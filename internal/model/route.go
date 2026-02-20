package model

import "time"

type Blueprint struct {
	Routes []Route `piml:"routes>route"`
}

type Route struct {
	Path    string `piml:"path"`
	Method  string `piml:"method"`
	Status  int    `piml:"status"`
	Body    string `piml:"body"`
	Latency string `piml:"latency"` // e.g. "500ms"
}

func (r Route) GetLatency() time.Duration {
	d, err := time.ParseDuration(r.Latency)
	if err != nil {
		return 0
	}
	return d
}
