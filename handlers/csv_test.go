package handlers

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"testing"

	calclib "github.com/smartyjenny/calc-lib"
)

var csvInput = strings.Join([]string{
	"1,+,2",
	"2,-,1",
	//"NaN,+,2",
	//"1,+,NaN",
	"1,nop,2",
	"3,+,4",
	"20,/,10",
}, "\n")

var csvOutput = strings.Join([]string{
	"1,+,2,3",
	"3,+,4,7",
	"",
}, "\n")

func TestCSVHandler(t *testing.T) {
	var logBuffer bytes.Buffer
	logger := log.New(&logBuffer, "", log.LstdFlags)
	reader := strings.NewReader(csvInput)
	var outputBuffer bytes.Buffer
	calculators := map[string]Calculator{"+": &calclib.Addition{}}
	handler := NewCSVHandler(logger, reader, &outputBuffer, calculators)
	err := handler.Handle()
	assertError(t, err, nil)
	if outputBuffer.String() != csvOutput {
		t.Errorf("got %q, expected %q", outputBuffer.String(), csvOutput)
	}

	t.Log(logBuffer.String())
}

func TestCSVHandler_WriteError(t *testing.T) {
	var logBuffer bytes.Buffer
	logger := log.New(&logBuffer, "", log.LstdFlags)
	reader := strings.NewReader(csvInput)
	boink := errors.New("boink")
	output := ErrWriter{err: boink}
	calculators := map[string]Calculator{"+": &calclib.Addition{}}
	handler := NewCSVHandler(logger, reader, &output, calculators)
	err := handler.Handle()
	assertError(t, err, boink)
}

func TestCSVHandler_ReadError(t *testing.T) {
	var logBuffer bytes.Buffer
	logger := log.New(&logBuffer, "", log.LstdFlags)
	boink := errors.New("boink")
	reader := ErrReader{
		err: boink,
	}
	var outputBuffer bytes.Buffer
	calculators := map[string]Calculator{"+": &calclib.Addition{}}
	handler := NewCSVHandler(logger, &reader, &outputBuffer, calculators)
	err := handler.Handle()
	assertError(t, err, boink)
}
