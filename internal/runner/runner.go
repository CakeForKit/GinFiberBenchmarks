package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/cnfg"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/models"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/utils"
)

type BenchmarkRunner struct {
	config *cnfg.Config
	client *http.Client
}

func NewBenchmarkRunner(config *cnfg.Config) *BenchmarkRunner {
	if config == nil {
		config = cnfg.NewDefaultConfig()
	}

	return &BenchmarkRunner{
		config: config,
		client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,              // controls the maximum number of idle (keep-alive) connections across all hosts
				MaxIdleConnsPerHost: 100,              // if non-zero, controls the maximum idle (keep-alive) connections to keep per-host
				IdleConnTimeout:     90 * time.Second, // the maximum amount of time an idle (keep-alive) connection will remain idle before closing itself
			},
		},
	}
}

// WarmUp выполняет прогрев приложения
func (b *BenchmarkRunner) WarmUp(url string) error {
	for i := 0; i < b.config.WarmupRequests; i++ {
		resp, err := b.client.Get(url + "/health")
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()
			return nil
		}
		if resp != nil {
			resp.Body.Close()
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("warmup failed for %s", url)
}

// SingleRequest выполняет один запрос и измеряет время
func (b *BenchmarkRunner) SingleRequest(url, endpoint string) models.RequestResult {
	start := time.Now()

	fullURL := fmt.Sprintf("%s/%s", url, endpoint)
	resp, err := b.client.Get(fullURL)
	duration := utils.DurationToMillis(time.Since(start))

	result := models.RequestResult{
		Duration:  duration,
		Timestamp: utils.Now(),
	}

	if err != nil {
		result.Error = err.Error()
		result.Success = false
		return result
	}
	defer resp.Body.Close()

	// Читаем тело ответа чтобы гарантировать полную передачу данных
	_, readErr := io.Copy(io.Discard, resp.Body)
	if readErr != nil {
		result.Error = readErr.Error()
		result.Success = false
		return result
	}

	result.Status = resp.StatusCode
	result.Success = resp.StatusCode == http.StatusOK

	return result
}

// LoadTest выполняет нагрузочное тестирование
func (b *BenchmarkRunner) LoadTest(ctx context.Context, url, endpoint string, concurrentRequests int, duration time.Duration) []models.RequestResult {
	var (
		results []models.RequestResult
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	endTime := time.Now().Add(duration)

	// Горутина для отправки запросов
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if time.Now().After(endTime) {
					return
				}

				// Отправляем batch запросов
				batchWG := sync.WaitGroup{}
				batchResults := make([]models.RequestResult, concurrentRequests)

				for i := 0; i < concurrentRequests; i++ {
					batchWG.Add(1)
					go func(index int) {
						defer batchWG.Done()
						batchResults[index] = b.SingleRequest(url, endpoint)
					}(i)
				}
				batchWG.Wait()

				mu.Lock()
				results = append(results, batchResults...)
				mu.Unlock()
			}
		}
	}()

	wg.Wait()
	return results
}

// DegradationTest выполняет тест на поиск точки деградации
func (b *BenchmarkRunner) DegradationTest(ctx context.Context, url, endpoint string) []models.DegradationResult {
	log.Printf("Running degradation test for %s/%s", url, endpoint)

	var results []models.DegradationResult

	for _, concurrent := range b.config.ConcurrentLevels {
		log.Printf("Testing with %d concurrent requests", concurrent)

		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		requestResults := b.LoadTest(ctx, url, endpoint, concurrent, 10*time.Second)
		cancel()

		successful := utils.FilterSuccessful(requestResults)
		if len(successful) == 0 {
			break
		}

		stats := utils.CalculateStats(requestResults)

		result := models.DegradationResult{
			Concurrent:         concurrent,
			TotalRequests:      stats.TotalRequests,
			SuccessfulRequests: stats.SuccessfulRequests,
			SuccessRate:        stats.SuccessRate,
			AvgDuration:        stats.AvgDuration,
			P50:                stats.P50,
			P75:                stats.P75,
			P90:                stats.P90,
			P95:                stats.P95,
			P99:                stats.P99,
			MaxDuration:        stats.MaxDuration,
		}

		results = append(results, result)

		// Останавливаемся если успешность ниже 95%
		if result.SuccessRate < 0.95 {
			break
		}
	}

	return results
}

// OverloadRecoveryTest выполняет тест на перегруз и восстановление
func (b *BenchmarkRunner) OverloadRecoveryTest(ctx context.Context, url, endpoint string) *models.OverloadRecoveryResult {
	log.Printf("Running overload recovery test for %s/%s", url, endpoint)

	// Фаза перегрузки
	log.Printf("Phase 1: Overload (300 concurrent requests for 30s)")
	ctx1, cancel1 := context.WithTimeout(ctx, 30*time.Second)
	overloadResults := b.LoadTest(ctx1, url, endpoint, 300, 30*time.Second)
	cancel1()

	// Пауза для восстановления
	log.Printf("Recovery phase: waiting 10s")
	time.Sleep(10 * time.Second)

	// Фаза проверки восстановления
	log.Printf("Phase 2: Recovery test (10 concurrent requests for 10s)")
	ctx2, cancel2 := context.WithTimeout(ctx, 10*time.Second)
	recoveryResults := b.LoadTest(ctx2, url, endpoint, 10, 10*time.Second)
	cancel2()

	return &models.OverloadRecoveryResult{
		Overload: utils.CalculateStats(overloadResults),
		Recovery: utils.CalculateStats(recoveryResults),
	}
}

