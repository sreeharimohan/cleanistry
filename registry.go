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
func GetAllReposInRegistry() ([]string, error) {
	var catalogResponse CatalogResponse
	_, _, err := Get("http://"+GetDockerHostURL()+"/v2/_catalog?n="+GetCatalogLimit(), false, &catalogResponse)
	if err != nil {
		return []string{}, err
	}
	return catalogResponse.Repositories, nil
}

// GetListOfTagsForRepo ...
func GetListOfTagsForRepo(repo string) ([]string, error) {
	var tagListResponse TagListResponse
	_, _, err := Get("http://"+GetDockerHostURL()+"/v2/"+repo+"/tags/list", false, &tagListResponse)
	if err != nil {
		return []string{}, err
	}
	return tagListResponse.Tags, nil
}

// GetTagCreatedDate ...
func GetTagCreatedDate(repo string, tag string) (string, error) {
	var tagDetailsResponse TagDetailsResponse
	var hs HistorySet
	_, _, err := Get("http://"+GetDockerHostURL()+"/v2/"+repo+"/manifests/"+tag, false, &tagDetailsResponse)
	if err != nil {
		return "", err
	}
	objectString := strings.Replace(tagDetailsResponse.History[0].V1Compatibility, "\\", "", -1)
	// log.Println(objectString)
	json.Unmarshal([]byte(objectString), &hs)
	// log.Println(hs)
	return hs.Created, nil
}

// GetContentDigest ...
func GetContentDigest(repo string, tag string) (string, error) {
	_, res, err := Get("http://"+GetDockerHostURL()+"/v2/"+repo+"/manifests/"+tag, true, nil)
	if err != nil {
		return "", err
	}
	return res.Header["Docker-Content-Digest"][0], nil
}

// DeleteDigest ...
func DeleteDigest(repo string, digest string) (int, error) {
	var data struct{}
	resCode, err := Delete("http://"+GetDockerHostURL()+"/v2/"+repo+"/manifests/"+digest, true, data, nil)
	if err != nil {
		return resCode, err
	}
	return resCode, nil
}
