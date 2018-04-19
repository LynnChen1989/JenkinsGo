package jenkins

import (
	"lib"
	"encoding/json"
)

type ItemDetail struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type JenkinsItems struct {
	// point，received json filed id `jobs` not `Items`, you can just assign flag as `jobs`, not `items`
	Jobs []ItemDetail `json:"jobs"`
}

// Get list of job's url
func GetJobUrl(token string, host string) (url string) {
	url = "http://" + token + "@" + host + "/api/json?pretty=true"
	return
}

// Get all jobs of current user
func (api *API) GetJobs() (items []ItemDetail, err error) {
	url := GetJobUrl(api.JenkinsToken, api.JenkinsHost)
	api.Printf("request job url: %s", url)
	header := map[string]string{}
	data := lib.HttpGet(url, header)
	var jobs JenkinsItems
	err = json.Unmarshal([]byte(data), &jobs)
	if err != nil {
		api.Printf("Error: %s", err)
	}
	items = jobs.Jobs
	return
}

// Get one job of current user, filter by job name, if you have duplicate job name, just return last one
func (api *API) GetJobByName(name string) (job map[string]string, err error) {
	items, err := api.GetJobs()
	if err != nil {
		api.Printf("Error: %s", err)
	}
	for _, v := range items {
		if v.Name == name {
			job = map[string]string{
				"name": v.Name,
				"url":  v.Url,
			}
		}
	}
	return
}
