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

// Get job task url
func GetTaskUrl(host string, job string) (url string) {
	url = "http://" + host + "/job/" + job + "/api/json?pretty=true"
	return
}

// Get one job's task list, require args job name
func (api *API) GetJobTasks(name string) (tasks []TaskDetail, err error) {
	url := GetTaskUrl(api.JenkinsHost, name)
	lib.Info.Printf("request task url: %s", url)
	header := map[string]string{}
	call := lib.Call{
		Url:      url,
		Header:   header,
		Username: api.JenkinsUser,
		Password: api.JenkinsToken,
	}
	call.HttpGet()
	var ti TaskItems
	err = json.Unmarshal([]byte(call.ReturnData), &ti)
	if err != nil {
		lib.Error.Printf("Error: %s", err)
	}
	tasks = ti.Tasks
	return
}

// Get jobs's task, require args, job name, task name
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

// Disable specified task

func (api *API) DisableTask(name string) {

}

// Enable specified task
func (api *API) EnableTask(name string) {

}
