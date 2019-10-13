package metrics

import (
	"expvar"

	"github.com/elastic/go-elasticsearch/v8/estransport"
)

func init() {
	estransport.EnableMetrics()
	expvar.Publish("goelasticsearch", expvar.Func(estransport.MetricFunc))
}
