package main

import (
	"fmt"
	"strconv"
	"strings"
)

type SolrStatisticData map[string]ISolrHandlerStat
type ISolrHandlerStat interface {
	Parse(info interface{}) error
	GetName() string
	GetValue(key string) float64
}

type SolrHandlerStat struct {
	Name        string
	ClassName   string
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
	case *SolrSystemResponse:
		{
			stat.Name = "Solr"
			stat.MetricaData = make(map[string]float64, 12)
			for _, item := range info.Info {
				itemName := strings.TrimSpace(item.ItemName)

				if itemName == "jvm" {
					for _, intValue := range item.MemoryValues {
						valueParts := strings.Split(strings.TrimSpace(intValue.Value), " ")
						value, err := strconv.ParseFloat(valueParts[0], 64)
						if err == nil {
							switch valueParts[1] {
							case "KB":
								value = value * 1024
							case "MB":
								value = value * 1024 * 1024
							case "GB":
								value = value * 1024 * 1024 * 1024
							}
							stat.MetricaData["jvm_memory_"+intValue.Name] = value
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
	case *SolrQueryHandlerInfo:
		{
			stat.Name = strings.TrimSpace(info.Name)
			stat.MetricaData = make(map[string]float64, len(info.Stats))
			for _, statItem := range info.Stats {
				value, err := strconv.ParseFloat(strings.TrimSpace(statItem.Value), 64)
				if err == nil {
                    stat.MetricaData[statItem.Name] = value
                }
			}
		}
	}
	return nil
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
	ItemName     string                    `xml:"name,attr"`
	IntValues    []SolrSystemInfoItemValue `xml:"long"`
	MemoryValues []SolrSystemInfoItemValue `xml:"lst>str"`
}
type SolrSystemInfoItemValue struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",innerxml"`
}
