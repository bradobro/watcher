package watch

import (
	"testing"

	"github.com/kindrid/gotest"
	"github.com/kindrid/gotest/should"
)

func TestJob(t *testing.T) {
	job := &Job{}
	gotest.Assert(t, job, should.NotBeNil)

}
