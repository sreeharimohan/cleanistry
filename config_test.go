package main

import (
	"log"
	"os"
	"testing"
)

func TestIsHostURLPresent(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URLS", "dockerhub.myntra.com:5001")
	CheckAndGetConfigs()
}

func TestIsDockerHostReachable(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "dockerhub.myntra.com:5001")
	CheckAndGetConfigs()
}

func TestForDefaultCatalogLimit(t *testing.T) {
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "dockerhub.myntra.com:8080")
	CheckAndGetConfigs()
	if GetCatalogLimit() != "50000" {
		log.Fatal("Default value for Catalog Limit not picked up")
	}
}
