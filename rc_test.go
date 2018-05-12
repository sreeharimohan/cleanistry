package main

import "testing"

func TestIsTagExcemptedFromDeletion(t *testing.T) {
	CheckAndGetConfigs()
	_, err := IsTagExcemptedFromDeletion([]string{"tagName1"}, "tagName1")
	FailOnError(err)
}

func TestLatestTag(t *testing.T) {
	CheckAndGetConfigs()
	isExcempt, err := IsTagExcemptedFromDeletion([]string{"j7t7"}, "latest")
	if !isExcempt || err != nil {
		t.Fail()
	}
}
