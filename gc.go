package main

import (
	"os"
	"os/exec"
)

// RunRegistryGarbageCollection ...
func RunRegistryGarbageCollection() error {
	fields := GetGarbageCollectCommand()
	cmd := exec.Command(fields[0], fields[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}
