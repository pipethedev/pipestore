package main

import (
	"fmt"
	"server/config"
	"server/core"
	"server/core/operations"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := config.LoadConfig().PORT

	fmt.Printf("Starting pipestore at port %d\n", port)

	operations.StartIndexing()

	core.StartTCP(port)
}
