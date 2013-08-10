package main

import (
	"flag"
	"github.com/yvasiyarov/newrelic_platform_go"
	"log"
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

func addMetrcasToComponent(component newrelic_platform_go.IComponent, metricas []newrelic_platform_go.IMetrica) {
	for _, m := range metricas {
		component.AddMetrica(m)
	}
}

func plainMetricasBuilder(metricas []*Metrica, dataSource *MetricsDataSource) []newrelic_platform_go.IMetrica {
	result := make([]newrelic_platform_go.IMetrica, len(metricas))
	for i, m := range metricas {
		m.DataSource = dataSource
		result[i] = m
	}
	return result
}
func incrementalMetricasBuilder(metricas []*Metrica, dataSource *MetricsDataSource) []newrelic_platform_go.IMetrica {
	incMetricas := make([]newrelic_platform_go.IMetrica, len(metricas))
	for i, m := range metricas {
		m.DataSource = dataSource
		incMetricas[i] = &IncrementalMetrica{*m}
	}
	return incMetricas
}

func main() {
	flag.Parse()

	if *newrelicLicense == "" {
		log.Fatalf("Please, pass a valid newrelic license key.\n Use --help to get more information about available options\n")
	}
    log.Printf("Total metrics:%d\n", len(plainMetricas) + len(incrementalMetricas))

	plugin := newrelic_platform_go.NewNewrelicPlugin(AGENT_VERSION, *newrelicLicense, NEWRELIC_POLL_INTERVAL)
	component := newrelic_platform_go.NewPluginComponent(COMPONENT_NAME, AGENT_GUID)
	plugin.AddComponent(component)

	ds := NewMetricsDataSource(*solrUrl, SOLR_CONNECTION_TIMEOUT)
	addMetrcasToComponent(component, plainMetricasBuilder(plainMetricas, ds))
	addMetrcasToComponent(component, incrementalMetricasBuilder(incrementalMetricas, ds))

	plugin.Verbose = *verbose
	plugin.Run()
}
