package main

import (
	"encoding/json"
	"strings"
)

// CatalogResponse ...
type CatalogResponse struct {
	Repositories []string `json:"repositories"`
}

// TagListResponse ...
type TagListResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

// HistorySet ...
type HistorySet struct {
	Created string `json:"created"`
}

// V1Compatibility ...
type V1Compatibility struct {
	V1Compatibility string `json:"v1Compatibility"`
}

// TagDetailsResponse ...
type TagDetailsResponse struct {
	Name    string            `json:"name"`
	History []V1Compatibility `json:"history"`
}

// GetAllReposInRegistry ...
func GetAllReposInRegistry() []string {
	var catalogResponse CatalogResponse
	_, _, err := Get("http://dockerhub.myntra.com:8080/v2/_catalog?n=50000", false, &catalogResponse)
	FailOnError(err)
	return catalogResponse.Repositories
}

// GetListOfTagsForRepo ...
func GetListOfTagsForRepo(repo string) []string {
	var tagListResponse TagListResponse
	_, _, err := Get("http://dockerhub.myntra.com:8080/v2/"+repo+"/tags/list", false, &tagListResponse)
	FailOnError(err)
	return tagListResponse.Tags
}

// GetTagCreatedDate ...
func GetTagCreatedDate(repo string, tag string) string {
	var tagDetailsResponse TagDetailsResponse
	var hs HistorySet
	_, _, err := Get("http://dockerhub.myntra.com:8080/v2/"+repo+"/manifests/"+tag, false, &tagDetailsResponse)
	FailOnError(err)
	objectString := strings.Replace(tagDetailsResponse.History[0].V1Compatibility, "\\", "", -1)
	// log.Println(objectString)
	json.Unmarshal([]byte(objectString), &hs)
	// log.Println(hs)
	return hs.Created
}

// GetContentDigest ...
func GetContentDigest(repo string, tag string) string {
	// acceptContent = append(acceptContent, "")
	// requestHeaders[""] = acceptContent
	_, res, err := Get("http://dockerhub.myntra.com:8080/v2/"+repo+"/manifests/"+tag, true, nil)
	FailOnError(err)
	// log.Println()
	return res.Header["Docker-Content-Digest"][0]
}

// DeleteDigest ...
func DeleteDigest(repo string, digest string) int {
	var data struct{}
	resCode, err := Delete("http://dockerhub.myntra.com:8080/v2/"+repo+"/manifests/"+digest, true, data, nil)
	FailOnError(err)
	return resCode
}
