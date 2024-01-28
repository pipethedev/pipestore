package main

import (
	"fmt"
	"pipebase/server/core"
)

func main() {
	fmt.Println("Pipestore is starting..")

	core.StartTCP("pipethedev", "123")
}
