package main

import (
	"jenkins"
	"os"
)

func main() {
	api := jenkins.API{
		JenkinsUser:  os.Getenv("JENKINS_API_USER"),
		JenkinsToken: os.Getenv("JENKINS_API_TOKEN"),
		JenkinsHost:  os.Getenv("JENKINS_HOST"),
	}
	para := map[string]string{
		"name": "chenlin",
		"age":  "18",
		"sex":  "male",
	}
	api.StartBuild("snake", "fuck")
	api.StartBuildWithParameters("snake", "fuck_args", para)
}
