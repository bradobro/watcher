package main

import (
	"testing"

	"github.com/kindrid/gotest"
	"github.com/kindrid/gotest/should"
)

func TestCli(t *testing.T) {
	gotest.Assert(t, cli(), should.Equal, 0)
	gotest.Assert(t, cli("simple"), should.Equal, 1)
}
