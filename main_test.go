package main

import (
	"os"
	"path/filepath"
	"testing"
)

var cwd = ""

func TestMain(m *testing.M) {
	cwd, _ = os.Getwd()

	os.Exit(m.Run())
}

func TestHelpAction(t *testing.T) {
	if run(cwd, "-h") == exitCodeError {
		t.Error("--help option should return exit code = 0")
	}
}

func TestVersionAction(t *testing.T) {
	if run(cwd, "-v") == exitCodeError {
		t.Error("--version option should return exit code = 0")
	}
}

func TestInvalidFlag(t *testing.T) {
	if run(cwd, "--") == exitCodeOk {
		t.Error("Should exits with error if called with --")
	}
}

func TestConfigNotFound(t *testing.T) {
	if run(cwd, "-c", "foo.md") == exitCodeOk {
		t.Error("Should exit with error if the specified config not found")
	}
}

func TestConfigNotFoundReadmeNotFound(t *testing.T) {
	if run(filepath.Join(cwd, "fixture", "no-files")) == exitCodeOk {
		t.Error("Should exit with error if both saku.md and readme.md not found")
	}
}

func TestConfigNotFoundDirectiveNotFound(t *testing.T) {
	if run(filepath.Join(cwd, "fixture", "no-saku-md-no-directive")) == exitCodeOk {
		t.Error("Should exit with error if saku.md not found and <!-- saku start --><!-- saku end --> directive not found in readme.md")
	}
}

func TestReadmeDirectiveConfig(t *testing.T) {
	if run(filepath.Join(cwd, "fixture", "readme-directive")) == exitCodeError {
		t.Error("saku can read config from the contents between <!-- saku start --><!-- saku end --> in README.md")
	}
}

func TestInfoAction(t *testing.T) {
	if run(cwd) == exitCodeError {
		t.Error("no option invokes info action and exits with 0")
	}
}

func TestParallelAndSeriallOptions(t *testing.T) {
	if run(cwd, "--serial", "--parallel", "hello") == exitCodeOk {
		t.Error("It is error if both --serial and --parallel specified")
	}
}

func TestNoTask(t *testing.T) {
	if run(cwd, "bar") == exitCodeOk {
		t.Error("invoking with non task fails")
	}
}

func TestSingleTask(t *testing.T) {
	if run(cwd, "hello") == exitCodeError {
		t.Error("invoking with a single task passes")
	}
}

func TestSingleErrorTask(t *testing.T) {
	if run(cwd, "foo") == exitCodeOk {
		t.Error("invoking with failing task causes command exit with error")
	}
}