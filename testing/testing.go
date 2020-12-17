package testing

// TestingFlag is a global flag, set with the
// command line flag -test. It can be used to
// disable certain features during development
// such as superfluous calls to external APIs.
var TestingFlag bool
