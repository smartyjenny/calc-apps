package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	calclib "github.com/smartyjenny/calc-lib"
)

type Handler struct {
	stdout     io.Writer
	calculator *calclib.Addition
}

func NewHandler(stdout io.Writer, calculator *calclib.Addition) Handler {
	return Handler{
		stdout:     stdout,
		calculator: calculator,
	}
}

func (this *Handler) Handle(args []string) error {
	if len(args) < 2 {
		return errWrongNumberOfArgs
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return errInvalidArg
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return errInvalidArg
	}

	result := this.calculator.Calculate(a, b)

	_, err = fmt.Fprint(this.stdout, result)
	if err != nil {
		return errWriterFailure
	}

	return nil
}

var (
	errWrongNumberOfArgs = errors.New("usages: calc [a] [b]")
	errInvalidArg        = errors.New("invalidArgument")
	errWriterFailure     = errors.New("writer failure")
)
