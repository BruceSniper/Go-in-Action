package main

import (
	"github.com/go-programming-tour-book/tour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Excute err : %v", err)
	}
}
