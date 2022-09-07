package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println(id.String())

	file, err := os.Create("token.tmp")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(id.String())
}
