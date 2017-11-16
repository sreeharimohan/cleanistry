package main

import "testing"
import "os"

func TestIsTagExcemptedFromDeletion(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "private-15e67c-cleanistry.apiary-mock.com:80")
	os.Setenv("CLEANISTRY_IMAGE_TAG_EXCEMPTION_TEST_API", "private-15e67c-cleanistry.apiary-mock.com/isTagExcempt")
	CheckAndGetConfigs()
	_, err := IsTagExcemptedFromDeletion("imageName", "tagName1")
	FailOnError(err)
}

func TestLatestTag(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "private-15e67c-cleanistry.apiary-mock.com:80")
	os.Setenv("CLEANISTRY_IMAGE_TAG_EXCEMPTION_TEST_API", "private-15e67c-cleanistry.apiary-mock.com/isTagExcempt")
	CheckAndGetConfigs()
	isExcempt, err := IsTagExcemptedFromDeletion("j7t7", "latest")
	if !isExcempt || err != nil {
		t.Fail()
	}
}
