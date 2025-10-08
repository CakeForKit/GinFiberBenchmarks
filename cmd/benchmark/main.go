package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/analyzer"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/cnfg"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/runner"
)

func main() {
	analyzeOnly := false // flag.Bool("analyze", false, "Only analyze existing results")
	config := cnfg.NewDefaultConfig()

	// Создание директорий если не существуют
	if err := os.MkdirAll(config.ResultsDir, 0755); err != nil {
		log.Fatalf("Failed to create results directory: %v", err)
	}

	// Обработка сигналов для graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("Received shutdown signal...")
		cancel()
	}()

	if analyzeOnly {
		// Только анализ существующих результатов
		if err := runAnalysis(config.ResultsDir); err != nil {
			log.Fatalf("Analysis failed: %v", err)
		}
	} else {
		// Полный прогон бенчмарка + анализ
		if err := runFullBenchmark(ctx, config); err != nil {
			log.Fatalf("Benchmark failed: %v", err)
		}
	}
}

func runFullBenchmark(ctx context.Context, config *cnfg.Config) error {
	log.Println("Starting full benchmark...")

	// Создание и запуск бенчмарк раннера
	benchmarkRunner := runner.NewBenchmarkRunner(config)

	startTime := time.Now()
	results, err := benchmarkRunner.RunFullBenchmark(ctx)
	if err != nil {
		return err
	}

	duration := time.Since(startTime)
	log.Printf("Benchmark completed in %v", duration)

	// Вывод путей к результатам
	for framework, path := range results {
		log.Printf("%s results: %s", framework, path)
	}

	// Анализ результатов
	log.Println("Starting results analysis...")
	if err := runAnalysis(config.ResultsDir); err != nil {
		return err
	}

	log.Println("Full benchmark and analysis completed successfully!")
	return nil
}

func runAnalysis(resultsDir string) error {
	analyzer := analyzer.NewResultAnalyzer(resultsDir)
	return analyzer.GenerateFullReport()
}
