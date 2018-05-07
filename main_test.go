package main

import (
	"os"
	"path/filepath"
	"testing"
)

// Test fixture directories
var mainDir = ""
var parallelDir = ""
var noFileDir = ""
var noSakuNoDirectiveDir = ""
var readmeDirectiveDir = ""

func TestMain(m *testing.M) {
	cwd, _ := os.Getwd()

	mainDir = filepath.Join(cwd, "fixture", "main")
	parallelDir = filepath.Join(cwd, "fixture", "parallel")
	noFileDir = filepath.Join(cwd, "fixture", "no-files")
	noSakuNoDirectiveDir = filepath.Join(cwd, "fixture", "no-saku-md-no-directive")
	readmeDirectiveDir = filepath.Join(cwd, "fixture", "readme-directive")

	os.Exit(m.Run())
}

func TestHelpAction(t *testing.T) {
	if Run(mainDir, "-h") == exitCodeError {
		t.Error("--help option should return exit code = 0")
	}
}

func TestVersionAction(t *testing.T) {
	if Run(mainDir, "-v") == exitCodeError {
		t.Error("--version option should return exit code = 0")
	}
}

func TestInvalidFlag(t *testing.T) {
	if Run(mainDir, "--") == exitCodeOk {
		t.Error("Should exits with error if called with --")
	}
}

func TestConfigNotFound(t *testing.T) {
	if Run(mainDir, "-c", "foo.md") == exitCodeOk {
		t.Error("Should exit with error if the specified config not found")
	}
}

func TestConfigNotFoundReadmeNotFound(t *testing.T) {
	if Run(noFileDir) == exitCodeOk {
		t.Error("Should exit with error if both saku.md and readme.md not found")
	}
}

func TestConfigNotFoundDirectiveNotFound(t *testing.T) {
	if Run(noSakuNoDirectiveDir) == exitCodeOk {
		t.Error("Should exit with error if saku.md not found and <!-- saku start --><!-- saku end --> directive not found in readme.md")
	}
}

func TestReadmeDirectiveConfig(t *testing.T) {
	if Run(readmeDirectiveDir) == exitCodeError {
		t.Error("saku can read config from the contents between <!-- saku start --><!-- saku end --> in README.md")
	}
}

func TestInfoAction(t *testing.T) {
	if Run(mainDir) == exitCodeError {
		t.Error("no option invokes info action and exits with 0")
	}
}

func TestParallelAndSeriallOptions(t *testing.T) {
	if Run(mainDir, "--serial", "--parallel", "hello") == exitCodeOk {
		t.Error("It is error if both --serial and --parallel specified")
	}
}

func TestNoTask(t *testing.T) {
	if Run(mainDir, "bar") == exitCodeOk {
		t.Error("invoking with non task fails")
	}
}

func TestSingleTask(t *testing.T) {
	if Run(mainDir, "hello") == exitCodeError {
		t.Error("invoking with a single task passes")
	}
}

func TestSingleErrorTask(t *testing.T) {
	if Run(mainDir, "foo") == exitCodeOk {
		t.Error("invoking with failing task causes command exit with error")
	}
}

func TestMultipleTask(t *testing.T) {
	if Run(mainDir, "hello", "hello") == exitCodeError {
		t.Error("Should exit with 0 if all the sequencial tasks pass")
	}
}

func TestParallel(t *testing.T) {
	code := Run(parallelDir, "-p", "1sec-ok", "2sec-ok")

	if code == exitCodeError {
		t.Error("Should exit with 0 if all parallel tasks exit with 0")
	}

	// TODO: Assert about the commands are invoked concurrently
}

func TestParallelFail(t *testing.T) {
	code := Run(parallelDir, "-p", "1sec-fail", "2sec-ok")

	if code == exitCodeOk {
		t.Error("Should fail if one of task failed")
	}
}

func TestParallelRace(t *testing.T) {
	code := Run(parallelDir, "-r", "-p", "1sec-ok", "2sec-fail")

	if code == exitCodeError {
		t.Error("Should exit with 0 if the first task exit with 0")
	}
}

func TestParallelRaceFail(t *testing.T) {
	code := Run(parallelDir, "-r", "-p", "1sec-fail", "2sec-ok")

	if code == exitCodeOk {
		t.Error("Should exit with error if the first task failed")
	}
}
