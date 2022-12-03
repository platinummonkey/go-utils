// Package integration provides an argument to determine whether or not a set of tests should run.
package integration

import (
	"os"
	"testing"
)

// ShouldRun determines whether tests should run or be skipped.
var ShouldRun = false

// ShouldRunBenchmark determines whether benchmark tests should run or be skipped.
var ShouldRunBenchmark = false

func init() {
	for _, arg := range os.Args[1:] {
		if arg == "integration" {
			ShouldRun = true
		}
		if arg == "integration_bench" {
			ShouldRunBenchmark = true
		}
	}
	if !ShouldRun {
		if _, ok := os.LookupEnv("INTEGRATION_TEST"); ok {
			ShouldRun = true
		}
	}
	if !ShouldRunBenchmark {
		if _, ok := os.LookupEnv("INTEGRATION_BENCH_TEST"); ok {
			ShouldRunBenchmark = true
		}
	}
}

// Check skips the current testing context if ShouldRun is true. If `-v` is not
// enabled for the test run, it won't be obvious the test was skipped.
func Check(t *testing.T) {
	if !ShouldRun {
		t.Skipf("warning: skipped %s because integrations tests are disabled. enable with `-args integration`.\n", t.Name())
	}
}

// CheckBenchmark check skips the current benchmark testing context if ShouldRunBenchmark is false. If `-v` is not
// enabled for the benchmark run, it won't be obvious the benchmark test was skipped.
func CheckBenchmark(b *testing.B) {
	if !ShouldRunBenchmark {
		b.Skipf("warning: skipped %s because integrations benchmarks are disabled. enable with `-args integration_bench`.\n", b.Name())
	}
}
