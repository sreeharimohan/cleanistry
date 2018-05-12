package main

import "log"

// // ReleaseCandidate ...
// type ReleaseCandidate struct {
// 	IsTagExcempt bool `json:"isTagExcempt"`
// }

// IsTagExcemptedFromDeletion ...
func IsTagExcemptedFromDeletion(excemptedTags []string, tag string) (bool, error) {
	if tag == "latest" && KeepLatestTag() {
		log.Println("Excempt since it is the latest tag")
		return true, nil
	}
	if IsImageTagExcemptionListPresent() {
		// var rcs *ReleaseCandidate
		// _, _, err := Get("http://"+GetImageTagExcemptionListAPI()+"/"+image+"/"+tag, false, &rcs)
		// if err != nil {
		// 	log.Println("API error")
		// 	log.Fatal(err)
		// }
		// log.Printf("Excempt response from API: %v", rcs.IsTagExcempt)
		// return rcs.IsTagExcempt, nil
		for _, singleExcemptedTag := range excemptedTags {
			if singleExcemptedTag == tag {
				return true, nil
			}
		}
	}
	return false, nil
}
