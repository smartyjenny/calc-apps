package main

import (
	"fmt"
	"os"

	"github.com/smartyjenny/calc-apps/handlers"
	calclib "github.com/smartyjenny/calc-lib"
)

func main() {

	handler := handlers.NewHandler(os.Stdout, &calclib.Addition{})
	err := handler.Handle(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
}
