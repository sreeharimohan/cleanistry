package main

import (
	"log"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

// Configs ...
type Configs struct {
	DockerHostURL             string  `split_words:"true"`
	CatalogLimit              string  `split_words:"true" default:"50000"`
	KeepLatestTag             bool    `split_words:"true" default:"true"`
	ImageTagExcemptionTestAPI string  `split_words:"true"`
	ImageTagExcemption        bool    `split_words:"true" default:"true"`
	MaxImageLifetime          float64 `split_words:"true" default:"720"`
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
	if ShouldWeKeepLatestTag() {
		log.Println("Keeping all images with the 'latest' tag")
	}
	if IsImageTagExcemptionAPIPresent() {
		if GlobalConfigs.ImageTagExcemptionTestAPI == "" {
			log.Fatalf("Ensure an ImageTag excemption API is configured using CLEANISTRY_IMAGE_TAG_EXCEMPTION_TEST_API or set CLEANISTRY_IMAGE_TAG_EXCEMPTION to false")
		} else {
			if !AbleToConnect("tcp", strings.Split(GlobalConfigs.ImageTagExcemptionTestAPI, "/")[0]) {
				log.Fatalf("Unable to connect to %s", GlobalConfigs.ImageTagExcemptionTestAPI)
			}
		}
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

// ShouldWeKeepLatestTag ...
func ShouldWeKeepLatestTag() bool {
	return GlobalConfigs.KeepLatestTag
}

// GetImageTagExcemptionTestAPI ...
func GetImageTagExcemptionTestAPI() string {
	return GlobalConfigs.ImageTagExcemptionTestAPI
}

// IsImageTagExcemptionAPIPresent ...
func IsImageTagExcemptionAPIPresent() bool {
	return GlobalConfigs.ImageTagExcemption
}

// GetMaxImageLifetime ...
func GetMaxImageLifetime() float64 {
	return GlobalConfigs.MaxImageLifetime
}
