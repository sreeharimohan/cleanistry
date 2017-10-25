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
func IsTagExcemptedFromDeletion(image string, tag string) (bool, error) {
	if tag == "latest" {
		return true, nil
	}
	var rcs []*ReleaseCandidate
	_, _, err := Get("http://localhost:4000/data/getRC/"+image+"/"+tag, false, &rcs)
	if err != nil {
		return true, err
	}
	if len(rcs) == 1 {
		return true, nil
	}
	return false, nil
}
