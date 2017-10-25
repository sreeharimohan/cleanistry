package main

import "log"

func main() {
	allRepos := GetAllReposInRegistry()
	for _, singleRepo := range allRepos {
		allTags := GetListOfTagsForRepo(singleRepo)
		for _, singleTag := range allTags {
			if !IsTagExcemptedFromDeletion(singleRepo, singleTag) {
				if IsImageTagOutdated(singleRepo, singleTag) {
					log.Printf("%s:%s", singleRepo, singleTag)
					contentDigest := GetContentDigest(singleRepo, singleTag)
					log.Println(DeleteDigest(singleRepo, contentDigest))
				}
			}
		}
	}
}

// FailOnError ...
func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
