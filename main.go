package main

import (
	"jenkins"
	"os"
	"fmt"
)

func main() {
	api := jenkins.API{
		JenkinsUser: os.Getenv("JENKINS_API_USER"),
		JenkinsToken: os.Getenv("JENKINS_API_TOKEN"),
		JenkinsHost:  os.Getenv("JENKINS_HOST"),
	}
	data, _ := api.GetJobs()
	fmt.Println(data)
}
