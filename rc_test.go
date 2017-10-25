package main

import "testing"

func TestIsTagExcemptedFromDeletion(t *testing.T) {
	IsTagExcemptedFromDeletion("apigateway-master", "20170908142341")
	t.Fail()
}
