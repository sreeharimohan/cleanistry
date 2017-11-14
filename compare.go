package main

import (
	"log"
	"time"
)

// IsImageTagOutdated ...
func IsImageTagOutdated(image string, tag string) (bool, error) {
	dateString, err := GetTagCreatedDate(image, tag)
	if err != nil {
		return false, err
	}
	if dateString == "" {
		log.Printf("Unknown Date - %s:%s", image, tag)
		return false, nil
	}
	date, err := time.Parse(time.RFC3339Nano, dateString)
	if err != nil {
		return false, err
	}
	duration := time.Since(date)
	log.Printf("Hours since build time : %f", duration.Hours())
	if duration.Hours() > GetMaxImageLifetime() {
		return true, nil
	}
	return false, nil
}
