package main

import "log"
import "github.com/robfig/cron"

func main() {
	FailOnError(CheckAndGetConfigs(), "Step: Get Config")
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
	<-make(chan struct{})
}

func run() {
	tagsDeleted := 0
	// Getting all repositories from a private docker hub
	allRepos, err := GetAllReposInRegistry()
	FailOnError(err, "Step: Get All Repos")
	err = GetExcemptedTagsList()
	FailOnError(err, "Step: Get Excempted Tags")
	for index, singleRepo := range allRepos {
		excemptedTags := GetExcemptedTagsForImage(singleRepo)
		allTags, err := GetListOfTagsForRepo(singleRepo)
		if err != nil {
			log.Printf("Unable to get list of tags for repo %s", singleRepo)
			continue
		}
		for _, singleTag := range allTags {
			log.Printf("%s:%s", singleRepo, singleTag)
			isExcempt, err := IsTagExcemptedFromDeletion(excemptedTags, singleTag)
			if err != nil {
				log.Printf("Unable to determine if tag %s is excempt for repo %s. Error : %s", singleTag, singleRepo, err.Error())
				continue
			}
			if !isExcempt {
				isOutdated, err := IsImageTagOutdated(singleRepo, singleTag)
				if err != nil {
					log.Printf("Unable to determine if tag %s is outdate for repo %s", singleTag, singleRepo)
					continue
				}
				if isOutdated {
					contentDigest, err := GetContentDigest(singleRepo, singleTag)
					if err != nil {
						log.Printf("Unable to get content digest for tag %s in repo %s", singleTag, singleRepo)
					} else {
						err = DeleteDigest(singleRepo, contentDigest)
						if err != nil {
							log.Printf("Unable to delete digest for tag %s in repo %s", singleTag, singleRepo)
						} else {
							tagsDeleted++
							log.Printf(">>>>>> Deletion Done <<<<<< Total number of tags deleted: %d", tagsDeleted)
						}
					}
				}
			}
		}
		log.Printf("Checking if repository index %d is divisible by 5", index)
		// Running garbage collection after every 5 repositories
		if index%5 == 0 {
			log.Printf("running garbage collection after index %d", index)
			// Run garbage-collect on registry
			FailOnError(RunRegistryGarbageCollection(), "Step: Garbage Collection")
		}
	}
}

// FailOnError ...
func FailOnError(err error, step string) {
	if err != nil {
		log.Println(step)
		log.Fatal(err)
	}
}
