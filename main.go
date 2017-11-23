package main

import "log"
import "github.com/robfig/cron"

func main() {
	FailOnError(CheckAndGetConfigs())
	if ShouldRunOnStart() {
		run()
	}
	c := cron.New()
	err := c.AddFunc(GetRunSchedule(), func() {
		run()
	})
	if err != nil {
		log.Fatalf("Unable to start cron due to error - %s", err.Error())
	}
	c.Start()
}

func run() {
	tagsDeleted := 0
	// Getting all repositories from a private docker hub
	allRepos, err := GetAllReposInRegistry()
	FailOnError(err)
	for _, singleRepo := range allRepos {
		allTags, err := GetListOfTagsForRepo(singleRepo)
		if err != nil {
			log.Printf("Unable to get list of tags for repo %s", singleRepo)
			continue
		}
		for _, singleTag := range allTags {
			log.Printf("%s:%s", singleRepo, singleTag)
			isExcempt, err := IsTagExcemptedFromDeletion(singleRepo, singleTag)
			if err != nil {
				log.Printf("Unable to determine if tag %s is excempt for repo %s. Error : %s", singleTag, singleRepo, err.Error())
				continue
			}
			if !isExcempt {
				isOutdate, err := IsImageTagOutdated(singleRepo, singleTag)
				if err != nil {
					log.Printf("Unable to determine if tag %s is outdate for repo %s", singleTag, singleRepo)
					continue
				}
				if isOutdate {
					contentDigest, err := GetContentDigest(singleRepo, singleTag)
					if err != nil {
						log.Printf("Unable to get content digest for tag %s in repo %s", singleTag, singleRepo)
					}
					log.Println(DeleteDigest(singleRepo, contentDigest))
					log.Printf(">>>>>> Deletion Done <<<<<<")
					tagsDeleted++
					log.Printf("Total number of tags deleted: %d", tagsDeleted)
				}
			}
		}
	}
	// Run garbage-collect on registry
	FailOnError(RunRegistryGarbageCollection())
}

// FailOnError ...
func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
