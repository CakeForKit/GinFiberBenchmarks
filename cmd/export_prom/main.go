package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	total_http_request_counter = "total_http_request_counter"
	memory_allocations_bytes   = "memory_allocations_bytes"
	goroutines_count           = "goroutines_count"

	// timestamp секнды загрузка процессора в процентах
	container_cpu_usage_seconds_total = "rate(container_cpu_usage_seconds_total{name=\"deployment-gin-app-1\"}[1m])*100"
	// timestamp байты (екущее использование памяти в байтах
	container_memory_usage_bytes = "container_memory_usage_bytes{name=\"deployment-gin-app-1\"}"
	// container_fs_reads_bytes_total = "container_fs_reads_bytes_total{name=\"deployment-gin-app-1\"} "
)

var save_json = true

var (
	metrics_save = []string{
		total_http_request_counter, memory_allocations_bytes, goroutines_count,
		container_cpu_usage_seconds_total, container_memory_usage_bytes}
)

type PrometheusExporter struct {
	prometheusURL string
	startTime     string
	endTime       string
	step          string
	maxWorkers    int
	client        *http.Client
}

type PrometheusLabelResponse struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

type PrometheusQueryRangeResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Values [][]interface{}   `json:"values"`
		} `json:"result"`
	} `json:"data"`
	Error     string `json:"error"`
	ErrorType string `json:"errorType"`
}

type ExportInfo struct {
	ExportInfo struct {
		PrometheusURL         string `json:"prometheus_url"`
		StartTime             string `json:"start_time"`
		EndTime               string `json:"end_time"`
		Step                  string `json:"step"`
		ExportTime            string `json:"export_time"`
		TotalMetricsAvailable int    `json:"total_metrics_available"`
		TotalMetricsExported  int    `json:"total_metrics_exported"`
	} `json:"export_info"`
	MetricGroups map[string]interface{} `json:"metric_groups,omitempty"`
}

type ExportResult struct {
	MetricName  string
	Message     string
	Success     bool
	HasData     bool
	SeriesCount int
}

