package main

import (
	"testing"
)

func TestGetTagCreatedDate(t *testing.T) {
	CheckAndGetConfigs()
	_, err := GetTagCreatedDate("imageName", "tagName1")
	FailOnError(err)
}

func TestGetListOfTagsForRepo(t *testing.T) {
	CheckAndGetConfigs()
	_, err := GetListOfTagsForRepo("imageName")
	FailOnError(err)
}

func TestGetContentDigest(t *testing.T) {
	CheckAndGetConfigs()
	_, err := GetContentDigest("imageName", "tagName1")
	FailOnError(err)
}
