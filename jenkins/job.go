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
	// 知识点，接收json的是jobs，不是Items, 也就是说是flag的名称和json的字段名称对应即可
	Jobs []ItemDetail `json:"jobs"`
}

func GetJobUrl(token string, host string) (url string) {
	url = "http://" + token + "@" + host + "/api/json?pretty=true"
	return
}

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