// RunSingleExperiment запускает один полный эксперимент
func (b *BenchmarkRunner) RunSingleExperiment(ctx context.Context, framework, url string) *models.FrameworkResults {
	log.Printf("Starting experiment for %s", framework)

	endpoints := []string{"simple", "tree", "wide"}
	results := &models.FrameworkResults{
		Framework: framework,
		Endpoints: make(map[string]*models.EndpointResults),
		Timestamp: utils.Now(),
	}

	for _, endpoint := range endpoints {
		log.Printf("Testing endpoint: %s", endpoint)

		// Прогрев
		if err := b.WarmUp(url); err != nil {
			log.Printf("Warning: warmup failed: %v", err)
		}

		// 100 измерений для статистики
		log.Printf("Running %d measurements for statistics", b.config.Measurements)
		var allMeasurements []models.RequestResult

		for i := 0; i < b.config.Measurements; i++ {
			measurementCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
			measurementResults := b.LoadTest(measurementCtx, url, endpoint, 10, 2*time.Second)
			cancel()

			allMeasurements = append(allMeasurements, measurementResults...)
			log.Printf("Measurement %d/%d completed", i+1, b.config.Measurements)
		}

		// Тест деградации
		degradationResults := b.DegradationTest(ctx, url, endpoint)

		// Тест перегрузки
		overloadResults := b.OverloadRecoveryTest(ctx, url, endpoint)

		results.Endpoints[endpoint] = &models.EndpointResults{
			Measurements:     utils.CalculateStats(allMeasurements),
			Degradation:      degradationResults,
			OverloadRecovery: overloadResults,
			RawMeasurements:  allMeasurements[:min(1000, len(allMeasurements))], // Сохраняем сырые данные для графиков
		}
	}

	return results
}

// SaveResults сохраняет результаты в файлы
func (b *BenchmarkRunner) SaveResults(results *models.FrameworkResults) (string, error) {
	timestamp := results.Timestamp.Format("20060102_150405")

	// Создаем директории
	rawDir := filepath.Join(b.config.ResultsDir, "raw")
	if err := os.MkdirAll(rawDir, 0755); err != nil {
		return "", err
	}

	// Сохраняем JSON
	jsonPath := filepath.Join(rawDir, fmt.Sprintf("%s_%s.json", results.Framework, timestamp))
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return "", err
	}

	if err := os.WriteFile(jsonPath, jsonData, 0644); err != nil {
		return "", err
	}

	// Сохраняем CSV для сырых измерений
	for endpoint, data := range results.Endpoints {
		csvPath := filepath.Join(rawDir, fmt.Sprintf("%s_%s_%s.csv", results.Framework, endpoint, timestamp))
		if err := b.saveRawMeasurementsToCSV(csvPath, data.RawMeasurements); err != nil {
			log.Printf("Warning: failed to save CSV for %s: %v", endpoint, err)
		}
	}

	return jsonPath, nil
}

// saveRawMeasurementsToCSV сохраняет сырые измерения в CSV
func (b *BenchmarkRunner) saveRawMeasurementsToCSV(path string, measurements []models.RequestResult) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Заголовок CSV
	file.WriteString("timestamp,duration,success,status\n")

	for _, m := range measurements {
		file.WriteString(fmt.Sprintf("%s,%.6f,%t,%d\n",
			m.Timestamp.Format(time.RFC3339),
			m.Duration,
			m.Success,
			m.Status,
		))
	}

	return nil
}

// RunFullBenchmark запускает полный бенчмарк для обоих фреймворков
func (b *BenchmarkRunner) RunFullBenchmark(ctx context.Context) (map[string]string, error) {
	frameworks := map[string]string{

		"gin":   fmt.Sprintf("http://%s:%d", b.config.Host, b.config.GinPort),
		"fiber": fmt.Sprintf("http://%s:%d", b.config.Host, b.config.FiberPort),
	}

	results := make(map[string]string)

	for name, url := range frameworks {
		log.Printf("\n%s", "==================================================")
		log.Printf("BENCHMARKING %s", name)
		log.Printf("==================================================")

		frameworkResults := b.RunSingleExperiment(ctx, name, url)
		resultsPath, err := b.SaveResults(frameworkResults)
		if err != nil {
			return nil, fmt.Errorf("failed to save results for %s: %v", name, err)
		}

		results[name] = resultsPath
		log.Printf("Results saved to: %s", resultsPath)
	}

	return results, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
