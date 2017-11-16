package main

import (
	"os"
	"testing"
)

func TestIsImageTagOutdated(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "private-15e67c-cleanistry.apiary-mock.com:80")
	os.Setenv("CLEANISTRY_IMAGE_TAG_EXCEMPTION_TEST_API", "private-15e67c-cleanistry.apiary-mock.com/isTagExcempt")
	CheckAndGetConfigs()
	_, err := IsImageTagOutdated("imageName", "tagName1")
	FailOnError(err)
}
