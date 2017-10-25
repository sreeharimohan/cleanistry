package main

import (
	"log"
	"testing"
)

func TestIsImageTagOutdated(t *testing.T) {
	isOutdated := IsImageTagOutdated("absolut-saloni_vb", "0910506-1ba49d7")
	// FailOnError(err)
	log.Println(isOutdated)
	t.Fatal()
}
