package models

import (
	"encoding/json"
	"os"
)

const (
	FlatStructFilename    = "./json_data/flat_struct.json"
	TreeNodeFilename      = "./json_data/deep_nested.json"
	FlatHierarchyFilename = "./json_data/flat_hierarchy.json"
)

func LoadTestData() (simple FlatStruct, tree TreeNode, wide FlatHierarchyStruct) {
	// Загрузка тестовых данных из файлов
	simpleData, _ := os.ReadFile(FlatStructFilename)
	json.Unmarshal(simpleData, &simple)

	treeData, _ := os.ReadFile(TreeNodeFilename)
	json.Unmarshal(treeData, &tree)

	wideData, _ := os.ReadFile(FlatHierarchyFilename)
	json.Unmarshal(wideData, &wide)

	return simple, tree, wide
}
