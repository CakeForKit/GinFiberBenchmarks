package analyzer

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/models"
)

type ResultAnalyzer struct {
	resultsDir string
	reportsDir string
}

func NewResultAnalyzer(resultsDir string) *ResultAnalyzer {
	reportsDir := filepath.Join(resultsDir, "reports")
	return &ResultAnalyzer{
		resultsDir: resultsDir,
		reportsDir: reportsDir,
	}
}

// LoadResults загружает результаты для фреймворка
func (a *ResultAnalyzer) LoadResults(framework string) (*models.FrameworkResults, error) {
	pattern := filepath.Join(a.resultsDir, "raw", framework+"_*.json")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	if len(matches) == 0 {
		return nil, fmt.Errorf("no results found for %s", framework)
	}

	// Берем последний файл
	sort.Sort(sort.Reverse(sort.StringSlice(matches)))
	latestFile := matches[0]

	data, err := os.ReadFile(latestFile)
	if err != nil {
		return nil, err
	}

	var results models.FrameworkResults
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

// CalculateResourceStats вычисляет статистику по ресурсам
func (a *ResultAnalyzer) CalculateResourceStats(results *models.FrameworkResults) *models.ResourceUsage {
	// В реальной реализации здесь будут данные из Prometheus
	// Заглушка для демонстрации
	usage := &models.ResourceUsage{}
	usage.CPU.Min = 0.1
	usage.CPU.Max = 85.5
	usage.CPU.Median = 45.2
	usage.CPU.P95 = 78.3

	usage.Memory.Min = 50.2
	usage.Memory.Max = 256.8
	usage.Memory.Median = 120.4
	usage.Memory.P95 = 230.1

	return usage
}

// GenerateComparisonReport генерирует сравнительный отчет
func (a *ResultAnalyzer) GenerateComparisonReport(ginResults, fiberResults *models.FrameworkResults) ([]models.ComparisonResult, error) {
	var comparison []models.ComparisonResult

	endpoints := []string{"simple", "tree", "wide"}

	for _, endpoint := range endpoints {
		ginStats := ginResults.Endpoints[endpoint].Measurements
		fiberStats := fiberResults.Endpoints[endpoint].Measurements

		performanceDiff := 0.0
		if ginStats.P50 > 0 {
			performanceDiff = (fiberStats.P50 - ginStats.P50) / ginStats.P50 * 100
		}

		comparison = append(comparison, models.ComparisonResult{
			Endpoint:         endpoint,
			GinP50:           ginStats.P50,
			FiberP50:         fiberStats.P50,
			GinP95:           ginStats.P95,
			FiberP95:         fiberStats.P95,
			GinP99:           ginStats.P99,
			FiberP99:         fiberStats.P99,
			GinSuccessRate:   ginStats.SuccessRate,
			FiberSuccessRate: fiberStats.SuccessRate,
			PerformanceDiff:  performanceDiff,
		})
	}

	return comparison, nil
}

// SaveComparisonCSV сохраняет сравнение в CSV
func (a *ResultAnalyzer) SaveComparisonCSV(comparison []models.ComparisonResult) error {
	if err := os.MkdirAll(a.reportsDir, 0755); err != nil {
		return err
	}

	csvPath := filepath.Join(a.reportsDir, "comparison_report.csv")
	file, err := os.Create(csvPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Заголовок
	headers := []string{
		"endpoint", "gin_p50", "fiber_p50", "gin_p95", "fiber_p95",
		"gin_p99", "fiber_p99", "gin_success_rate", "fiber_success_rate",
		"performance_diff_p50",
	}
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Данные
	for _, result := range comparison {
		record := []string{
			result.Endpoint,
			fmt.Sprintf("%.6f", result.GinP50),
			fmt.Sprintf("%.6f", result.FiberP50),
			fmt.Sprintf("%.6f", result.GinP95),
			fmt.Sprintf("%.6f", result.FiberP95),
			fmt.Sprintf("%.6f", result.GinP99),
			fmt.Sprintf("%.6f", result.FiberP99),
			fmt.Sprintf("%.6f", result.GinSuccessRate),
			fmt.Sprintf("%.6f", result.FiberSuccessRate),
			fmt.Sprintf("%.6f", result.PerformanceDiff),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	log.Printf("Comparison report saved to: %s", csvPath)
	return nil
}

// SaveResourceUsage сохраняет использование ресурсов в JSON
func (a *ResultAnalyzer) SaveResourceUsage(ginUsage, fiberUsage *models.ResourceUsage) error {
	resourceReport := map[string]*models.ResourceUsage{
		"gin":   ginUsage,
		"fiber": fiberUsage,
	}

	jsonPath := filepath.Join(a.reportsDir, "resource_usage.json")
	data, err := json.MarshalIndent(resourceReport, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(jsonPath, data, 0644); err != nil {
		return err
	}

	log.Printf("Resource usage report saved to: %s", jsonPath)
	return nil
}

// GenerateCharts генерирует графики (заглушка - в реальности используйте gonum/plot или другую библиотеку)
func (a *ResultAnalyzer) GenerateCharts(ginResults, fiberResults *models.FrameworkResults) error {
	log.Printf("Generating charts...")
	// В реальной реализации здесь будет код для создания графиков
	// с использованием gonum/plot, go-chart или другой библиотеки

	log.Printf("Charts would be generated in: %s", a.reportsDir)
	return nil
}

// GenerateFullReport генерирует полный отчет
func (a *ResultAnalyzer) GenerateFullReport() error {
	if err := os.MkdirAll(a.reportsDir, 0755); err != nil {
		return err
	}

	log.Printf("Loading results...")
	ginResults, err := a.LoadResults("gin")
	if err != nil {
		return fmt.Errorf("failed to load gin results: %v", err)
	}

	fiberResults, err := a.LoadResults("fiber")
	if err != nil {
		return fmt.Errorf("failed to load fiber results: %v", err)
	}

	log.Printf("Generating comparison report...")
	comparison, err := a.GenerateComparisonReport(ginResults, fiberResults)
	if err != nil {
		return err
	}

	if err := a.SaveComparisonCSV(comparison); err != nil {
		return err
	}

	log.Printf("Calculating resource usage...")
	ginUsage := a.CalculateResourceStats(ginResults)
	fiberUsage := a.CalculateResourceStats(fiberResults)

	if err := a.SaveResourceUsage(ginUsage, fiberUsage); err != nil {
		return err
	}

	log.Printf("Generating charts...")
	if err := a.GenerateCharts(ginResults, fiberResults); err != nil {
		log.Printf("Warning: chart generation failed: %v", err)
	}

	log.Printf("Full report generated in: %s", a.reportsDir)

	// Вывод краткого отчета в консоль
	a.PrintSummary(comparison)

	return nil
}

// PrintSummary выводит краткий отчет в консоль
func (a *ResultAnalyzer) PrintSummary(comparison []models.ComparisonResult) {
	fmt.Printf("\n%s\n", strings.Repeat("=", 60))
	fmt.Printf("BENCHMARK SUMMARY\n")
	fmt.Printf("%s\n", strings.Repeat("=", 60))

	for _, result := range comparison {
		fmt.Printf("\nEndpoint: %s\n", result.Endpoint)
		fmt.Printf("  P50:    Gin=%.3fms, Fiber=%.3fms (diff=%.2f%%)\n",
			result.GinP50, result.FiberP50, result.PerformanceDiff)
		fmt.Printf("  P95:    Gin=%.3fms, Fiber=%.3fms\n",
			result.GinP95, result.FiberP95)
		fmt.Printf("  P99:    Gin=%.3fms, Fiber=%.3fms\n",
			result.GinP99, result.FiberP99)
		fmt.Printf("  Success: Gin=%.1f%%, Fiber=%.1f%%\n",
			result.GinSuccessRate*100, result.FiberSuccessRate*100)
	}
}
