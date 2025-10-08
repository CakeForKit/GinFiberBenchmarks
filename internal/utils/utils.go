package utils

import (
	"math"
	"sort"
	"time"

	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/models"
)

// CalculateStats вычисляет статистику по результатам запросов
func CalculateStats(results []models.RequestResult) *models.EndpointStats {
	if len(results) == 0 {
		return &models.EndpointStats{}
	}

	successful := FilterSuccessful(results)
	if len(successful) == 0 {
		return &models.EndpointStats{
			TotalRequests: len(results),
			SuccessRate:   0,
		}
	}

	durations := make([]float64, len(successful))
	for i, r := range successful {
		durations[i] = r.Duration
	}

	sort.Float64s(durations)

	return &models.EndpointStats{
		TotalRequests:      len(results),
		SuccessfulRequests: len(successful),
		SuccessRate:        float64(len(successful)) / float64(len(results)),
		AvgDuration:        calculateAverage(durations),
		P50:                calculatePercentile(durations, 50),
		P75:                calculatePercentile(durations, 75),
		P90:                calculatePercentile(durations, 90),
		P95:                calculatePercentile(durations, 95),
		P99:                calculatePercentile(durations, 99),
		MinDuration:        durations[0],
		MaxDuration:        durations[len(durations)-1],
	}
}

// FilterSuccessful фильтрует успешные запросы
func FilterSuccessful(results []models.RequestResult) []models.RequestResult {
	var successful []models.RequestResult
	for _, r := range results {
		if r.Success {
			successful = append(successful, r)
		}
	}
	return successful
}

// calculateAverage вычисляет среднее значение
func calculateAverage(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// calculatePercentile вычисляет перцентиль
func calculatePercentile(sortedValues []float64, percentile float64) float64 {
	if len(sortedValues) == 0 {
		return 0
	}

	index := (percentile / 100) * float64(len(sortedValues)-1)

	if index == float64(int64(index)) {
		// Целый индекс
		return sortedValues[int(index)]
	}

	// Интерполяция между двумя ближайшими значениями
	lower := int(math.Floor(index))
	upper := int(math.Ceil(index))
	weight := index - float64(lower)

	return sortedValues[lower]*(1-weight) + sortedValues[upper]*weight
}

// Now возвращает текущее время в UTC
func Now() time.Time {
	return time.Now().UTC()
}

// DurationToMillis преобразует duration в миллисекунды
func DurationToMillis(duration time.Duration) float64 {
	return float64(duration.Nanoseconds()) / 1e6
}
