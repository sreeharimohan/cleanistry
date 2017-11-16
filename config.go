package main

import (
	"fmt"
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
	RunSchedule               string  `split_words:"true" default:"@daily"`
	RunOnStart                bool    `split_words:"true" default:"false"`
}

// GlobalConfigs ...
var GlobalConfigs Configs

// CheckAndGetConfigs ...
func CheckAndGetConfigs() error {
	envconfig.Process("cleanistry", &GlobalConfigs)
	if GlobalConfigs.DockerHostURL == "" {
		return fmt.Errorf("Environment variable CLEANISTRY_DOCKER_HOST_URL not found")
	}
	if !AbleToConnect("tcp", GlobalConfigs.DockerHostURL) {
		return fmt.Errorf("Unable to connect to %s", GlobalConfigs.DockerHostURL)
	}
	if ShouldWeKeepLatestTag() {
		log.Println("Keeping all images with the 'latest' tag")
	}
	if IsImageTagExcemptionAPIPresent() {
		if GlobalConfigs.ImageTagExcemptionTestAPI == "" {
			return fmt.Errorf("Ensure an ImageTag excemption API is configured using CLEANISTRY_IMAGE_TAG_EXCEMPTION_TEST_API or set CLEANISTRY_IMAGE_TAG_EXCEMPTION to false")
		}
		address := strings.Split(GlobalConfigs.ImageTagExcemptionTestAPI, "/")[0]
		if !strings.Contains(address, ":") {
			address += ":80"
		}
		if !AbleToConnect("tcp", address) {
			return fmt.Errorf("Unable to connect to %s", address)
		}
	}
	return nil
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

// GetRunSchedule ...
func GetRunSchedule() string {
	return GlobalConfigs.RunSchedule
}

// ShouldRunOnStart ...
func ShouldRunOnStart() bool {
	return GlobalConfigs.RunOnStart
}
