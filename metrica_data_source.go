package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

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

func (ds *MetricsDataSource) CheckAndGetData(key *MetricaDataKey) (float64, error) {
	if err := ds.CheckAndUpdateData(); err != nil {
		return 0, err
	}

	prev, last, err := ds.GetOriginalData(key)

	if err != nil {
		return 0, err
	}
	return last - prev, nil
}
func (ds *MetricsDataSource) CheckAndGetLastData(key *MetricaDataKey) (float64, error) {
	if err := ds.CheckAndUpdateData(); err != nil {
		return 0, err
	}

	_, last, err := ds.GetOriginalData(key)

	if err != nil {
		return 0, err
	}
	return last, nil
}

func (ds *MetricsDataSource) GetOriginalData(key *MetricaDataKey) (float64, float64, error) {
	previousValueBlock, ok := ds.PreviousData[key.StatBlockKey]
	if !ok {
		return 0, 0, fmt.Errorf("Can not get data block from source \n")
	}
	currentValueBlock, ok := ds.LastData[key.StatBlockKey]
	if !ok {
		return 0, 0, fmt.Errorf("Can not get data block from source \n")
	}

	return previousValueBlock.GetValue(key.KeyInsideStatBlock), currentValueBlock.GetValue(key.KeyInsideStatBlock), nil
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
	// TODO: implement uptime check
	//If uptime is less then in previous run - then server were restarted
	/*
		if prev, last, err := ds.GetOriginalData("uptime"); err != nil {
			return err
		} else {
			if last < prev {
				ds.PreviousData = ds.LastData
			}
		}
	*/
	return nil
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

	data := make(SolrStatisticData, len(response.SolrInfo.QueryHandler.QueryHandlerInfo)+len(response.SolrInfo.UpdateHandler.QueryHandlerInfo)+len(response.SolrInfo.CacheHandler.QueryHandlerInfo)+1)
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
