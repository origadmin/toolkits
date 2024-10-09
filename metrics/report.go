package metrics

type Report struct {
	Type     MetricType
	TraceID  string
	Endpoint string
	Method   string
	Code     string
	SendSize int64
	RecvSize int64
	Latency  int64
	Succeed  bool
}
