package logmetrics

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/vmihailenco/msgpack/v5"
)

func SaveStat(fileNameToSave string, metrics []SerializeMetric) error {
	fmt.Printf("fileNameToSave: %s\n\n", fileNameToSave)
	sorted := make([]SerializeMetric, len(metrics))
	copy(sorted, metrics)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].SerializeStartTime.Before(sorted[j].SerializeStartTime)
	})

	f, err := os.Create(fileNameToSave)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	if _, err := f.WriteString("timeStart(ms) duration(us)\n"); err != nil {
		return err
	}

	startTimePoint := sorted[0].SerializeStartTime
	for _, v := range sorted {
		timeStart := v.SerializeStartTime.Sub(startTimePoint)
		duration := v.SerializeEndTime.Sub(v.SerializeStartTime)
		if _, err := f.WriteString(fmt.Sprintf("%d %d\n", timeStart.Milliseconds(), duration.Microseconds())); err != nil {
			return err
		}
	}
	return nil
}

func rReadMsgpackWithTime(filename string) ([]SerializeMetric, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data []SerializeMetric
	err = msgpack.Unmarshal(bytes, &data)
	return data, err
}

func aAnalyzeLogs(metrics []SerializeMetric) {
	if len(metrics) == 0 {
		fmt.Println("Нет данных для анализа")
		return
	}

	// Собираем все durations
	var durations []time.Duration
	for _, m := range metrics {
		durations = append(durations, m.SerializeEndTime.Sub(m.SerializeStartTime))
	}

	// Расчет перцентилей
	percentiles := []float64{0.5, 0.75, 0.9, 0.95, 0.99}
	percentileResults := calculatePercentiles(durations, percentiles)

	// Базовая статистика
	var total time.Duration
	min, max := durations[0], durations[0]

	for _, d := range durations {
		total += d
		if d < min {
			min = d
		}
		if d > max {
			max = d
		}
	}

	average := total / time.Duration(len(durations))

	// Вывод результатов
	fmt.Println("=== АНАЛИЗ ВРЕМЕНИ СЕРИАЛИЗАЦИИ ===")
	fmt.Print("µs = 1e-6 s\n")
	fmt.Printf("Общее количество запросов: %d\n", len(metrics))
	fmt.Printf("Минимальное время: %v\n", min)
	fmt.Printf("Максимальное время: %v\n", max)
	fmt.Printf("Среднее время: %v\n", average)

	fmt.Println("\nПерцентили:")
	for _, p := range percentiles {
		fmt.Printf("P%.0f: %v\n", p*100, percentileResults[p])
	}

}

func calculatePercentiles(durations []time.Duration, percentiles []float64) map[float64]time.Duration {
	// Сортируем durations
	sorted := make([]time.Duration, len(durations))
	copy(sorted, durations)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	result := make(map[float64]time.Duration)
	for _, p := range percentiles {
		if len(sorted) == 0 {
			result[p] = 0
			continue
		}

		index := int(float64(len(sorted)) * p)
		if index >= len(sorted) {
			index = len(sorted) - 1
		}
		result[p] = sorted[index]
	}

	return result
}

/*
func printHistogram(metrics []SerializeMetric) {
	if len(metrics) == 0 {
		return
	}
	sorted := make([]SerializeMetric, len(metrics))
	copy(sorted, metrics)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].SerializeStartTime.Before(sorted[j].SerializeStartTime)
	})

	// Находим min и max для создания интервалов
	min, max := metrics[0].Duration, metrics[0].Duration
	for _, m := range metrics {
		if m.Duration < min {
			min = m.Duration
		}
		if m.Duration > max {
			max = m.Duration
		}
	}

	// Создаем интервалы (10 интервалов)
	bucketCount := 10
	bucketSize := (max - min) / time.Duration(bucketCount)

	buckets := make([]int, bucketCount)
	for _, m := range metrics {
		bucket := int((m.Duration - min) / bucketSize)
		if bucket >= bucketCount {
			bucket = bucketCount - 1
		}
		buckets[bucket]++
	}

	fmt.Println("\n=== ГИСТОГРАММА РАСПРЕДЕЛЕНИЯ ===")
	for i, count := range buckets {
		start := min + time.Duration(i)*bucketSize
		end := min + time.Duration(i+1)*bucketSize
		bar := strings.Repeat("█", count*50/len(metrics))
		fmt.Printf("%v - %v: %s %d запросов\n",
			start, end, bar, count)
	}
}

*/
