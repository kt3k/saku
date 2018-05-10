package saku

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	md := []byte("# hello\n> hello\n\n    echo hello")

	tasks := ParseConfig(&md)

	if len(tasks.tasks) != 1 {
		t.Fail()
	}
}
