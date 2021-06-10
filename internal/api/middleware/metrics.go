package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/cors"
)

// source: https://github.com/go-swagger/go-swagger/issues/1120#issuecomment-419577629

var histogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_duration_seconds",
	Help: "Time taken to execute endpoint.",
}, []string{"path", "method", "status"})

type metricResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newMetricResponseWriter(w http.ResponseWriter) *metricResponseWriter {
	return &metricResponseWriter{w, http.StatusOK}
}

func (lrw *metricResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// SetupHandler enable CORS, handler metrics.
func SetupHandler(handler http.Handler) http.Handler {
	handleCORS := cors.AllowAll().Handler
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := newMetricResponseWriter(w)
		handler.ServeHTTP(lrw, r)
		statusCode := lrw.statusCode
		duration := time.Since(start)
		histogram.WithLabelValues(r.URL.String(), r.Method, fmt.Sprintf("%d", statusCode)).Observe(duration.Seconds())
	})
	return handleCORS(h)
}
