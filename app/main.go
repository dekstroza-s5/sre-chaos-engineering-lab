package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync/atomic"
	"syscall"
	"time"
)

var mode atomic.Value
var requests atomic.Uint64
var errorsTotal atomic.Uint64
var durationMillis atomic.Uint64

func main() {
	mode.Store("healthy")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("GET /work", work)
	mux.HandleFunc("GET /metrics", metrics)
	mux.HandleFunc("POST /admin/failure", setFailure)
	server := &http.Server{Addr: ":8080", Handler: accessLog(mux), ReadHeaderTimeout: 5 * time.Second}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	go func() { log.Printf("demo-api listening on :8080"); if err := server.ListenAndServe(); err != http.ErrServerClosed { log.Fatal(err) } }()
	<-stop
	_ = server.Close()
}

func health(w http.ResponseWriter, _ *http.Request) {
	if mode.Load() == "unhealthy" { writeJSON(w, 503, map[string]string{"status":"unhealthy"}); return }
	writeJSON(w, 200, map[string]string{"status":"ok"})
}

func work(w http.ResponseWriter, _ *http.Request) {
	start := time.Now(); requests.Add(1)
	switch mode.Load() {
	case "error":
		errorsTotal.Add(1); writeJSON(w, 500, map[string]string{"error":"injected failure"})
	case "latency":
		time.Sleep(750 * time.Millisecond); writeJSON(w, 200, map[string]string{"result":"slow"})
	default:
		time.Sleep(time.Duration(10+rand.Intn(40)) * time.Millisecond); writeJSON(w, 200, map[string]string{"result":"ok"})
	}
	durationMillis.Add(uint64(time.Since(start).Milliseconds()))
}

func setFailure(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("mode")
	if value != "healthy" && value != "error" && value != "latency" && value != "unhealthy" { writeJSON(w, 400, map[string]string{"error":"unsupported mode"}); return }
	mode.Store(value); writeJSON(w, 200, map[string]string{"mode":value})
}

func metrics(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	total := requests.Load(); failed := errorsTotal.Load()
	_, _ = w.Write([]byte("# TYPE demo_http_requests_total counter\n"))
	_, _ = w.Write([]byte("demo_http_requests_total{status=\"2xx\"} " + strconv.FormatUint(total-failed,10) + "\n"))
	_, _ = w.Write([]byte("demo_http_requests_total{status=\"5xx\"} " + strconv.FormatUint(failed,10) + "\n"))
	_, _ = w.Write([]byte("# TYPE demo_request_duration_milliseconds_total counter\n"))
	_, _ = w.Write([]byte("demo_request_duration_milliseconds_total " + strconv.FormatUint(durationMillis.Load(),10) + "\n"))
}

func writeJSON(w http.ResponseWriter, status int, value any) { w.Header().Set("Content-Type","application/json"); w.WriteHeader(status); _ = json.NewEncoder(w).Encode(value) }
func accessLog(next http.Handler) http.Handler { return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){start:=time.Now();next.ServeHTTP(w,r);log.Printf(`{"method":%q,"path":%q,"duration_ms":%d}`,r.Method,r.URL.Path,time.Since(start).Milliseconds())}) }
