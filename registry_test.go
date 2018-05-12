package main

import (
	"testing"
)

func TestGetTagCreatedDate(t *testing.T) {
	CheckAndGetConfigs()
	_, err := GetTagCreatedDate("imageName", "tagName1")
	FailOnError(err, "Step: Test - Tag Create Date")
}

func TestGetListOfTagsForRepo(t *testing.T) {
	CheckAndGetConfigs()
	_, err := GetListOfTagsForRepo("imageName")
	FailOnError(err, "Step: Test - List of Tags")
}

func TestGetContentDigest(t *testing.T) {
	CheckAndGetConfigs()
	_, err := GetContentDigest("imageName", "tagName1")
	FailOnError(err, "Step: Test - Get Content Digest")
}
