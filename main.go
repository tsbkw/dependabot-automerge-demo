package main

import (
	// All modules here are choosen because of it simplicity (few dependencies) and existence of patch version
	"github.com/google/uuid"
	"github.com/hashicorp/go-version"
	"google.golang.org/grpc/balancer/grpclb/state"
	"google.golang.org/grpc/resolver"
)

func main() {
	// Module for testing skip auto merge when check fail.
	// Note: 
	//  As grpc revert changing module name `state` to `grpclbstate` in patch version v1.52.1,
	//  code below cannot be built when using module greater than v1.52.1
	state.Get(resolver.State{})


	// Module for testing auto merge works
	uuid.ClockSequence()

	// Module for testing skip auto merge for specific module works
	version.NewVersion("1.1")
}