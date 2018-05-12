package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

// Configs ...
type Configs struct {
	DockerHostURL             string                  `split_words:"true"`
	CatalogLimit              string                  `split_words:"true" default:"50000"`
	KeepLatestTag             bool                    `split_words:"true" default:"true"`
	ImageTagExcemptionListAPI string                  `split_words:"true"`
	ImageTagExcemptionList    []ExcemptedTagsForImage `split_words:"true" default:"[]"`
	ImageTagExcemption        bool                    `split_words:"true" default:"true"`
	MaxImageLifetime          float64                 `split_words:"true" default:"720"`
	RunSchedule               string                  `split_words:"true" default:"@daily"`
	RunOnStart                bool                    `split_words:"true" default:"false"`
	GarbageCollectCommand     string                  `split_words:"true" default:"docker exec registry registry garbage-collect /etc/docker/registry/config.yml"`
}

// GlobalConfigs ...
var GlobalConfigs Configs

// ExcemptedTagsForImage ...
type ExcemptedTagsForImage struct {
	Image string   `json:"image"`
	Tags  []string `json:"tags"`
}

// ExcemptedTags ...
var ExcemptedTags []ExcemptedTagsForImage

// CheckAndGetConfigs ...
func CheckAndGetConfigs() error {
	envconfig.Process("cleanistry", &GlobalConfigs)
	if GlobalConfigs.DockerHostURL == "" {
		return fmt.Errorf("Environment variable CLEANISTRY_DOCKER_HOST_URL not found")
	}
	if !AbleToConnect("tcp", GlobalConfigs.DockerHostURL) {
		return fmt.Errorf("Unable to connect to %s", GlobalConfigs.DockerHostURL)
	}
	if KeepLatestTag() {
		log.Println("Keeping all images with the 'latest' tag")
	}
	if IsImageTagExcemptionListPresent() {
		if GlobalConfigs.ImageTagExcemptionListAPI == "" && len(GlobalConfigs.ImageTagExcemptionList) == 0 {
			return fmt.Errorf("Ensure an ImageTag excemption list API is configured using " +
				"CLEANISTRY_IMAGE_TAG_EXCEMPTION_LIST_API OR\n set CLEANISTRY_IMAGE_TAG_EXCEMPTION to false OR\n" +
				"Add the list using CLEANISTRY_IMAGE_TAG_EXCEMPTION_LIST")
		}
		if GlobalConfigs.ImageTagExcemptionListAPI != "" && len(GlobalConfigs.ImageTagExcemptionList) != 0 {
			log.Println("Both CLEANISTRY_IMAGE_TAG_EXCEMPTION_LIST_API and CLEANISTRY_IMAGE_TAG_EXCEMPTION_LIST " +
				"have been mentioned. API overrides the List")
		}
		address := strings.Split(GlobalConfigs.ImageTagExcemptionListAPI, "/")[0]
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

// KeepLatestTag ...
func KeepLatestTag() bool {
	return GlobalConfigs.KeepLatestTag
}

// GetImageTagExcemptionListAPI ...
func GetImageTagExcemptionListAPI() string {
	return GlobalConfigs.ImageTagExcemptionListAPI
}

// IsImageTagExcemptionListPresent ...
func IsImageTagExcemptionListPresent() bool {
	return GlobalConfigs.ImageTagExcemption
}

// GetImageTagExcemptionList ...
func GetImageTagExcemptionList() []ExcemptedTagsForImage {
	return GlobalConfigs.ImageTagExcemptionList
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

// GetGarbageCollectCommand ...
func GetGarbageCollectCommand() []string {
	return strings.Fields(GlobalConfigs.GarbageCollectCommand)
}

// GetExcemptedTagsList ...
func GetExcemptedTagsList() (err error) {
	if GlobalConfigs.ImageTagExcemptionListAPI != "" {
		_, _, err = Get("http://"+GetImageTagExcemptionListAPI(), false, &ExcemptedTags)
	} else {
		ExcemptedTags = GlobalConfigs.ImageTagExcemptionList
	}
	return err
}

// GetExcemptedTagsForImage ...
func GetExcemptedTagsForImage(image string) []string {
	var excemptedTagsForImage []string
	for _, singleImage := range ExcemptedTags {
		if singleImage.Image == image {
			excemptedTagsForImage = singleImage.Tags
			break
		}
	}
	return excemptedTagsForImage
}
