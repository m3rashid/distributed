package main

import (
	"fmt"
	permissionModels "permissions/models"

	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {
	// permissionModels permissionConverter
	fileName, permissionConverter := newStructConverter("permissions.ts", permissionModels.Group{}, permissionModels.Permission{})
	// converter.AddEnum(permissionModels.ALL_RELATION_NAMES) // enums cannot have numeric keys
	permissionConverter.AddEnum(permissionModels.ALL_RELATIONS)
	convert(fileName, permissionConverter)

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
		WithBackupDir("backup")

	for _, model := range models {
		converter.Add(model)
	}

	return fmt.Sprintf("generated/%s", fileName), converter
}
