package main

import "testing"

func TestIsTagExcemptedFromDeletion(t *testing.T) {
	_, err := IsTagExcemptedFromDeletion("apigateway-master", "20170908142341")
	FailOnError(err)
}

func TestLatestTag(t *testing.T) {
	isExcempt, err := IsTagExcemptedFromDeletion("j7t7", "latest")
	if !isExcempt || err != nil {
		t.Fail()
	}
}
