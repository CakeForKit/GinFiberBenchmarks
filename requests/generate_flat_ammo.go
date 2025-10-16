package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const (
	FLAT_FILENAME        = "./requests/ammo/flat_ammo.json"
	FLAT_STRUCT_FILENAME = "./json_data/flat_struct.json"
)

func main() {
	// Читаем ваш существующий JSON файл
	jsonData, err := os.ReadFile(FLAT_STRUCT_FILENAME)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Валидируем JSON
	var jsonContent interface{}
	if err := json.Unmarshal(jsonData, &jsonContent); err != nil {
		log.Fatalf("Invalid JSON file: %v", err)
	}

	// Конвертируем обратно в строку для тела запроса
	jsonString := string(jsonData)

	// Создаем два одинаковых запроса
	ammo := []AmmoRequest{
		{
			Method: "POST",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: jsonString,
			Tag:  "flat_request_1",
			URI:  "/flat",
		},
		{
			Method: "POST",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: jsonString,
			Tag:  "flat_request_2",
			URI:  "/flat",
		},
	}

	// Сохраняем ammo файл
	ammoJSON, err := json.MarshalIndent(ammo, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling ammo: %v", err)
	}

	if err := os.WriteFile(FLAT_FILENAME, ammoJSON, 0644); err != nil {
		log.Fatalf("Error writing ammo file: %v", err)
	}

	fmt.Println("Ammo file with 2 identical requests generated!")
}
