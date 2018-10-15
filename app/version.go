package main

import (
	"fmt"
	"runtime"
)

var (
	Version = "0.0.1"
	DATE    = "2018/10/12"
)

// SetVersion for setup Version string.
func SetVersion(ver string) {
	Version = ver
}

// PrintVersion provide print server engine
func PrintVersion() {
	fmt.Printf(`px-proxy %s, Compiler: %s %s, Copyright (C) %s Alex Stocks.\n`,
		Version,
		runtime.Compiler,
		runtime.Version(),
		DATE,
	)
}
