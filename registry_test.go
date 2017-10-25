package main

import "testing"

func TestGetTagCreatedDate(t *testing.T) {
	_, err := GetTagCreatedDate("absolut-saloni_vb", "0910506-1ba49d7")
	FailOnError(err)
}

func TestGetListOfTagsForRepo(t *testing.T) {
	_, err := GetListOfTagsForRepo("absolut-saloni_vb")
	FailOnError(err)
}

func TestGetContentDigest(t *testing.T) {
	_, err := GetContentDigest("absolut-saloni_vb", "0910506-1ba49d7")
	FailOnError(err)
}
