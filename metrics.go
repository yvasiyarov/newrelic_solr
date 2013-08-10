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
    //update handler errors
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "updateHandler",
			KeyInsideStatBlock: "errors",
		},
		Name:  "handler/errors/DirectUpdateHandler2",
		Units: "errors/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "updateHandler",
			KeyInsideStatBlock: "cumulative_errors",
		},
		Name:  "handler/errors/DirectUpdateHandler2 cumulative errors",
		Units: "errors/seconds",
	},

	//timeouts  
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "spell",
			KeyInsideStatBlock: "timeouts",
		},
		Name:  "handler/timeouts/spell",
		Units: "timeouts/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "/update",
			KeyInsideStatBlock: "timeouts",
		},
		Name:  "handler/timeouts/update",
		Units: "timeouts/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "org.apache.solr.handler.XmlUpdateRequestHandler",
			KeyInsideStatBlock: "timeouts",
		},
		Name:  "handler/timeouts/org.apache.solr.handler.XmlUpdateRequestHandler",
		Units: "timeouts/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "standard",
			KeyInsideStatBlock: "timeouts",
		},
		Name:  "handler/timeouts/standard",
		Units: "timeouts/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "/suggest",
			KeyInsideStatBlock: "timeouts",
		},
		Name:  "handler/timeouts/suggest",
		Units: "timeouts/seconds",
	},

    //Direct updates handler metrics
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "updateHandler",
			KeyInsideStatBlock: "commits",
		},
		Name:  "DirectUpdateHandler2/commits",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "updateHandler",
			KeyInsideStatBlock: "autocommits",
		},
		Name:  "DirectUpdateHandler2/autocommits",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "updateHandler",
			KeyInsideStatBlock: "optimizes",
		},
		Name:  "DirectUpdateHandler2/optimizes",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "updateHandler",
			KeyInsideStatBlock: "rollbacks",
		},
		Name:  "DirectUpdateHandler2/rollbacks",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "updateHandler",
			KeyInsideStatBlock: "expungeDeletes",
		},
		Name:  "DirectUpdateHandler2/expungeDeletes",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "updateHandler",
			KeyInsideStatBlock: "adds",
		},
		Name:  "DirectUpdateHandler2/adds",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "updateHandler",
			KeyInsideStatBlock: "deletesById",
		},
		Name:  "DirectUpdateHandler2/deletesById",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "updateHandler",
			KeyInsideStatBlock: "deletesByQuery",
		},
		Name:  "DirectUpdateHandler2/deletesByQuery",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "updateHandler",
			KeyInsideStatBlock: "cumulative_adds",
		},
		Name:  "DirectUpdateHandler2/cumulative_adds",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "updateHandler",
			KeyInsideStatBlock: "cumulative_deletesById",
		},
		Name:  "DirectUpdateHandler2/cumulative_deletesById",
		Units: "requests/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "updateHandler",
			KeyInsideStatBlock: "cumulative_deletesByQuery",
		},
		Name:  "DirectUpdateHandler2/cumulative_deletesByQuery",
		Units: "requests/seconds",
	},
}
