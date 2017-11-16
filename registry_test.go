package main

import (
	"os"
	"testing"
)

func TestGetTagCreatedDate(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "private-15e67c-cleanistry.apiary-mock.com:80")
	os.Setenv("CLEANISTRY_IMAGE_TAG_EXCEMPTION_TEST_API", "private-15e67c-cleanistry.apiary-mock.com/isTagExcempted")
	CheckAndGetConfigs()
	_, err := GetTagCreatedDate("imageName", "tagName1")
	FailOnError(err)
}

func TestGetListOfTagsForRepo(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "private-15e67c-cleanistry.apiary-mock.com:80")
	os.Setenv("CLEANISTRY_IMAGE_TAG_EXCEMPTION_TEST_API", "private-15e67c-cleanistry.apiary-mock.com/isTagExcempted")
	CheckAndGetConfigs()
	_, err := GetListOfTagsForRepo("imageName")
	FailOnError(err)
}

func TestGetContentDigest(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "private-15e67c-cleanistry.apiary-mock.com:80")
	os.Setenv("CLEANISTRY_IMAGE_TAG_EXCEMPTION_TEST_API", "private-15e67c-cleanistry.apiary-mock.com/isTagExcempted")
	CheckAndGetConfigs()
	_, err := GetContentDigest("imageName", "tagName1")
	FailOnError(err)
}
