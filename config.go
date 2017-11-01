package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Configs ...
type Configs struct {
	DockerHostURL string `split_words:"true"`
	CatalogLimit  string `split_words:"true" default:"50000"`
}

// GlobalConfigs ...
var GlobalConfigs Configs

// CheckAndGetConfigs ...
func CheckAndGetConfigs() {
	envconfig.Process("cleanistry", &GlobalConfigs)
	if GlobalConfigs.DockerHostURL == "" {
		log.Fatal("Environment variable CLEANISTRY_DOCKER_HOST_URL not found")
	}
	if !AbleToConnect("tcp", GlobalConfigs.DockerHostURL) {
		log.Fatalf("Unable to connect to %s", GlobalConfigs.DockerHostURL)
	}
}

// GetDockerHostURL ...
func GetDockerHostURL() string {
	return GlobalConfigs.DockerHostURL
}

// GetCatalogLimit ...
func GetCatalogLimit() string {
	return GlobalConfigs.CatalogLimit
}
