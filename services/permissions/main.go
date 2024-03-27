package main

import (
	"permissions/utils"
)

func main() {
	err := utils.ParseEnv()
	if err != nil {
		panic(err)
	}

}
