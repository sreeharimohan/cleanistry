package main

import (
	"os"
	"testing"
)

func TestForAnyCommand(t *testing.T) {
	os.Setenv("CLEANISTRY_GARBAGE_COLLECT_COMMAND", "docker ps -a")
	CheckAndGetConfigs()
	FailOnError(RunRegistryGarbageCollection(), "Step: Test - Any Command")
}
