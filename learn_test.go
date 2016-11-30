package learn_test

import (
	"testing"

	learn "github.com/minodisk/learn-exit-panic"
)

// func TestExit(t *testing.T) {
// 	if !learn.Exit() {
// 		t.Error("fail to exit")
// 	}
// }

func TestPanic(t *testing.T) {
	if !learn.Panic() {
		t.Errorf("fail to panic")
	}
}
