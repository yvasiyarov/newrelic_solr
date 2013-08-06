package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/yvasiyarov/newrelic_platform_go"
	"io/ioutil"
	"log"
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

type SolrStatisticData map[string]ISolrHandlerStat
type ISolrHandlerStat interface {
	Parse(info *SolrQueryHandlerInfo) error
	GetName() string
}

type SolrHandlerStat struct {
	Name      string
	ClassName string
}

func (stat *SolrHandlerStat) GetName() string {
	return stat.Name
}

func NewSolrHandlerStat(solrClassName string) ISolrHandlerStat {
	solrClassName = strings.TrimSpace(solrClassName)
	switch solrClassName {
	case "org.apache.solr.handler.component.SearchHandler", "org.apache.solr.handler.XmlUpdateRequestHandler":
		{
			return &SolrHandlerStatSearch{SolrHandlerStat: SolrHandlerStat{ClassName: solrClassName}}
		}
	case "org.apache.solr.update.DirectUpdateHandler2":
		{
			return &SolrHandlerStatUpdate{SolrHandlerStat: SolrHandlerStat{ClassName: solrClassName}}
		}
	case "org.apache.solr.search.LRUCache", "org.apache.solr.search.FastLRUCache":
		{
			return &SolrHandlerStatCache{SolrHandlerStat: SolrHandlerStat{ClassName: solrClassName}}
		}
	}
	return nil
}

type SolrHandlerStatSearch struct {
	SolrHandlerStat
	Requests             float64
	Errors               float64
	Timeouts             float64
	TotalTime            float64
	AvgTimePerRequest    float64
	AvgRequestsPerSecond float64
}

func (stat *SolrHandlerStatSearch) Parse(info *SolrQueryHandlerInfo) error {
	stat.Name = strings.TrimSpace(info.Name)
	for _, statItem := range info.Stats {
		value, err := strconv.ParseFloat(strings.TrimSpace(statItem.Value), 64)
		if err != nil {
			return err
		}
		switch statItem.Name {
		case "avgRequestsPerSecond":
			{
				stat.AvgRequestsPerSecond = value
			}
		case "avgTimePerRequest":
			{
				stat.AvgTimePerRequest = value
			}
		case "totalTime":
			{
				stat.TotalTime = value
			}
		case "timeouts":
			{
				stat.Timeouts = value
			}
		case "errors":
			{
				stat.Errors = value
			}
		case "requests":
			{
				stat.Requests = value
			}
		}
	}
	return nil
}

type SolrHandlerStatUpdate struct {
	SolrHandlerStat
	Commits                  float64
	Autocommits              float64
	Optimizes                float64
	Rollbacks                float64
	ExpungeDeletes           float64
	DocsPending              float64
	Adds                     float64
	DeletesById              float64
	DeletesByQuery           float64
	Errors                   float64
	CumulativeAdds           float64
	CumulativeDeletesById    float64
	CumulativeDeletesByQuery float64
	CumulativeErrors         float64
}

func (stat *SolrHandlerStatUpdate) Parse(info *SolrQueryHandlerInfo) error {
	stat.Name = strings.TrimSpace(info.Name)
	for _, statItem := range info.Stats {
		value, err := strconv.ParseFloat(strings.TrimSpace(statItem.Value), 64)
		if err != nil {
			return err
		}
		switch statItem.Name {
		case "commits":
			{
				stat.Commits = value
			}
		case "autocommits":
			{
				stat.Autocommits = value
			}
		case "optimizes":
			{
				stat.Optimizes = value
			}
		case "rollbacks":
			{
				stat.Rollbacks = value
			}
		case "expungeDeletes":
			{
				stat.ExpungeDeletes = value
			}
		case "docsPending":
			{
				stat.DocsPending = value
			}
		case "adds":
			{
				stat.Adds = value
			}
		case "deletesById":
			{
				stat.DeletesById = value
			}
		case "deletesByQuery":
			{
				stat.DeletesByQuery = value
			}
		case "errors":
			{
				stat.Errors = value
			}
		case "cumulative_adds":
			{
				stat.CumulativeAdds = value
			}
		case "cumulative_deletesById":
			{
				stat.CumulativeDeletesById = value
			}
		case "cumulative_deletesByQuery":
			{
				stat.CumulativeDeletesByQuery = value
			}
		case "cumulative_errors":
			{
				stat.CumulativeErrors = value
			}
		}
	}
	return nil
}

