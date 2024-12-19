package gunit_test

import (
	"testing"

	"backend-training/cohort-c-2/calc-apps/externals/gunit"
	"backend-training/cohort-c-2/calc-apps/externals/should"
)

func TestMySuperCoolFixture(t *testing.T) {
	gunit.Run(t, new(MySuperCoolFixture))
}

type MySuperCoolFixture struct {
	*gunit.Fixture
}

func (this *MySuperCoolFixture) Test1() {
	this.So(1, should.Equal, 1)
}
