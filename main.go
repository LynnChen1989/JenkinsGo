package main

import (
	"jenkins"
	"os"
	"fmt"
)

func main() {
	api := jenkins.API{
		JenkinsUser:  os.Getenv("JENKINS_API_USER"),
		JenkinsToken: os.Getenv("JENKINS_API_TOKEN"),
		JenkinsHost:  os.Getenv("JENKINS_HOST"),
	}
	//para := map[string]string{
	//	"name": "chenlin",
	//	"age":  "18",
	//	"sex":  "male",
	//}
	//api.StartBuildWithParameters("snake", "fuck_args", para)
	api.GetBuildByNumber("snake", "fuck_args", 18)
	fmt.Println(api.C.ResponseData)
	//f.Println(api.C.ResponseData, api.C.ResponseStatus, api.C.ResponseHeader)
}
