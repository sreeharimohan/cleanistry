package main

import "log"

// ReleaseCandidate ...
type ReleaseCandidate struct {
	IsTagExcempt bool `json:"isTagExcempt"`
}

// IsTagExcemptedFromDeletion ...
func IsTagExcemptedFromDeletion(image string, tag string) (bool, error) {
	if tag == "latest" && ShouldWeKeepLatestTag() {
		log.Println("Excempt since it is the latest tag")
		return true, nil
	}
	if IsImageTagExcemptionAPIPresent() {
		var rcs *ReleaseCandidate
		_, _, err := Get("http://"+GetImageTagExcemptionTestAPI()+"/"+image+"/"+tag, false, &rcs)
		if err != nil {
			log.Println("API error")
			log.Fatal(err)
		}
		log.Printf("Excempt response from API: %v", rcs.IsTagExcempt)
		return rcs.IsTagExcempt, nil
	}
	log.Println("API not present")
	return false, nil
}
