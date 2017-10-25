package main

import (
	"log"
	"time"
)

// IsImageTagOutdated ...
func IsImageTagOutdated(image string, tag string) bool {
	dateString := GetTagCreatedDate(image, tag)
	if dateString == "" {
		log.Printf("Unknown Date - %s:%s", image, tag)
		return false
	}
	date, err := time.Parse(time.RFC3339Nano, dateString)
	FailOnError(err)
	duration := time.Since(date)
	if duration.Hours() > 720 {
		return true
	}
	return false
}