func NewPrometheusExporter(prometheusURL, startTime, endTime, step string, maxWorkers int) *PrometheusExporter {
	return &PrometheusExporter{
		prometheusURL: strings.TrimRight(prometheusURL, "/"),
		startTime:     startTime,
		endTime:       endTime,
		step:          step,
		maxWorkers:    maxWorkers,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (p *PrometheusExporter) GetAllMetrics() ([]string, error) {
	log.Println("📋 Fetching list of all metrics...")

	url := fmt.Sprintf("%s/api/v1/label/__name__/values", p.prometheusURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch metrics list: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var labelResponse PrometheusLabelResponse
	if err := json.Unmarshal(body, &labelResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if labelResponse.Status != "success" {
		return nil, fmt.Errorf("prometheus API error: status %s", labelResponse.Status)
	}

	metrics := labelResponse.Data
	sort.Strings(metrics)

	log.Printf("✅ Found %d metrics", len(metrics))
	return metrics, nil
}

func (p *PrometheusExporter) SanitizeFilename(filename string) string {
	// Заменяем недопустимые символы в именах файлов
	reg := regexp.MustCompile(`[<>:"/\\|?*]`)
	safeName := reg.ReplaceAllString(filename, "_")

	// Убираем множественные подчеркивания
	reg = regexp.MustCompile(`_{2,}`)
	safeName = reg.ReplaceAllString(safeName, "_")

	return strings.Trim(safeName, "_")
}

type metric struct {
	Timestamp int
	Val       float64
}

func parseFloatVals(vals [][]interface{}) []metric {
	resultVals := []metric{}
	for _, v := range vals {
		timestamp := int(v[0].(float64))

		// metricVal, err := strconv.Atoi(v[1].(string))
		metricVal, err := strconv.ParseFloat(v[1].(string), 64)
		if err != nil {
			panic(err)
		}
		resultVals = append(resultVals, metric{Timestamp: timestamp, Val: metricVal})
	}
	return resultVals
}

// func parseValues(vals [][]interface{}) []metric {
// 	resultVals := [][]int{}
// 	for _, v := range vals {
// 		timestamp := int(v[0].(float64))

// 		metricVal, err := strconv.Atoi(v[1].(string))
// 		if err != nil {
// 			panic(err)
// 		}
// 		resultVals = append(resultVals, metric{Timestamp: timestamp, Val: metricVal})
// 	}
// 	return resultVals
// }

func saveMetrics(queryResponse PrometheusQueryRangeResponse, metricName string) error {
	var resultVals []metric
	if metricName == memory_allocations_bytes {
		for i := range 2 {
			filename := fmt.Sprintf("./metrics_data/prometheus/%s_%s.txt", metricName, queryResponse.Data.Result[i].Metric["type"])
			fmt.Printf("filename:= %s\n", filename)

			file, err := os.Create(filename)
			if err != nil {
				return fmt.Errorf("❌ Failed to create file for %s: %v", metricName, err)
			}
			defer file.Close()

			vals := queryResponse.Data.Result[i].Values
			resultVals = parseFloatVals(vals)
			for _, v := range resultVals {
				file.WriteString(fmt.Sprintf("%d %.2f\n", v.Timestamp, v.Val))
			}
		}
	} else {
		filename := fmt.Sprintf("./metrics_data/prometheus/%s.txt", metricName)
		file, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("❌ Failed to create file for %s: %v", metricName, err)
		}
		defer file.Close()

		vals := queryResponse.Data.Result[0].Values
		resultVals = parseFloatVals(vals)
		for _, v := range resultVals {
			file.WriteString(fmt.Sprintf("%d %.2f\n", v.Timestamp, v.Val))
		}
	}
	return nil
}

func (p *PrometheusExporter) ExportMetric(metricName string) *ExportResult {
	result := &ExportResult{
		MetricName: metricName,
	}
	// Подготавливаем параметры запроса
	params := url.Values{}
	params.Add("query", metricName)
	params.Add("step", p.step)

	// Устанавливаем временной диапазон
	if p.startTime != "" && p.endTime != "" {
		params.Add("start", p.startTime)
		params.Add("end", p.endTime)
	} else {
		endTime := time.Now()
		startTime := endTime.Add(-3*time.Hour - 3*time.Minute)
		params.Add("start", startTime.Format(time.RFC3339))
		params.Add("end", endTime.Format(time.RFC3339))
	}

	// Выполняем запрос
	queryURL := fmt.Sprintf("%s/api/v1/query_range?%s", p.prometheusURL, params.Encode())
	req, err := http.NewRequest("GET", queryURL, nil)
	if err != nil {
		result.Message = fmt.Sprintf("❌ Failed to create request for %s: %v", metricName, err)
		return result
	}

	resp, err := p.client.Do(req)
	if err != nil {
		result.Message = fmt.Sprintf("❌ Network error for %s: %v", metricName, err)
		return result
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		result.Message = fmt.Sprintf("❌ Failed to read response for %s: %v", metricName, err)
		return result
	}

	var queryResponse PrometheusQueryRangeResponse
	if err := json.Unmarshal(body, &queryResponse); err != nil {
		result.Message = fmt.Sprintf("❌ Failed to parse response for %s: %v", metricName, err)
		return result
	}

	if queryResponse.Status != "success" {
		result.Message = fmt.Sprintf("❌ API error for %s: %s", metricName, queryResponse.Error)
		return result
	}

	// Проверяем наличие данных
	if len(queryResponse.Data.Result) == 0 {
		result.Message = fmt.Sprintf("⚠️  No data for metric: %s", metricName)
		result.HasData = false
		return result
	}

	// Сохраняем в файл json
	if save_json {
		file, err := os.Create(fmt.Sprintf("./prometheus-export/%s.json", metricName))
		if err != nil {
			result.Message = fmt.Sprintf("❌ Failed to create file for %s: %v", metricName, err)
			return result
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(queryResponse); err != nil {
			result.Message = fmt.Sprintf("❌ Failed to write file for %s: %v", metricName, err)
			return result
		}
	}
	err = saveMetrics(queryResponse, metricName)
	if err != nil {
		result.Message = fmt.Sprintf("Error: %v", err)
		return result
	}
	result.Success = true
	result.HasData = true
	result.SeriesCount = len(queryResponse.Data.Result)
	result.Message = fmt.Sprintf("✅ %s (%d series)", metricName, result.SeriesCount)

	return result
}

func (p *PrometheusExporter) ExportAllMetrics() error {
	all_metrics, err := p.GetAllMetrics()
	if err != nil {
		return fmt.Errorf("failed to get metrics list: %w", err)
	}
	fmt.Printf("all metrics: %v\n", all_metrics)
	metrics := metrics_save
	fmt.Printf("metrics save: %v\n", metrics)
	log.Printf("🚀 Starting export of %d metrics with %d workers...", len(metrics), p.maxWorkers)

	results := p.exportMetricsParallel(metrics)
	// Статистика
	successCount := 0
	noDataCount := 0
	errorCount := 0

	for _, result := range results {
		log.Println(result.Message)

		if result.Success {
			successCount++
		} else if !result.HasData && strings.Contains(result.Message, "No data") {
			noDataCount++
		} else {
			errorCount++
		}
	}

	log.Printf("\n📊 Export completed:")
	log.Printf("   ✅ Success: %d", successCount)
	log.Printf("   ⚠️  No data: %d", noDataCount)
	log.Printf("   ❌ Errors:  %d", errorCount)
	log.Printf("   Time:  %s - %s", p.startTime, p.endTime)
	return nil
}

func (p *PrometheusExporter) exportMetricsParallel(metrics []string) []*ExportResult {
	var mu sync.Mutex
	results := make([]*ExportResult, 0, len(metrics))

	g := new(errgroup.Group)
	g.SetLimit(p.maxWorkers)

	for _, metric := range metrics {
		metric := metric // capture range variable
		g.Go(func() error {
			result := p.ExportMetric(metric)

			mu.Lock()
			results = append(results, result)
			mu.Unlock()

			return nil
		})
	}

	// Ждем завершения всех горутин
	if err := g.Wait(); err != nil {
		log.Printf("Error in worker: %v", err)
	}

	return results
}

// go run ./cmd/export_prom/main.go -start="1760796858" -end="1760796896"
func main() {
	var (
		prometheusURL = flag.String("url", "http://localhost:9090", "Prometheus URL")
		startTime     = flag.String("start", "0", "Start time (RFC3339 or Unix timestamp)")
		endTime       = flag.String("end", "0", "End time (RFC3339 or Unix timestamp)")
		step          = flag.String("step", "1s", "Query step duration")
		workers       = flag.Int("workers", 5, "Number of parallel workers")
	)
	flag.Parse()

	log.Println("🚀 Prometheus Metrics Exporter")
	log.Println(strings.Repeat("=", 50))

	exporter := NewPrometheusExporter(
		*prometheusURL,
		*startTime,
		*endTime,
		*step,
		*workers,
	)

	if err := exporter.ExportAllMetrics(); err != nil {
		log.Fatalf("❌ Export failed: %v", err)
	}
}
