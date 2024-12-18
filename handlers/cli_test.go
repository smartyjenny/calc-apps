package handlers

import (
	"bytes"
	"errors"
	"testing"

	"backend-training/cohort-c-2/calc-apps/externals/should"

	calclib "github.com/smartyjenny/calc-lib"
)

func assertError(t *testing.T, actual, target error) {
	if should.Equal(actual, target) != nil {
		t.Errorf("got: %v, want: %v", actual, target)
	}

}

func TestHandler_WrongNumberOfArgs(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	should.So(t, err, should.WrapError, errWrongNumberOfArgs)
}

func TestHandler_InvalidFirstArg(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"INVALID", "2"})
	should.So(t, err, should.WrapError, errInvalidArg)
}

func TestHandler_InvalidSecondArg(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"2", "INVALID"})
	assertError(t, err, errInvalidArg)
}

func TestHandler_OutputWriter(t *testing.T) {
	boink := errors.New("boink")
	writer := &ErrWriter{err: boink}
	handler := NewHandler(writer, nil)
	err := handler.Handle([]string{"1", "2"})
	assertError(t, err, errWriterFailure)
}

func TestHandler_HappyPath(t *testing.T) {
	writer := &bytes.Buffer{}
	handler := NewHandler(writer, &calclib.Addition{})
	err := handler.Handle([]string{"1", "2"})
	assertError(t, err, nil)
	if writer.String() != "3" {
		t.Errorf("expected 3, got %s", writer.String())
	}
}

type ErrWriter struct {
	err error
}

func (this *ErrWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}

type ErrReader struct {
	err error
}

func (this *ErrReader) Read(p []byte) (n int, err error) {
	return 0, this.err
}
