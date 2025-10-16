package gendata

import "github.com/CakeForKit/GinFiberBenchmarks.git/internal/models"

func GenerateFlatAmmo(count int) ([]AmmoRequest, error) {
	ammoRes := []AmmoRequest{}
	for range count {
		jsonData, err := models.GenerateFlatStructJSON()
		if err != nil {
			return ammoRes, err
		}

		ammoRes = append(ammoRes, AmmoRequest{
			Method: "POST",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: string(jsonData),
			// Tag:  "flat_request_1",
			URI: "/flat",
		})
	}
	return ammoRes, nil
}

func GenerateFlatHierarchyAmmo(count int) ([]AmmoRequest, error) {
	ammoRes := []AmmoRequest{}
	for range count {
		jsonData, err := models.GenerateFlatHierarchyJSON()
		if err != nil {
			return ammoRes, err
		}

		ammoRes = append(ammoRes, AmmoRequest{
			Method: "POST",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: string(jsonData),
			// Tag:  "flat_request_1",
			URI: "/hierarchy",
		})
	}
	return ammoRes, nil
}

func GenerateDeepAmmo(count, levelsInTree int) ([]AmmoRequest, error) {
	ammoRes := []AmmoRequest{}
	for range count {
		jsonData, err := models.GenerateDeepNestedJSON(levelsInTree)
		if err != nil {
			return ammoRes, err
		}

		ammoRes = append(ammoRes, AmmoRequest{
			Method: "POST",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: string(jsonData),
			// Tag:  "flat_request_1",
			URI: "/deep",
		})
	}
	return ammoRes, nil
}
