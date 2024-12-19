package should_test

import (
	"errors"
	"fmt"
	"testing"

	"backend-training/cohort-c-2/calc-apps/externals/should"
)

type FakeT struct {
	err    error
	helped bool
}

func (this *FakeT) Helper() {
	this.helped = true
}

func (this *FakeT) Error(values ...any) {
	this.err = values[0].(error)
}

func TestShouldEqual_PrimitiveTypes(t *testing.T) {
	pass(t, 1, should.Equal, 1)
}

func TestShouldEqual_ComplicatedTypes(t *testing.T) {
	fail(t, []int{1, 2, 3}, should.Equal, []int{4, 5, 6})
}

func TestShouldEqual_ComplicatedTypes_Equal(t *testing.T) {
	fakeT := &FakeT{}

	should.So(fakeT, []int{1, 2, 3}, should.Equal, []int{1, 2, 3})
	if fakeT.err != nil {
		t.Errorf("should not get an error, but did: %v", fakeT.err)
	}

}

func TestShouldBeTrue(t *testing.T) {
	pass(t, true, should.BeTrue)
	fail(t, false, should.BeTrue)

}

func TestShouldBeFalse(t *testing.T) {
	pass(t, false, should.BeFalse)
	fail(t, true, should.BeFalse)
}

func TestShouldBeNil(t *testing.T) {
	// var nilTime *time.Time
	pass(t, nil, should.BeNil)
	// pass(t, nilTime, should.BeNil)
	fail(t, 1, should.BeNil)
	fail(t, []int{1}, should.BeNil)
	fail(t, false, should.BeNil)
	fail(t, true, should.BeNil)
}

func TestShouldNotEqual(t *testing.T) {
	pass(t, 1, should.NOT.Equal, 2)
}

func TestShouldWrapError(t *testing.T) {
	inner := errors.New("inner")
	outer := fmt.Errorf("output %w", inner)
	pass(t, outer, should.WrapError, inner)
	fail(t, inner, should.WrapError, outer)
}

func pass(t *testing.T, actual any, assert should.Assertion, expected ...any) {
	fakeT := &FakeT{}

	should.So(fakeT, actual, assert, expected...)
	if fakeT.err != nil {
		t.Errorf("should not get an error, but did: %v", fakeT.err)
	}
}

func fail(t *testing.T, actual any, assert should.Assertion, expected ...any) {
	fakeT := &FakeT{}

	should.So(fakeT, actual, assert, expected...)
	if !errors.Is(fakeT.err, should.ErrAssertionFailure) {
		t.Errorf("should get an Assertion error, but didn't: %v", fakeT.err)
	} else {
		t.Log(fakeT.err)
	}
	if !fakeT.helped {
		t.Errorf("should have called t.Helper(), but didn't")
	}
}
