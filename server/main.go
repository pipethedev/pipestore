package main

import (
	"fmt"
	"pipebase/server/core"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	fmt.Println("Pipestore is starting..")

	core.StartTCP()
}
