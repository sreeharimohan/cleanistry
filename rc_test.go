package main

import "testing"

func TestIsTagExcemptedFromDeletion(t *testing.T) {
	IsTagExcemptedFromDeletion("apigateway-master", "20170908142341")
	t.Fail()
}

func TestLatestTag(t *testing.T) {
	isExcempt := IsTagExcemptedFromDeletion("j7t7", "latest")
	if !isExcempt {
		t.Fail()
	}
}
