package main

import (
	authzModels "authz/models"
	"fmt"

	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {
	err := newConverter("authz.ts", authzModels.Group{}, authzModels.Permission{})
	if err != nil {
		fmt.Println(err)
	}
}

func newConverter(fileName string, models ...interface{}) error {
	converter := typescriptify.New().
		WithConstructor(true).
		WithInterface(true).
		WithIndent("  "). // Two spaces
		WithBackupDir("../types/backup")

	for _, model := range models {
		converter.Add(model)
	}

	return converter.ConvertToFile(fmt.Sprintf("../types/generated/%s", fileName))
}
