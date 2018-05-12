package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get ...
func Get(url string, useHeader bool, response interface{}) (int, *http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 400, nil, err
	}
	if useHeader {
		req.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 502, res, err
	}
	if res.StatusCode != 200 {
		return res.StatusCode, res, fmt.Errorf("GET %s - Bad status code: %d ", url, res.StatusCode)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 500, res, err
	}
	if response != nil {
		err = json.Unmarshal(body, response)
		// log.Println(response)
		if err != nil {
			return 500, res, err
		}
	}
	return 200, res, nil
}

// Delete ...
func Delete(url string, useHeader bool, data interface{}, response interface{}) (int, error) {
	var payload *bytes.Buffer
	if data != nil {
		d, err := json.Marshal(data)
		if err != nil {
			return 500, err
		}
		payload = bytes.NewBuffer(d)
	} else {
		payload = nil
	}
	req, err := http.NewRequest("DELETE", url, payload)
	if err != nil {
		return 500, err
	}
	if useHeader {
		req.Header.Add("'Accept", "application/vnd.docker.distribution.manifest.v2+json")
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 500, err
	}
	if res.StatusCode > 202 {
		return res.StatusCode, fmt.Errorf("GET %s - Bad status code: %d ", url, res.StatusCode)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 500, err
	}
	if response != nil {
		err = json.Unmarshal(body, response)
		if err != nil {
			return 500, err
		}
	}
	return 200, nil
}
