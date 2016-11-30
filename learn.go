package learn

import "os"

func Exit() bool {
	go func() {
		os.Exit(0)
	}()
	return true
}

func Panic() bool {
	go func() {
		panic("OMG")
	}()
	return true
}
