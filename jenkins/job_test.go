package jenkins

import (
	"testing"
	"os"
)


func TestAPI_GetJobs(t *testing.T) {
	api := API{
		JenkinsToken: os.Getenv("JENKINS_API_TOKEN"),
		JenkinsHost:  os.Getenv("JENKINS_HOST"),
	}
	_, err := api.GetJobs()
	if err != nil {
		t.Error("Test do not passed")
	} else {
		t.Log("PASS")
	}
}

func TestAPI_GetJobByName(t *testing.T) {

}
