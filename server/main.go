package main

import (
	"fmt"
	"pipebase/cli/utils"
	"pipebase/server/core"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	fmt.Println("Pipestore is starting..")

	port, err := utils.AvailablePort()

	if err != nil {
		fmt.Println("Unable to assign port")
	}

	core.StartTCP(port)
}
