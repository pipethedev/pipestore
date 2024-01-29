package main

import (
	"fmt"
	"pipebase/cli/utils"
	"pipebase/server/core"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	port, err := utils.AvailablePort()

	if err != nil {
		fmt.Println("Unable to assign port")
	}

	fmt.Printf("Starting pipestore at port %d\n", port)

	core.StartTCP(port)
}
