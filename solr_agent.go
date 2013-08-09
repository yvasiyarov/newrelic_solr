package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/yvasiyarov/newrelic_platform_go"
	"io/ioutil"
//	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var solrUrl = flag.String("solr-url", "127.0.0.1:8080/", "Solr url")
var newrelicLicense = flag.String("newrelic-license", "", "Newrelic license")
var verbose = flag.Bool("verbose", false, "Verbose mode")

const (
	MIN_PAUSE_TIME          = 30 //do not query sphinx often than once in 30 seconds
	SOLR_CONNECTION_TIMEOUT = 0  //no timeout
	NEWRELIC_POLL_INTERVAL  = 60 //Send data to newrelic every 60 seconds

	COMPONENT_NAME = "Solr"
	AGENT_GUID     = "com.github.yvasiyarov.Solr"
	AGENT_VERSION  = "0.0.1"
)

type MetricsDataSource struct {
	SolrUrl           string
	Port              int
	ConnectionTimeout int

	PreviousData   SolrStatisticData
	LastData       SolrStatisticData
	LastUpdateTime time.Time
}

type SolrStatisticData map[string]ISolrHandlerStat
type ISolrHandlerStat interface {
	Parse(info interface{}) error
	GetName() string
    GetValue(key string) float64
}

type SolrHandlerStat struct {
	Name      string
	ClassName string
    MetricaData map[string]float64
}

func (stat *SolrHandlerStat) GetName() string {
	return stat.Name
}


func (stat *SolrHandlerStat) GetValue(key string) float64 {
    if v, ok := stat.MetricaData[key]; ok {
        return v
    }
    return 0
}

func (stat *SolrHandlerStat) Parse(handlerInfo interface{}) error {
    switch info := handlerInfo.(type) {
        default:
            return fmt.Errorf("Parse of %#v is not implemented\n")
        case *SolrSystemResponse: {
            stat.Name = "Solr"
            stat.MetricaData = make(map[string]float64, 12)
            for _, item := range info.Info {
                itemName := strings.TrimSpace(item.ItemName);

                if itemName == "jvm" {
                    for _, intValue := range item.MemoryValues {
                        valueParts := strings.Split(strings.TrimSpace(intValue.Value), " ")
                        value, err := strconv.ParseFloat(valueParts[0], 64)
                        if err == nil {
                            switch valueParts[1] {
                            case "KB":
                                value = value * 1024;
                            case "MB":
                                value = value * 1024 * 1024;
                            case "GB":
                                value = value * 1024 * 1024 * 1024;
                            }
                            stat.MetricaData["jvm_memory_" + intValue.Name] = value
                        }
                    }
                }
                if itemName == "system" {
                    for _, intValue := range item.IntValues {
                        value, err := strconv.ParseFloat(strings.TrimSpace(intValue.Value), 64)
                        if err != nil {
                            return err
                        }
                        stat.MetricaData[intValue.Name] = value
                    }
                }
            }
            return nil
        }
        case *SolrQueryHandlerInfo: {
            stat.Name = strings.TrimSpace(info.Name)
            stat.MetricaData = make(map[string]float64, len(info.Stats))
            for _, statItem := range info.Stats {
                value, err := strconv.ParseFloat(strings.TrimSpace(statItem.Value), 64)
                if err != nil {
                    return err
                }
                stat.MetricaData[statItem.Name] = value
            }
        }
    }
	return nil
}


func NewMetricsDataSource(solrUrl string, connectionTimeout int) *MetricsDataSource {
	ds := &MetricsDataSource{
		SolrUrl:           solrUrl,
		ConnectionTimeout: connectionTimeout,
	}
	return ds
}

