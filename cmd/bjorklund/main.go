package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/eduardostalinho/bjorklund"
)

func main() {
	args := os.Args[1:]

	pulses, _ := strconv.Atoi(args[0])
	slots, _ := strconv.Atoi(args[1])

	rotation := 0
	if len(args) > 2 {
		rotation, _ = strconv.Atoi(args[2])
	}
	pattern, err := bjorklund.CreatePattern(pulses, slots, rotation)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pattern)
}
