package metrics

func (d dummyMetrics) Enabled() bool {
	return false
}
func (d dummyMetrics) Observe(reporter Reporter) {}

func (d dummyMetrics) Log(handler, method, code string, sendBytes, recvBytes, latency float64) {}

func (d dummyMetrics) RequestTotal(module, handler, method, code string) {}

func (d dummyMetrics) RequestDurationSeconds(module, handler, method string, latency float64) {}

func (d dummyMetrics) SummaryLatencyLog(module, handler, method string, latency float64) {}

func (d dummyMetrics) ErrorsTotal(module, errors string) {}

func (d dummyMetrics) Event(module, event string) {}

func (d dummyMetrics) SiteEvent(module, event, site string) {}

func (d dummyMetrics) RequestsInFlight(module, state string, value float64) {}

func (d dummyMetrics) ResponseSize(module, handler, method, code string, length float64) {}

func (d dummyMetrics) RequestSize(module, handler, method, code string, length float64) {}

func (d dummyMetrics) ResetCounter() {}

var DummyMetrics Metrics = dummyMetrics{}
