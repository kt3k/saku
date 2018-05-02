package main

import (
	"testing"
)

func TestHelpAction(t *testing.T) {
	if run("saku", "-h") == exitCodeError {
		t.Error("--help option should return exit code = 0")
	}
}

func TestVersionAction(t *testing.T) {
	if run("saku", "-v") == exitCodeError {
		t.Error("--version option should return exit code = 0")
	}
}

func TestInvalidFlag(t *testing.T) {
	if run("saku", "--") == exitCodeOk {
		t.Error("Should exits with error if called with --")
	}
}

func TestConfigNotFound(t *testing.T) {
	if run("saku", "-c", "foo.md") == exitCodeOk {
		t.Error("Should exit with error if the config not found")
	}
}

func TestInfoAction(t *testing.T) {
	if run("saku") == exitCodeError {
		t.Error("no option invokes info action and exits with 0")
	}
}

func TestParallelAndSeriallOptions(t *testing.T) {
	if run("saku", "--serial", "--parallel", "hello") == exitCodeOk {
		t.Error("It is error if both --serial and --parallel specified")
	}
}

func TestNoTask(t *testing.T) {
	if run("saku", "bar") == exitCodeOk {
		t.Error("invoking with non task fails")
	}
}

func TestSingleTask(t *testing.T) {
	if run("saku", "hello") == exitCodeError {
		t.Error("invoking with a single task passes")
	}
}

func TestSingleErrorTask(t *testing.T) {
	if run("saku", "foo") == exitCodeOk {
		t.Error("invoking with failing task causes command exit with error")
	}
}
