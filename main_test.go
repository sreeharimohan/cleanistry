package main

import (
	"net/http"
	"os"
	"testing"

	"gopkg.in/jarcoal/httpmock.v1"
)

func TestMain(m *testing.M) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "http://cleanistry.sreeharimohan.com:80/v2/imageName/manifests/tagName1",
		func(req *http.Request) (*http.Response, error) {
			singleV1Compatibility := V1Compatibility{
				V1Compatibility: "{\"created\":\"2017-11-02T15:04:05.999999999Z\"}",
			}
			v1Compatibility := make([]V1Compatibility, 1)
			v1Compatibility = append(v1Compatibility, singleV1Compatibility)
			res, _ := httpmock.NewJsonResponse(200, TagDetailsResponse{
				Name:    "imageName",
				History: v1Compatibility,
			})
			res.Header.Add("Docker-Content-Digest", "[\"thisisaveryhugecontextdigest\"]")
			return res, nil
		},
	)
	httpmock.RegisterResponder("GET", "http://cleanistry.sreeharimohan.com:80/v2/imageName/tags/list",
		func(req *http.Request) (*http.Response, error) {
			res, _ := httpmock.NewJsonResponse(200, TagListResponse{
				Name: "imageName",
				Tags: []string{
					"tagName1",
					"tagName2",
					"tagName3",
				},
			})
			return res, nil
		},
	)
	httpmock.RegisterResponder("GET", "http://cleanistry.sreeharimohan.com:80/isTagExcempted/imageName/tagName1",
		func(req *http.Request) (*http.Response, error) {
			res, _ := httpmock.NewJsonResponse(200, struct {
				IsTagExcempt bool `json:"isTagExcempt"`
			}{
				IsTagExcempt: true,
			})
			return res, nil
		},
	)
	os.Setenv("CLEANISTRY_DOCKER_HOST_URL", "cleanistry.sreeharimohan.com:80")
	os.Setenv("CLEANISTRY_IMAGE_TAG_EXCEMPTION_TEST_API", "cleanistry.sreeharimohan.com:80/isTagExcempted")
	os.Exit(m.Run())
}
