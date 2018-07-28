package saku

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	md := []byte(`
# hello
> hello
	echo hello
`)

	tasks := ParseConfig(&md)

	if len(tasks.tasks) != 1 {
		t.Fail()
	}
}

func TestParserComplexConfig(t *testing.T) {
	md := []byte(`
# foo
> desc foo

	echo foo

## bar
> desc bar

	echo bar

## baz
> desc baz

	echo baz

### qux
> desc qux

	echo qux

### quux
> desc quux

	echo quux

# test
> desc test

    echo test
`)
	tasks := ParseConfig(&md)

	if len(tasks.tasks) != 2 {
		t.Fail()
	}

	if len(tasks.tasks[0].children.tasks) != 2 {
		t.Fail()
	}

	if tasks.taskCount() != 6 {
		t.Fail()
	}
}
