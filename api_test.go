package main

import (
	"fmt"
	"log"
	"testing"
)

func TestRequestCreationError(t *testing.T) {
	_, _, err := Get("%", false, nil)
	if err == nil {
		FailOnError(fmt.Errorf("Did not get an error for a erroneous url"))
	}
}

func TestBadGateway(t *testing.T) {
	code, _, err := Get("yyguehe.cahy", false, nil)
	if err == nil {
		FailOnError(fmt.Errorf("Did not get an error for a erroneous url"))
	}
	log.Printf("%v - %v", err, code)
}

func TestBadStatusCode(t *testing.T) {
	code, _, err := Get("http://cleanistry.sreeharimohan.com:80/test", false, nil)
	if err == nil {
		FailOnError(fmt.Errorf("Did not get an error for a erroneous url"))
	}
	log.Printf("%v - %v", err, code)
}