//Set of structures, used to parse XML response with solr handler statistic
type SolrResponse struct {
	SolrInfo SolrInfo `xml:"solr-info"`
}
type SolrInfo struct {
	QueryHandler  SolrQueryHandler `xml:"QUERYHANDLER"`
	UpdateHandler SolrQueryHandler `xml:"UPDATEHANDLER"`
	CacheHandler  SolrQueryHandler `xml:"CACHE"`
}
type SolrQueryHandler struct {
	QueryHandlerInfo []SolrQueryHandlerInfo `xml:"entry"`
}
type SolrQueryHandlerInfo struct {
	Name        string                     `xml:"name"`
	ClassName   string                     `xml:"class"`
	Version     string                     `xml:"version"`
	Description string                     `xml:"description"`
	Stats       []SolrQueryHandlerInfoItem `xml:"stats>stat"`
}
type SolrQueryHandlerInfoItem struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",innerxml"`
}

// Set of structures used to parse XML response with information about OS and JVM
type SolrSystemResponse struct {
	Info []SolrSystemInfoItem `xml:"lst"`
}
type SolrSystemInfoItem struct {
    ItemName string `xml:"name,attr"`
    IntValues []SolrSystemInfoItemValue `xml:"long"`
    MemoryValues []SolrSystemInfoItemValue `xml:"lst>str"`
}
type SolrSystemInfoItemValue struct {
    Name string `xml:"name,attr"`
    Value string `xml:",innerxml"`
}

//Query Solr handlers statistics
func (ds *MetricsDataSource) QueryData() (SolrStatisticData, error) {
	resp, err := http.Get("http://" + ds.SolrUrl + "admin/stats.jsp")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := SolrResponse{}
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	data := make(SolrStatisticData, len(response.SolrInfo.QueryHandler.QueryHandlerInfo) + len(response.SolrInfo.UpdateHandler.QueryHandlerInfo) + len(response.SolrInfo.CacheHandler.QueryHandlerInfo) + 1)
	parseQueryHandlers(response.SolrInfo.QueryHandler.QueryHandlerInfo, data)
	parseQueryHandlers(response.SolrInfo.UpdateHandler.QueryHandlerInfo, data)
	parseQueryHandlers(response.SolrInfo.CacheHandler.QueryHandlerInfo, data)

    if stat, err := ds.QuerySystemData(); err == nil {
        data["solr"] = stat
    }
	return data, nil
}

// parse statistic tag blocks
func parseQueryHandlers(queryHandlerInfo []SolrQueryHandlerInfo, data SolrStatisticData) {
	for _, handler := range queryHandlerInfo {
        solrClassName := strings.TrimSpace(handler.ClassName)
	    if solrClassName != "org.apache.solr.handler.component.SearchHandler" && solrClassName != "org.apache.solr.handler.XmlUpdateRequestHandler" && solrClassName != "org.apache.solr.update.DirectUpdateHandler2" && solrClassName != "org.apache.solr.search.LRUCache" && solrClassName != "org.apache.solr.search.FastLRUCache" {
        
                continue
        }

        stat := &SolrHandlerStat{ClassName: solrClassName}
		err := stat.Parse(&handler)
		if err != nil {
			continue
		}
		data[stat.GetName()] = stat
	}
}

//Query solr system information - OS and JVM memory consumption
func (ds *MetricsDataSource) QuerySystemData() (*SolrHandlerStat, error) {
	resp, err := http.Get("http://" + ds.SolrUrl + "admin/system/")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := SolrSystemResponse{}
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

    stat := &SolrHandlerStat{ClassName: "solr"}
    err = stat.Parse(&response)
    if err == nil {
        return stat, nil
    }

	return nil, err
}

func main() {
	flag.Parse()
	/*
		if *newrelicLicense == "" {
			log.Fatalf("Please, pass a valid newrelic license key.\n Use --help to get more information about available options\n")
		}
	*/
	plugin := newrelic_platform_go.NewNewrelicPlugin(AGENT_VERSION, *newrelicLicense, NEWRELIC_POLL_INTERVAL)
	component := newrelic_platform_go.NewPluginComponent(COMPONENT_NAME, AGENT_GUID)
	plugin.AddComponent(component)

	ds := NewMetricsDataSource(*solrUrl, SOLR_CONNECTION_TIMEOUT)
    d, _ := ds.QueryData()
	for k, v := range d {
		fmt.Printf("\n\nkey:%v, value:%#v\n", k, v)
	}
	//fmt.Printf("Error: %v\n, data:%#v\n", err, d)
	//AddMetrcas(component, ds)

	//	plugin.Verbose = *verbose
	//	plugin.Harvest()
	//plugin.Run()
}

/*
func (ds *MetricsDataSource) CheckAndGetData(key string) (float64, error) {
	if err := ds.CheckAndUpdateData(); err != nil {
		return 0, err
	}

	prev, last, err := ds.GetOriginalData(key)

	if err != nil {
		return 0, err
	}
	return last - prev, nil
}
func (ds *MetricsDataSource) CheckAndGetLastData(key string) (float64, error) {
	if err := ds.CheckAndUpdateData(); err != nil {
		return 0, err
	}

	_, last, err := ds.GetOriginalData(key)

	if err != nil {
		return 0, err
	}
	return last, nil
}

func (ds *MetricsDataSource) GetOriginalData(key string) (float64, float64, error) {
	previousValue, ok := ds.PreviousData[key]
	if !ok {
		return 0, 0, fmt.Errorf("Can not get data from source \n")
	}
	currentValue, ok := ds.LastData[key]
	if !ok {
		return 0, 0, fmt.Errorf("Can not get data from source \n")
	}

	//some metric calculation can be turned off by sphinx settings
	if previousValue == "OFF" || currentValue == "OFF" {
		return 0, 0, nil
	}

	previousValueConverted, err := strconv.ParseFloat(previousValue, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Can not convert previous value of %s to int \n", key)
	}
	currentValueConverted, err := strconv.ParseFloat(currentValue, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Can not convert current value of %s to int \n", key)
	}

	return previousValueConverted, currentValueConverted, nil
}

func (ds *MetricsDataSource) CheckAndUpdateData() error {
	startTime := time.Now()
	if startTime.Sub(ds.LastUpdateTime) > time.Second*MIN_PAUSE_TIME {
		newData, err := ds.QueryData()
		if err != nil {
			return err
		}

		if ds.PreviousData == nil {
			ds.PreviousData = newData
		} else {
			ds.PreviousData = ds.LastData
		}
		ds.LastData = newData
		ds.LastUpdateTime = startTime
	}

	// check uptime
	//If uptime is less then in previous run - then server were restarted
	if prev, last, err := ds.GetOriginalData("uptime"); err != nil {
		return err
	} else {
		if last < prev {
			ds.PreviousData = ds.LastData
		}
	}
	return nil
}
*/
/*
type Metrica struct {
Name       string
Units      string
DataKey    string
DataSource *MetricsDataSource
}

func (metrica *Metrica) GetName() string {
return metrica.Name
}
func (metrica *Metrica) GetUnits() string {
return metrica.Units
}
func (metrica *Metrica) GetValue() (float64, error) {
return metrica.DataSource.CheckAndGetData(metrica.DataKey)
}

type AvgMetrica struct {
Metrica
}

func (metrica *AvgMetrica) GetValue() (float64, error) {
return metrica.DataSource.CheckAndGetLastData(metrica.DataKey)
}

func AddMetrcas(component newrelic_platform_go.IComponent, dataSource *MetricsDataSource) {
metricas := []*Metrica{
&Metrica{
DataKey: "queries",
Name:    "Queries",
Units:   "Queries/second",
},
&Metrica{
DataKey: "connections",
Name:    "Connections",
Units:   "connections/second",
},
&Metrica{
DataKey: "maxed_out",
Name:    "Maxed out",
Units:   "connections/second",
},
&Metrica{
DataKey: "command_search",
			Name:    "Command search",
			Units:   "command/second",
		},
		&Metrica{
			DataKey: "command_excerpt",
			Name:    "Command excerpt",
			Units:   "command/second",
		},
		&Metrica{
			DataKey: "command_update",
			Name:    "Command update",
			Units:   "command/second",
		},
		&Metrica{
			DataKey: "command_keywords",
			Name:    "Command keywords",
			Units:   "command/second",
		},
		&Metrica{
			DataKey: "command_persist",
			Name:    "Command persist",
			Units:   "command/second",
		},
		&Metrica{
			DataKey: "command_flushattrs",
			Name:    "Command flushattrs",
			Units:   "command/second",
		},
	}
	for _, m := range metricas {
		m.DataSource = dataSource
		component.AddMetrica(m)
	}

	avgQueryWallTime := &AvgMetrica{
		Metrica{
			DataKey: "avg_query_wall",
			Name:    "Avg Query Wall Time",
			Units:   "milisecond",
		},
	}
	avgQueryWallTime.DataSource = dataSource
	component.AddMetrica(avgQueryWallTime)
}
*/
