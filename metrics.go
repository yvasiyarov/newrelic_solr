package main

type Metrica struct {
	Name       string
	Units      string
	DataKey    *MetricaDataKey
	DataSource *MetricsDataSource
}
type MetricaDataKey struct {
	StatBlockKey       string
	KeyInsideStatBlock string
}

func (metrica *Metrica) GetName() string {
	return metrica.Name
}
func (metrica *Metrica) GetUnits() string {
	return metrica.Units
}
func (metrica *Metrica) GetValue() (float64, error) {
	return metrica.DataSource.CheckAndGetLastData(metrica.DataKey)
}

type IncrementalMetrica struct {
	Metrica
}

func (metrica *IncrementalMetrica) GetValue() (float64, error) {
	return metrica.DataSource.CheckAndGetData(metrica.DataKey)
}

var plainMetricas = []*Metrica{
	// Solr memory metrics
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "solr",
			KeyInsideStatBlock: "jvm_memory_used",
		},
		Name:  "solr/memory/JVM memory used",
		Units: "bytes",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "solr",
			KeyInsideStatBlock: "jvm_memory_free",
		},
		Name:  "solr/memory/JVM memory free",
		Units: "bytes",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "solr",
			KeyInsideStatBlock: "jvm_memory_total",
		},
		Name:  "solr/memory/JVM memory total",
		Units: "bytes",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "solr",
			KeyInsideStatBlock: "freePhysicalMemorySize",
		},
		Name:  "solr/memory/Free Physical Memory Size",
		Units: "bytes",
	},

	//Avg request per second
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "spell",
			KeyInsideStatBlock: "avgRequestsPerSecond",
		},
		Name:  "handler/request_per_second/spell",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "/update",
			KeyInsideStatBlock: "avgRequestsPerSecond",
		},
		Name:  "handler/request_per_second/update",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "org.apache.solr.handler.XmlUpdateRequestHandler",
			KeyInsideStatBlock: "avgRequestsPerSecond",
		},
		Name:  "handler/request_per_second/org.apache.solr.handler.XmlUpdateRequestHandler",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "standard",
			KeyInsideStatBlock: "avgRequestsPerSecond",
		},
		Name:  "handler/request_per_second/standard",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "/suggest",
			KeyInsideStatBlock: "avgRequestsPerSecond",
		},
		Name:  "handler/request_per_second/suggest",
		Units: "requests/seconds",
	},
	//Time per request
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "spell",
			KeyInsideStatBlock: "avgTimePerRequest",
		},
		Name:  "handler/time_per_request/spell",
		Units: "seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "/update",
			KeyInsideStatBlock: "avgTimePerRequest",
		},
		Name:  "handler/time_per_request/update",
		Units: "seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "org.apache.solr.handler.XmlUpdateRequestHandler",
			KeyInsideStatBlock: "avgTimePerRequest",
		},
		Name:  "handler/time_per_request/org.apache.solr.handler.XmlUpdateRequestHandler",
		Units: "seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "standard",
			KeyInsideStatBlock: "avgTimePerRequest",
		},
		Name:  "handler/time_per_request/standard",
		Units: "seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "/suggest",
			KeyInsideStatBlock: "avgTimePerRequest",
		},
		Name:  "handler/time_per_request/suggest",
		Units: "seconds",
	},
}

//Incremental metricas
var incrementalMetricas = []*Metrica{
	//Errors 
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "spell",
			KeyInsideStatBlock: "errors",
		},
		Name:  "handler/errors/spell",
		Units: "errors/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "/update",
			KeyInsideStatBlock: "errors",
		},
		Name:  "handler/errors/update",
		Units: "errors/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "org.apache.solr.handler.XmlUpdateRequestHandler",
			KeyInsideStatBlock: "errors",
		},
		Name:  "handler/errors/org.apache.solr.handler.XmlUpdateRequestHandler",
		Units: "errors/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "standard",
			KeyInsideStatBlock: "errors",
		},
		Name:  "handler/errors/standard",
		Units: "errors/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "/suggest",
			KeyInsideStatBlock: "errors",
		},
		Name:  "handler/errors/suggest",
		Units: "errors/seconds",
	},
}
