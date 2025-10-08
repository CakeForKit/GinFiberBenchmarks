package models

import (
	"encoding/json"
	"time"
)

// RequestResult результат одного запроса
type RequestResult struct {
	Duration  float64   `json:"duration"` // в миллисекундах
	Status    int       `json:"status"`
	Success   bool      `json:"success"`
	Error     string    `json:"error,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

// EndpointStats статистика по endpoint
type EndpointStats struct {
	TotalRequests      int     `json:"total_requests"`
	SuccessfulRequests int     `json:"successful_requests"`
	SuccessRate        float64 `json:"success_rate"`
	AvgDuration        float64 `json:"avg_duration"`
	P50                float64 `json:"p50"`
	P75                float64 `json:"p75"`
	P90                float64 `json:"p90"`
	P95                float64 `json:"p95"`
	P99                float64 `json:"p99"`
	MinDuration        float64 `json:"min_duration"`
	MaxDuration        float64 `json:"max_duration"`
}

// DegradationResult результаты теста деградации
type DegradationResult struct {
	Concurrent         int     `json:"concurrent"`
	TotalRequests      int     `json:"total_requests"`
	SuccessfulRequests int     `json:"successful_requests"`
	SuccessRate        float64 `json:"success_rate"`
	AvgDuration        float64 `json:"avg_duration"`
	P50                float64 `json:"p50"`
	P75                float64 `json:"p75"`
	P90                float64 `json:"p90"`
	P95                float64 `json:"p95"`
	P99                float64 `json:"p99"`
	MaxDuration        float64 `json:"max_duration"`
}

// OverloadRecoveryResult результаты теста перегрузки
type OverloadRecoveryResult struct {
	Overload *EndpointStats `json:"overload"`
	Recovery *EndpointStats `json:"recovery"`
}

// EndpointResults все результаты по endpoint
type EndpointResults struct {
	Measurements     *EndpointStats          `json:"measurements"`
	Degradation      []DegradationResult     `json:"degradation"`
	OverloadRecovery *OverloadRecoveryResult `json:"overload_recovery"`
	RawMeasurements  []RequestResult         `json:"raw_measurements"`
}

// FrameworkResults результаты по фреймворку
type FrameworkResults struct {
	Framework string                      `json:"framework"`
	Endpoints map[string]*EndpointResults `json:"endpoints"`
	Timestamp time.Time                   `json:"timestamp"`
}

// ResourceUsage использование ресурсов
type ResourceUsage struct {
	CPU struct {
		Min    float64 `json:"min"`
		Max    float64 `json:"max"`
		Median float64 `json:"median"`
		P95    float64 `json:"p95"`
	} `json:"cpu"`
	Memory struct {
		Min    float64 `json:"min"`
		Max    float64 `json:"max"`
		Median float64 `json:"median"`
		P95    float64 `json:"p95"`
	} `json:"memory"`
}

// ComparisonResult результаты сравнения
type ComparisonResult struct {
	Endpoint         string  `json:"endpoint" csv:"endpoint"`
	GinP50           float64 `json:"gin_p50" csv:"gin_p50"`
	FiberP50         float64 `json:"fiber_p50" csv:"fiber_p50"`
	GinP95           float64 `json:"gin_p95" csv:"gin_p95"`
	FiberP95         float64 `json:"fiber_p95" csv:"fiber_p95"`
	GinP99           float64 `json:"gin_p99" csv:"gin_p99"`
	FiberP99         float64 `json:"fiber_p99" csv:"fiber_p99"`
	GinSuccessRate   float64 `json:"gin_success_rate" csv:"gin_success_rate"`
	FiberSuccessRate float64 `json:"fiber_success_rate" csv:"fiber_success_rate"`
	PerformanceDiff  float64 `json:"performance_diff_p50" csv:"performance_diff_p50"`
}

// MarshalJSON кастомная сериализация для RequestResult
func (r RequestResult) MarshalJSON() ([]byte, error) {
	type Alias RequestResult
	return json.Marshal(&struct {
		Timestamp string `json:"timestamp"`
		Alias
	}{
		Timestamp: r.Timestamp.Format(time.RFC3339),
		Alias:     (Alias)(r),
	})
}

// UnmarshalJSON кастомная десериализация для RequestResult
func (r *RequestResult) UnmarshalJSON(data []byte) error {
	type Alias RequestResult
	aux := &struct {
		Timestamp string `json:"timestamp"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Timestamp != "" {
		t, err := time.Parse(time.RFC3339, aux.Timestamp)
		if err != nil {
			return err
		}
		r.Timestamp = t
	}
	return nil
}
