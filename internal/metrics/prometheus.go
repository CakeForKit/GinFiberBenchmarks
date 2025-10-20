package metrics

import (
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Gatherer = prometheus.NewRegistry()
	// отслеживает количество currently выполняющихся HTTP-запросов
	// ActiveHttpRequestCounter = prometheus.NewGaugeVec(
	// 	prometheus.GaugeOpts{
	// 		Name: "active_http_request_counter",
	// 		Help: "The number of active http requests",
	// 	},
	// 	[]string{"method", "path"},
	// )
	// подсчитывает общее количество обработанных HTTP-запросов
	TotalHttpRequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_http_request_counter",
			Help: "The number of total http requests",
		},
		[]string{"method", "path"},
	)
	// отслеживает использование памяти приложением
	MemoryAllocations = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "memory_allocations_bytes",
		Help: "Memory allocations in bytes",
	}, []string{"type"})
	// отслеживает текущее количество активных горутин
	GoroutinesCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "goroutines_count",
		Help: "Number of goroutines",
	})
)

func RegisterMetricCollector(sleepMs int) {
	// Gatherer.MustRegister(ActiveHttpRequestCounter)
	Gatherer.MustRegister(TotalHttpRequestCounter)
	Gatherer.MustRegister(GoroutinesCount)
	Gatherer.MustRegister(MemoryAllocations)
	go func() {
		for {
			GoroutinesCount.Set(float64(runtime.NumGoroutine()))
			var memoryStats runtime.MemStats
			runtime.ReadMemStats(&memoryStats)
			MemoryAllocations.WithLabelValues("heap").Set(float64(memoryStats.HeapAlloc))
			MemoryAllocations.WithLabelValues("stack").Set(float64(memoryStats.StackInuse))
			time.Sleep(time.Duration(sleepMs) * time.Millisecond) // как часто приложение ОБНОВЛЯЕТ значения метрик
		}
	}()
}

func GetHttpHandler() http.Handler {
	return promhttp.HandlerFor(Gatherer, promhttp.HandlerOpts{})
}
