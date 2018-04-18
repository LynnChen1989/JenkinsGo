package jenkins

import (
	"lib"
	"encoding/json"
)

type TaskDetail struct {
	Name  string
	Url   string
	Color string
}

type TaskItems struct {
	FullName string       `json:"fullName"`
	Tasks    []TaskDetail `json:"jobs"`
}

func GetTaskUrl(token string, host string, job string) (url string) {
	url = "http://" + token + "@" + host + "/job/" + job + "/api/json?pretty=true"
	return
}

func (api *API) GetJobTasks(name string) (tasks []TaskDetail, err error) {
	// Get one job's task list, require args job name
	url := GetTaskUrl(api.JenkinsToken, api.JenkinsHost, name)
	api.Printf("request task url: %s", url)
	header := map[string]string{}
	data := lib.HttpGet(url, header)
	var ti TaskItems
	err = json.Unmarshal([]byte(data), &ti)
	if err != nil {
		api.Printf("Error: %s", err)
	}
	tasks = ti.Tasks
	return
}

func (api *API) GetJobTaskByName(jobName string, taskName string) (jt map[string]interface{}, err error) {
	/*
	Get jobs's task, require args, job name, task name, jt means: job task
	fmt.Println(data["job"])
	fmt.Println(data["task"].(jenkins.TaskDetail))
	*/
	tasks, _ := api.GetJobTasks(jobName)
	for _, v := range tasks {
		taskInfo := TaskDetail{
			Name:  v.Name,
			Url:   v.Url,
			Color: v.Color,
		}
		if v.Name == taskName {
			jt = map[string]interface{}{
				"job":  jobName,
				"task": taskInfo,
			}
		}
	}
	return
}
