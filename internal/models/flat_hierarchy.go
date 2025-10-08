package models

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type FlatHierarchyStruct struct {
	Users    []User            `json:"users"`
	Products []Product         `json:"products"`
	Settings map[string]string `json:"settings"`
}

type User struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Age       int      `json:"age"`
	Active    bool     `json:"active"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	InStock     bool    `json:"in_stock"`
	Description string  `json:"description"`
}

// Генерируем плоскую иерархию с массивами объектов
func GenerateFlatHierarchyJSON() ([]byte, error) {
	data := map[string]interface{}{
		"users":    generateUsers(50),
		"products": generateProducts(30),
		"settings": generateSettings(),
	}

	return json.MarshalIndent(data, "", "  ")
}

func generateUsers(count int) []User {
	firstNames := []string{"John", "Jane", "Alex", "Maria", "Bob", "Alice", "Tom", "Sarah"}
	lastNames := []string{"Smith", "Johnson", "Brown", "Davis", "Wilson", "Taylor", "Clark"}
	domains := []string{"gmail.com", "yahoo.com", "hotmail.com", "company.com"}
	tags := []string{"premium", "basic", "vip", "new", "active", "inactive"}

	users := make([]User, count)
	for i := 0; i < count; i++ {
		firstName := firstNames[rand.Intn(len(firstNames))]
		lastName := lastNames[rand.Intn(len(lastNames))]

		users[i] = User{
			ID:        i + 1,
			Name:      fmt.Sprintf("%s %s", firstName, lastName),
			Email:     fmt.Sprintf("%s.%s@%s", firstName, lastName, domains[rand.Intn(len(domains))]),
			Age:       rand.Intn(50) + 18,
			Active:    rand.Intn(2) == 1,
			Tags:      randomTags(tags, 2),
			CreatedAt: time.Now().Add(-time.Duration(rand.Intn(365)) * 24 * time.Hour).Format("2006-01-02"),
		}
	}
	return users
}

func generateProducts(count int) []Product {
	categories := []string{"Electronics", "Books", "Clothing", "Home", "Sports", "Toys"}
	adjectives := []string{"Amazing", "Premium", "Basic", "Professional", "Deluxe", "Standard"}
	nouns := []string{"Laptop", "Book", "Shirt", "Chair", "Ball", "Game"}

	products := make([]Product, count)
	for i := 0; i < count; i++ {
		products[i] = Product{
			ID:          i + 1,
			Name:        fmt.Sprintf("%s %s", adjectives[rand.Intn(len(adjectives))], nouns[rand.Intn(len(nouns))]),
			Price:       float64(rand.Intn(1000)*100+rand.Intn(100)) / 100,
			Category:    categories[rand.Intn(len(categories))],
			InStock:     rand.Intn(2) == 1,
			Description: fmt.Sprintf("This is a %s product", randomString(15)),
		}
	}
	return products
}

func generateSettings() map[string]interface{} {
	return map[string]interface{}{
		"theme":         "dark",
		"language":      "en",
		"notifications": true,
		"auto_save":     false,
		"refresh_rate":  30,
	}
}

func randomTags(tags []string, count int) []string {
	result := make([]string, 0, count)
	used := make(map[int]bool)

	for len(result) < count && len(result) < len(tags) {
		idx := rand.Intn(len(tags))
		if !used[idx] {
			result = append(result, tags[idx])
			used[idx] = true
		}
	}

	return result
}
