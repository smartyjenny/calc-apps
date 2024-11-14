package main

import (
	"flag"
	"fmt"
	"os"

	"backend-training/cohort-c-2/calc-apps/handlers"

	calclib "github.com/smartyjenny/calc-lib"
)

func main() {

	var operation string
	flag.StringVar(&operation, "op", "+", "The mathematical operation")
	flag.Parse()
	fmt.Println(operation)
	handler := handlers.NewHandler(os.Stdout, &calclib.Addition{})
	err := handler.Handle(flag.Args())
	if err != nil {
		fmt.Println(err)
	}
}

var calculators = map[string]handlers.Calculator{
	"+": &calclib.Addition{},
	"-": &calclib.Subtraction{},
	"*": &calclib.Multiplication{},
	"/": &calclib.Division{},
}
