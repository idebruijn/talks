package main

// START OMIT

// T is the interface that this package is using from standard testing.T

type T interface {
	Logf(format string, args ...interface{})

	Error(args ...interface{})

	Fatal(args ...interface{})
}

// END OMIT
