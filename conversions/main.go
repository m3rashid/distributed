package main

import (
	authzModels "authz/models"
	"fmt"

	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {
	// authz converter
	fileName, converter := newStructConverter("authz.ts", authzModels.Group{}, authzModels.Permission{})
	// converter.AddEnum(authzModels.ALL_RELATION_NAMES) // enums cannot have numeric keys
	converter.AddEnum(authzModels.ALL_RELATIONS)
	convert(fileName, converter)

	// .... other converters
}

func convert(fileName string, converter *typescriptify.TypeScriptify) {
	err := converter.ConvertToFile(fileName)
	if err != nil {
		fmt.Println("error in converting", fileName)
		fmt.Println(err)
	}
}

func newStructConverter(fileName string, models ...interface{}) (string, *typescriptify.TypeScriptify) {
	converter := typescriptify.New().
		WithConstructor(true).
		WithInterface(true).
		WithIndent("  "). // Two spaces
		WithBackupDir("../types/backup")

	for _, model := range models {
		converter.Add(model)
	}

	return fmt.Sprintf("../types/generated/%s", fileName), converter
}
