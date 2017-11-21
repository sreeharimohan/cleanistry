package main

import (
	"testing"
)

func TestIsImageTagOutdated(t *testing.T) {
	CheckAndGetConfigs()
	_, err := IsImageTagOutdated("imageName", "tagName1")
	FailOnError(err)
}
