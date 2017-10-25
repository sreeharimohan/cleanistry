package main

import "testing"

func TestGetTagCreatedDate(t *testing.T) {
	GetTagCreatedDate("absolut-saloni_vb", "0910506-1ba49d7")
	t.Fatal()
}

func TestGetListOfTagsForRepo(t *testing.T) {
	GetListOfTagsForRepo("absolut-saloni_vb")
	t.Fatal()
}

func TestGetContentDigest(t *testing.T) {
	GetContentDigest("absolut-saloni_vb", "0910506-1ba49d7")
	t.Fail()
}
