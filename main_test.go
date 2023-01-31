package main

import (
	"runtime/debug"
	"testing"
)


func TestMain(t *testing.T) {
	bi, _ := debug.ReadBuildInfo()
	for _, dp := range bi.Deps {
		if dp.Path == "google.golang.org/grpc" && dp.Version != "v1.52.1" {
			// fail in CI to check auto merge disabled on check failure
			t.Fail()
		}
	}
}