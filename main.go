package main

import (
	"console/src/config"
	"console/src/service"
	"fmt"
)

func main() {
	config.Load()

	fmt.Println("Rodando Console")

	service.ProcessMessages()
}