type SolrHandlerStatCache struct {
	SolrHandlerStat
	Lookups             float64
	Hits                float64
	HitRatio            float64
	Inserts             float64
	Evictions           float64
	Size                float64
	CumulativeLookups   float64
	CumulativeHits      float64
	CumulativeHitRatio  float64
	CumulativeInserts   float64
	CumulativeEvictions float64
}

func (stat *SolrHandlerStatCache) Parse(info *SolrQueryHandlerInfo) error {
	stat.Name = strings.TrimSpace(info.Name)
	for _, statItem := range info.Stats {
		value, err := strconv.ParseFloat(strings.TrimSpace(statItem.Value), 64)
		if err != nil {
			return err
		}
		switch statItem.Name {
		case "lookups":
			{
				stat.Lookups = value
			}
		case "hits":
			{
				stat.Hits = value
			}
		case "hitratio":
			{
				stat.HitRatio = value
			}
		case "inserts":
			{
				stat.Inserts = value
			}
		case "evictions":
			{
				stat.Evictions = value
			}
		case "size":
			{
				stat.Size = value
			}
		case "cumulative_lookups":
			{
				stat.CumulativeLookups = value
			}
		case "cumulative_hits":
			{
				stat.CumulativeHits = value
			}
		case "cumulative_hitratio":
			{
				stat.CumulativeHitRatio = value
			}
		case "cumulative_inserts":
			{
				stat.CumulativeInserts = value
			}
		case "cumulative_evictions":
			{
				stat.CumulativeEvictions = value
			}
		}
	}
	return nil
}

type MetricsDataSource struct {
	SolrUrl           string
	Port              int
	ConnectionTimeout int

	PreviousData   SolrStatisticData
	LastData       SolrStatisticData
	LastUpdateTime time.Time
}

func NewMetricsDataSource(solrUrl string, connectionTimeout int) *MetricsDataSource {
	ds := &MetricsDataSource{
		SolrUrl:           solrUrl,
		ConnectionTimeout: connectionTimeout,
	}
	return ds
}

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

	data := make(SolrStatisticData, len(response.SolrInfo.QueryHandler.QueryHandlerInfo) + len(response.SolrInfo.UpdateHandler.QueryHandlerInfo) + len(response.SolrInfo.CacheHandler.QueryHandlerInfo))
	parseQueryHandlers(response.SolrInfo.QueryHandler.QueryHandlerInfo, data)
	parseQueryHandlers(response.SolrInfo.UpdateHandler.QueryHandlerInfo, data)
	parseQueryHandlers(response.SolrInfo.CacheHandler.QueryHandlerInfo, data)

	return data, nil
}

func parseQueryHandlers(queryHandlerInfo []SolrQueryHandlerInfo, data SolrStatisticData) {
	for _, handler := range queryHandlerInfo {
		stat := NewSolrHandlerStat(handler.ClassName)
		if stat == nil {
			continue
		}
		err := stat.Parse(&handler)
		if err != nil {
			log.Printf("Handler stat parse error:%v. Handler: %#v", err, handler)
			continue
		}
		data[stat.GetName()] = stat
	}
}

type SolrSystemResponse struct {
	Info []SolrSystemInfoItem `xml:"lst"`
}
type SolrSystemInfoItem struct {
    ItemName string `xml:"name,attr"`
    IntValues []SolrSystemInfoItemValue `xml:"long"`
    MemoryValues []SolrSystemInfoItemValue `xml:"lst>str"`
}
type SolrSystemInfoItemValue struct {
    ValueName string `xml:"name,attr"`
    Value string `xml:",innerxml"`
}

func (ds *MetricsDataSource) QuerySystemData() (SolrStatisticData, error) {
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
    fmt.Printf("System resp %#v", response)

	return nil, nil
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
	ds.QueryData()
	ds.QuerySystemData()
	//for k, v := range d {
	//	fmt.Printf("key:%v, value:%#v\n", k, v)
	//}
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
