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

    //Cache hitratio non cumulative
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "queryResultCache",
			KeyInsideStatBlock: "hitratio",
		},
		Name:  "handler/cache/hitrates/queryResultCache",
		Units: "seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "documentCache",
			KeyInsideStatBlock: "hitratio",
		},
		Name:  "handler/cache/hitrates/documentCache",
		Units: "seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "fieldValueCache",
			KeyInsideStatBlock: "hitratio",
		},
		Name:  "handler/cache/hitrates/fieldValueCache",
		Units: "seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "filterCache",
			KeyInsideStatBlock: "hitratio",
		},
		Name:  "handler/cache/hitrates/filterCache",
		Units: "seconds",
	},
    //Cache hitratio cumulative
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "queryResultCache",
			KeyInsideStatBlock: "cumulative_hitratio",
		},
		Name:  "handler/cache/hitrates_cumulative/queryResultCache",
		Units: "seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "documentCache",
			KeyInsideStatBlock: "cumulative_hitratio",
		},
		Name:  "handler/cache/hitrates_cumulative/documentCache",
		Units: "seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "fieldValueCache",
			KeyInsideStatBlock: "cumulative_hitratio",
		},
		Name:  "handler/cache/hitrates_cumulative/fieldValueCache",
		Units: "seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "filterCache",
			KeyInsideStatBlock: "cumulative_hitratio",
		},
		Name:  "handler/cache/hitrates_cumulative/filterCache",
		Units: "seconds",
	},
    //Cache size
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "queryResultCache",
			KeyInsideStatBlock: "size",
		},
		Name:  "handler/cache/size/queryResultCache",
		Units: "seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "documentCache",
			KeyInsideStatBlock: "size",
		},
		Name:  "handler/cache/size/documentCache",
		Units: "seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "fieldValueCache",
			KeyInsideStatBlock: "size",
		},
		Name:  "handler/cache/size/fieldValueCache",
		Units: "seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "filterCache",
			KeyInsideStatBlock: "size",
		},
		Name:  "handler/cache/size/filterCache",
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

    //queryResultCache detail metrics
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "queryResultCache",
			KeyInsideStatBlock: "lookups",
		},
		Name:  "handler/cache/queryResultCache/lookups",
		Units: "request/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "queryResultCache",
			KeyInsideStatBlock: "hits",
		},
		Name:  "handler/cache/queryResultCache/hits",
		Units: "hits/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "queryResultCache",
			KeyInsideStatBlock: "inserts",
		},
		Name:  "handler/cache/queryResultCache/inserts",
		Units: "inserts/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "queryResultCache",
			KeyInsideStatBlock: "evictions",
		},
		Name:  "handler/cache/queryResultCache/evictions",
		Units: "evictions/seconds",
	},

    //queryResultCache cumulative detail metrics
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "queryResultCache",
			KeyInsideStatBlock: "cumulative_lookups",
		},
		Name:  "handler/cache/queryResultCache/cumulative_lookups",
		Units: "request/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "queryResultCache",
			KeyInsideStatBlock: "cumulative_hits",
		},
		Name:  "handler/cache/queryResultCache/cumulative_hits",
		Units: "hits/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "queryResultCache",
			KeyInsideStatBlock: "cumulative_inserts",
		},
		Name:  "handler/cache/queryResultCache/cumulative_inserts",
		Units: "inserts/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "queryResultCache",
			KeyInsideStatBlock: "cumulative_evictions",
		},
		Name:  "handler/cache/queryResultCache/cumulative_evictions",
		Units: "evictions/seconds",
	},

    //documentCache detail metrics
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "documentCache",
			KeyInsideStatBlock: "lookups",
		},
		Name:  "handler/cache/documentCache/lookups",
		Units: "request/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "documentCache",
			KeyInsideStatBlock: "hits",
		},
		Name:  "handler/cache/documentCache/hits",
		Units: "hits/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "documentCache",
			KeyInsideStatBlock: "inserts",
		},
		Name:  "handler/cache/documentCache/inserts",
		Units: "inserts/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "documentCache",
			KeyInsideStatBlock: "evictions",
		},
		Name:  "handler/cache/documentCache/evictions",
		Units: "evictions/seconds",
	},

    //documentCache cumulative detail metrics
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "documentCache",
			KeyInsideStatBlock: "cumulative_lookups",
		},
		Name:  "handler/cache/documentCache/cumulative_lookups",
		Units: "request/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "documentCache",
			KeyInsideStatBlock: "cumulative_hits",
		},
		Name:  "handler/cache/documentCache/cumulative_hits",
		Units: "hits/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "documentCache",
			KeyInsideStatBlock: "cumulative_inserts",
		},
		Name:  "handler/cache/documentCache/cumulative_inserts",
		Units: "inserts/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "documentCache",
			KeyInsideStatBlock: "cumulative_evictions",
		},
		Name:  "handler/cache/documentCache/cumulative_evictions",
		Units: "evictions/seconds",
	},

    //fieldValueCache detail metrics
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "fieldValueCache",
			KeyInsideStatBlock: "lookups",
		},
		Name:  "handler/cache/fieldValueCache/lookups",
		Units: "request/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "fieldValueCache",
			KeyInsideStatBlock: "hits",
		},
		Name:  "handler/cache/fieldValueCache/hits",
		Units: "hits/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "fieldValueCache",
			KeyInsideStatBlock: "inserts",
		},
		Name:  "handler/cache/fieldValueCache/inserts",
		Units: "inserts/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "fieldValueCache",
			KeyInsideStatBlock: "evictions",
		},
		Name:  "handler/cache/fieldValueCache/evictions",
		Units: "evictions/seconds",
	},

    //fieldValueCache cumulative detail metrics
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "fieldValueCache",
			KeyInsideStatBlock: "cumulative_lookups",
		},
		Name:  "handler/cache/fieldValueCache/cumulative_lookups",
		Units: "request/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "fieldValueCache",
			KeyInsideStatBlock: "cumulative_hits",
		},
		Name:  "handler/cache/fieldValueCache/cumulative_hits",
		Units: "hits/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "fieldValueCache",
			KeyInsideStatBlock: "cumulative_inserts",
		},
		Name:  "handler/cache/fieldValueCache/cumulative_inserts",
		Units: "inserts/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "fieldValueCache",
			KeyInsideStatBlock: "cumulative_evictions",
		},
		Name:  "handler/cache/fieldValueCache/cumulative_evictions",
		Units: "evictions/seconds",
	},

    //filterCache detail metrics
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "filterCache",
			KeyInsideStatBlock: "lookups",
		},
		Name:  "handler/cache/filterCache/lookups",
		Units: "request/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "filterCache",
			KeyInsideStatBlock: "hits",
		},
		Name:  "handler/cache/filterCache/hits",
		Units: "hits/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "filterCache",
			KeyInsideStatBlock: "inserts",
		},
		Name:  "handler/cache/filterCache/inserts",
		Units: "inserts/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "filterCache",
			KeyInsideStatBlock: "evictions",
		},
		Name:  "handler/cache/filterCache/evictions",
		Units: "evictions/seconds",
	},

    //filterCache cumulative detail metrics
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "filterCache",
			KeyInsideStatBlock: "cumulative_lookups",
		},
		Name:  "handler/cache/filterCache/cumulative_lookups",
		Units: "request/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "filterCache",
			KeyInsideStatBlock: "cumulative_hits",
		},
		Name:  "handler/cache/filterCache/cumulative_hits",
		Units: "hits/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "filterCache",
			KeyInsideStatBlock: "cumulative_inserts",
		},
		Name:  "handler/cache/filterCache/cumulative_inserts",
		Units: "inserts/seconds",
	},
	&Metrica{
		DataKey: &MetricaDataKey{
			StatBlockKey:       "filterCache",
			KeyInsideStatBlock: "cumulative_evictions",
		},
		Name:  "handler/cache/filterCache/cumulative_evictions",
		Units: "evictions/seconds",
	},
}
