package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	gendata "github.com/CakeForKit/GinFiberBenchmarks.git/internal/gen_data"
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
