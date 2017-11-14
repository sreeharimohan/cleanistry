package main

import (
	"os"
	"testing"
)

func TestIsImageTagOutdated(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "dockerhub.myntra.com:5000")
	os.Setenv("CLEANISTRY_IMAGE_TAG_EXCEMPTION_TEST_API", "dockins.myntra.com/data/getRC")
	_, err := IsImageTagOutdated("absolut-saloni_vb", "0910506-1ba49d7")
	FailOnError(err)
}
