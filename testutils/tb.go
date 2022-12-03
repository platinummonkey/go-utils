package testuils

// T is the interface common to testing.{T|B}
//
// This is copied from testing.TB to allow a
// non-test package to have functions that rely
// on testing.T without adding testing flags to
// the binary.
//
// To find packages that import "testing" in a non-test
// package run:
// grep -R --exclude '*_test.go' --include '*.go' '"testing"' .
type T interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
}
