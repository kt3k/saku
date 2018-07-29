package saku

import (
	"fmt"
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

aaa <!-- saku parallel --> bbb

### qux
> desc qux

	echo qux

### quux
> desc quux

	echo quux

# test
> desc test

	echo test

<!-- saku parallel race -->

## test-a
> desc test-a

	echo test-a

## test-b
> desc test-b

	echo test-b
`)
	tasks := ParseConfig(&md)

	if len(tasks.tasks) != 2 {
		t.Error("There should be 2 tasks.")
	}

	foo := tasks.tasks[0]

	if len(foo.children.tasks) != 2 {
		t.Error("There should be 2 children in foo.")
	}

	if foo.children.mode != RunModeSequence {
		t.Error("The children of foo should run in sequence")
	}

	baz := foo.children.tasks[1]

	if len(baz.children.tasks) != 2 {
		t.Error("There should be 2 children in baz.")
	}

	fmt.Printf("%#v", baz)

	if baz.children.mode != RunModeParallel {
		t.Error("The children of baz should run in parallel")
	}

	test := tasks.tasks[1]

	if test.children.mode != RunModeParallelRace {
		t.Error("The children of test should run in parallel-race")
	}

	if tasks.taskCount() != 8 {
		t.Error("There should be 8 tasks.")
	}
}
