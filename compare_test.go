package main

import (
	"testing"
)

func TestIsImageTagOutdated(t *testing.T) {
	isOutdated, err := IsImageTagOutdated("absolut-saloni_vb", "0910506-1ba49d7")
	FailOnError(err)
}
