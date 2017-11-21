package main

import (
	"log"
	"os"
	"testing"
)

func TestIsHostURLPresent(t *testing.T) {
	CheckAndGetConfigs()
}

func TestIsDockerHostReachable(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "unavailable.host.com:80")
	err := CheckAndGetConfigs()
	if err == nil {
		t.Fail()
	}
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "cleanistry.sreeharimohan.com:80")
}

func TestForDefaultCatalogLimit(t *testing.T) {
	CheckAndGetConfigs()
	if GetCatalogLimit() != "50000" {
		log.Fatal("Default value for Catalog Limit not picked up")
	}
}
