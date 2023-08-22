package resource

import "github.com/prometheus/client_golang/prometheus"

var (
	//HTTPReqDuration metric:http_request_duration_seconds
	HTTPReqDuration *prometheus.HistogramVec
	//HTTPReqTotal metric:http_request_total
	HTTPReqTotal *prometheus.CounterVec
)

const (
	PromMetricHTTPRequestDurationSecondsName = "http_request_duration_seconds"
	PromMetricHTTPRequestDurationSecondsHelp = "The HTTP request latencies in seconds."
	PromMetricHTTPRequestTotalName           = "http_requests_total"
	PromMetricHTTPRequestTotalHelp           = "Total number of HTTP requests made."
)

// prometheus labels name
const (
	PromLabelMethod = "method" //http method
	PromLabelPath   = "path"   //http path
	PromLabelStatus = "status" //http response status
)

func init() {
	HTTPReqDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    PromMetricHTTPRequestDurationSecondsName,
		Help:    PromMetricHTTPRequestDurationSecondsHelp,
		Buckets: nil,
	}, []string{PromLabelMethod, PromLabelPath})

	HTTPReqTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: PromMetricHTTPRequestTotalName,
		Help: PromMetricHTTPRequestTotalHelp,
	}, []string{PromLabelMethod, PromLabelPath, PromLabelStatus})

	prometheus.MustRegister(
		HTTPReqDuration,
		HTTPReqTotal,
	)
}
