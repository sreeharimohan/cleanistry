package main

import (
	"log"
	"os"
	"testing"
)

func TestIsHostURLPresent(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URLS", "private-15e67c-cleanistry.apiary-mock.com:80")
	CheckAndGetConfigs()
}

func TestIsDockerHostReachable(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "private-15e67c-cleanistry.apiary-mock.com:8080")
	err := CheckAndGetConfigs()
	if err == nil {
		t.Fail()
	}
}

func TestForDefaultCatalogLimit(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "private-15e67c-cleanistry.apiary-mock.com:80")
	CheckAndGetConfigs()
	if GetCatalogLimit() != "50000" {
		log.Fatal("Default value for Catalog Limit not picked up")
	}
}
