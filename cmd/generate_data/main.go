package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	gendata "github.com/CakeForKit/GinFiberBenchmarks.git/internal/gen_data"
	"github.com/CakeForKit/GinFiberBenchmarks.git/internal/models"
)

const (
	FLAT_AMMO_FILENAME      = "./requests/ammo/flat_ammo.json"
	HEIRARCHY_AMMO_FILENAME = "./requests/ammo/hierarchy_ammo.json"
	DEEP_AMMO_FILENAME      = "./requests/ammo/deep_ammo.json"
	count_different_ammo    = 5
	levels_in_tree          = 7
)

func main() {
	{
		flat_ammo, err := gendata.GenerateFlatAmmo(count_different_ammo)
		if err != nil {
			panic(err)
		}
		ammoJSON, err := json.MarshalIndent(flat_ammo, "", "  ")
		if err != nil {
			panic(fmt.Sprintf("Error marshaling ammo: %v", err))
		}
		if err := os.WriteFile(FLAT_AMMO_FILENAME, ammoJSON, 0644); err != nil {
			log.Fatalf("Error writing ammo file: %v", err)
		}

		fmt.Println("Flat ammo file with generated!")
	}
	{
		hierarchy_ammo, err := gendata.GenerateFlatHierarchyAmmo(count_different_ammo)
		if err != nil {
			panic(err)
		}
		ammoJSON, err := json.MarshalIndent(hierarchy_ammo, "", "  ")
		if err != nil {
			panic(fmt.Sprintf("Error marshaling ammo: %v", err))
		}
		if err := os.WriteFile(HEIRARCHY_AMMO_FILENAME, ammoJSON, 0644); err != nil {
			log.Fatalf("Error writing ammo file: %v", err)
		}

		fmt.Println("Hierarchy ammo file with generated!")
	}
	{
		deep_ammo, err := gendata.GenerateDeepAmmo(count_different_ammo, levels_in_tree)
		if err != nil {
			panic(err)
		}
		ammoJSON, err := json.MarshalIndent(deep_ammo, "", "  ")
		if err != nil {
			panic(fmt.Sprintf("Error marshaling ammo: %v", err))
		}
		if err := os.WriteFile(DEEP_AMMO_FILENAME, ammoJSON, 0644); err != nil {
			log.Fatalf("Error writing ammo file: %v", err)
		}

		fmt.Println("Deep ammo file with generated!")
	}
}

func main1() {
	// Создаем папку для результатов, если её нет
	outputDir := "json_data"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		panic(fmt.Sprintf("Не удалось создать директорию: %v", err))
	}

	// Генерируем временную метку для уникальности файлов
	// timestamp := time.Now().Format("20060102_150405")

	// 1. Генерируем плоский JSON со 100 полями
	flatJSON, err := models.GenerateFlatStructJSON()
	if err != nil {
		panic(err)
	}
	flatFilename := filepath.Join(outputDir, fmt.Sprintf("flat_struct.json"))
	if err := saveJSONToFile(flatJSON, flatFilename); err != nil {
		panic(err)
	}
	fmt.Printf("=== ПЛОСКИЙ JSON (100 полей) ===\n")
	fmt.Printf("Сохранен в: %s\n", flatFilename)
	fmt.Printf("Размер: %d байт\n\n", len(flatJSON))

	// 2. Генерируем глубоко вложенный JSON
	levelsInTree := 6
	deepJSON, err := models.GenerateDeepNestedJSON(levelsInTree)
	if err != nil {
		panic(err)
	}
	deepFilename := filepath.Join(outputDir, fmt.Sprintf("deep_nested.json"))
	if err := saveJSONToFile(deepJSON, deepFilename); err != nil {
		panic(err)
	}
	fmt.Printf("=== ГЛУБОКО ВЛОЖЕННЫЙ JSON ===\n")
	fmt.Printf("Сохранен в: %s\n", deepFilename)
	fmt.Printf("Уровней в дереве: %d\n", levelsInTree)
	fmt.Printf("Размер: %d байт\n\n", len(deepJSON))

	// 3. Генерируем плоскую иерархию
	hierarchyJSON, err := models.GenerateFlatHierarchyJSON()
	if err != nil {
		panic(err)
	}
	hierarchyFilename := filepath.Join(outputDir, fmt.Sprintf("flat_hierarchy.json"))
	if err := saveJSONToFile(hierarchyJSON, hierarchyFilename); err != nil {
		panic(err)
	}
	fmt.Printf("=== ПЛОСКАЯ ИЕРАРХИЯ ===\n")
	fmt.Printf("Сохранен в: %s\n", hierarchyFilename)
	fmt.Printf("Размер: %d байт\n\n", len(hierarchyJSON))

	// Выводим разделитель
	fmt.Println(strings.Repeat("=", 50))

	// Сводная информация
	fmt.Printf("\n📁 Все файлы сохранены в папке: %s/\n", outputDir)
	fmt.Printf("📊 Итого сгенерировано:\n")
	fmt.Printf("   • Плоская структура: %s\n", filepath.Base(flatFilename))
	fmt.Printf("   • Глубоко вложенный JSON: %s\n", filepath.Base(deepFilename))
	fmt.Printf("   • Плоская иерархия: %s\n", filepath.Base(hierarchyFilename))
}

// Функция для сохранения JSON в файл
func saveJSONToFile(data []byte, filename string) error {
	return os.WriteFile(filename, data, 0644)
}

func mainprint() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Генерируем плоский JSON со 100 полями
	flatJSON, err := models.GenerateFlatStructJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println("=== ПЛОСКИЙ JSON (100 полей) ===")
	fmt.Println(string(flatJSON))
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// Генерируем глубоко вложенный JSON
	levelsInTree := 6
	deepJSON, err := models.GenerateDeepNestedJSON(levelsInTree)
	if err != nil {
		panic(err)
	}
	fmt.Println("=== ГЛУБОКО ВЛОЖЕННЫЙ JSON ===")
	fmt.Println(string(deepJSON))
	fmt.Println("\n" + strings.Repeat("=", 50) + "\n")

	// Генерируем плоскую иерархию
	hierarchyJSON, err := models.GenerateFlatHierarchyJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println("=== ПЛОСКАЯ ИЕРАРХИЯ ===")
	fmt.Println(string(hierarchyJSON))
}
