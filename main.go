package main

import (
	"jenkins"
	"os"
	"fmt"
)

/* {"name":"chenlin", "age": 18, "sex":"male"}
	==> [{"name":"name", "value":"chenlin"}, {"name":"age", "value":28}, {"name":"sex", "value":"male"}]}
*/

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
	//paras := api.GetPostParas(para)
	data := api.StartBuild("snake", "fuck")
	fmt.Println(data)
}
