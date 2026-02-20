package server

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"atlas.facade/internal/model"
)

type LogEntry struct {
	Timestamp time.Time
	Method    string
	Path      string
	Status    int
	Duration  time.Duration
}

type Server struct {
	Port      int
	Routes    []model.Route
	LogChan   chan LogEntry
	isRunning bool
}

func NewServer(port int, routes []model.Route) *Server {
	return &Server{
		Port:    port,
		Routes:  routes,
		LogChan: make(chan LogEntry, 100),
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	for _, r := range s.Routes {
		route := r // Capture for closure
		mux.HandleFunc(route.Path, func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			
			if req.Method != route.Method && route.Method != "" {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}

			// Simulate Latency
			if lat := route.GetLatency(); lat > 0 {
				time.Sleep(lat)
			}

			// Response
			w.Header().Set("Content-Type", "application/json") // Default to JSON
			w.WriteHeader(route.Status)
			io.WriteString(w, route.Body)

			// Log to TUI
			s.LogChan <- LogEntry{
				Timestamp: time.Now(),
				Method:    req.Method,
				Path:      req.URL.Path,
				Status:    route.Status,
				Duration:  time.Since(start),
			}
		})
	}

	addr := fmt.Sprintf(":%d", s.Port)
	return http.ListenAndServe(addr, mux)
}
