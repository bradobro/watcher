package watch

import (
	"testing"

	"github.com/kindrid/gotest"
	"github.com/kindrid/gotest/should"
)

func TestSet(t *testing.T) {
	set := &Set{}
	gotest.Assert(t, set, should.NotBeNil)
}
