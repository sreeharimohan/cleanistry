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
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
