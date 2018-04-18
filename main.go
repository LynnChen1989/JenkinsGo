package main

import (
	"jenkins"
	"os"
	"fmt"
)

func main() {
	api := jenkins.API{
		JenkinsToken: os.Getenv("JENKINS_API_TOKEN"),
		JenkinsHost:  os.Getenv("JENKINS_HOST"),
	}
	data, _ := api.GetJobTaskByName("628", "628_compile")
	fmt.Println(data["job"])
	fmt.Println(data["task"].(jenkins.TaskDetail))

}
