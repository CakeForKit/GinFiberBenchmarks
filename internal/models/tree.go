package models

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type TreeNode struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Value    float64    `json:"value"`
	Children []TreeNode `json:"children,omitempty"`
}

func GenerateDeepNestedJSON(levels int) ([]byte, error) {
	root := generateDeepTree(levels, 0) // Дерево глубиной 6 уровней
	return json.MarshalIndent(root, "", "  ")
}

// Генерируем глубоко вложенное дерево
func generateDeepTree(maxDepth, currentDepth int) TreeNode {
	if currentDepth > maxDepth {
		return TreeNode{}
	}

	node := TreeNode{
		ID:    fmt.Sprintf("node_%d_%d", currentDepth, rand.Intn(1000)),
		Name:  randomString(8),
		Value: rand.Float64() * 100,
	}

	// Рекурсивно добавляем детей
	if currentDepth < maxDepth {
		numChildren := rand.Intn(3) + 1 // 1-3 детей
		for i := 0; i < numChildren; i++ {
			child := generateDeepTree(maxDepth, currentDepth+1)
			if child.ID != "" {
				node.Children = append(node.Children, child)
			}
		}
	}

	return node
}
