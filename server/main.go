package main

import (
	"fmt"
	"pipebase/server/config"
	"pipebase/server/core"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	port := config.LoadConfig().PORT

	fmt.Printf("Starting pipestore at port %d\n", port)

	core.StartTCP(port)
}
