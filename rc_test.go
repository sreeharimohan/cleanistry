package main

import "testing"
import "os"

func TestIsTagExcemptedFromDeletion(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "dockerhub.myntra.com:5000")
	os.Setenv("CLEANISTRY_IMAGE_TAG_EXCEMPTION_TEST_API", "dockins.myntra.com/data/getRC")
	CheckAndGetConfigs()
	_, err := IsTagExcemptedFromDeletion("apigateway-master", "20170908142341")
	FailOnError(err)
}

func TestLatestTag(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "dockerhub.myntra.com:5000")
	os.Setenv("CLEANISTRY_IMAGE_TAG_EXCEMPTION_TEST_API", "dockins.myntra.com/data/getRC")
	CheckAndGetConfigs()
	isExcempt, err := IsTagExcemptedFromDeletion("j7t7", "latest")
	if !isExcempt || err != nil {
		t.Fail()
	}
}
