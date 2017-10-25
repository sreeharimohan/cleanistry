package main

// ReleaseCandidate ...
type ReleaseCandidate struct {
	AppRpmArtifact string `json:"appRpmArtifact"`
	ClusterName    string `json:"clusterName"`
	Date           string `json:"date"`
	DockerImage    string `json:"dockerImage"`
	OpsRpmArtifact string `json:"opsRpmArtifact"`
	ServiceName    string `json:"serviceName"`
	UserEmail      string `json:"userEmail"`
	Username       string `json:"username"`
}

// IsTagExcemptedFromDeletion ...
func IsTagExcemptedFromDeletion(image string, tag string) bool {
	var rcs []*ReleaseCandidate
	_, _, err := Get("http://localhost:4000/data/getRC/"+image+"/"+tag, false, &rcs)
	FailOnError(err)
	if len(rcs) == 1 {
		return true
	}
	return false
}
