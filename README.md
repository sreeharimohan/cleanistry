[![Build Status](https://travis-ci.org/sreeharimohan/cleanistry.svg?branch=master)](https://travis-ci.org/sreeharimohan/cleanistry) [![Coverage Status](https://coveralls.io/repos/github/sreeharimohan/cleanistry/badge.svg)](https://coveralls.io/github/sreeharimohan/cleanistry) [![Go Report Card](https://goreportcard.com/badge/github.com/sreeharimohan/cleanistry)](https://goreportcard.com/report/github.com/sreeharimohan/cleanistry)

# cleanistry
Docker Registry Cleanup Cron runs as a docker container on the same machine as your v2 registry

## Constraints
* Only version v2 of the registry is supported
* Only docker version with the garbage-collect API supported

## Environment variables
* `CLEANISTRY_DOCKER_HOST_URL` - *Mandatory* environment variable which helps set the private registry to be cleaned's URL e.g. dockerhub.company.com:5000
* `CLEANISTRY_CATALOG_LIMIT` - Catalog limit to be used while getting a the list of all images. Default is 50000
* `CLEANISTRY_KEEP_LATEST_TAG` - Should we ignore all images with the latest tag?
* `CLEANISTRY_IMAGE_TAG_EXCEMPTION` - true/false. Default is true just to save unwanted cleanups
* `CLEANISTRY_IMAGE_TAG_EXCEMPTION_LIST_API` - API to fetch the image tag excemption list
* `CLEANISTRY_IMAGE_TAG_EXCEMPTION_LIST` - List of excepmted tags of images
* `CLEANISTRY_MAX_IMAGE_LIFETIME` - Default lifetime is 720 hours
* `CLEANISTRY_RUN_SCHEDULE` - Default schedule is `@daily`
* `CLEANISTRY_RUN_ON_START` - Should the cleanup be run once before doing it in schedule
* `CLEANISTRY_GARBAGE_COLLECT_COMMAND` - Default command is `docker exec registry registry garbage-collect /etc/docker/registry/config.yml`

## How to run
```
docker run -d \
-e CLEANISTRY_DOCKER_HOST_URL="dockerhub.company.com:5000" \
-e CLEANISTRY_CATALOG_LIMIT=5000 \
-e CLEANISTRY_IMAGE_TAG_EXCEMPTION=true \
-e CLEANISTRY_IMAGE_TAG_EXCEMPTION_LIST_API=releasecandidates.company.com/yourapiToFetchRCs \
sreeharimohan/cleanistry
```