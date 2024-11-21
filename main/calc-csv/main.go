package main

import (
	"log"
	"os"

	"backend-training/cohort-c-2/calc-apps/handlers"

	calclib "github.com/smartyjenny/calc-lib"
)

func main() {
	logger := log.New(os.Stderr, "", log.LstdFlags)
	handler := handlers.NewCSVHandler(logger, os.Stdin, os.Stdout, calculators)
	err := handler.Handle()
	if err != nil {
		log.Fatal(err)
	}
}

var calculators = map[string]handlers.Calculator{
	"+": &calclib.Addition{},
	"-": &calclib.Subtraction{},
	"*": &calclib.Multiplication{},
	"/": &calclib.Division{},
}
