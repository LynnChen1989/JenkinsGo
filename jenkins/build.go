package jenkins

import (
	"lib"
	"strconv"
)

type Parameters struct {
	Parameter []ParameterFormat `json:"parameter"`
}

type ParameterFormat struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

var paras []ParameterFormat

func JenkinsConstructParas(args map[string]string) {
	for k, v := range args {
		para := ParameterFormat{
			Name:  k,
			Value: v,
		}
		paras = append(paras, para)
	}
}

// 封装build参数 {"parameter":[{"name":"name", "value":"chenlin"}, {"name":"age", "value":28}, {"name":"sex", "value":"male"}]}
func (api *API) GetPostParas(args map[string]string) (parameters Parameters) {
	JenkinsConstructParas(args)
	parameters = Parameters{
		Parameter: paras,
	}
	return
}

// get job task build url
func (api *API) GetBuildUrl(jobName string, taskName string, withParameters bool) (url string) {
	jobTask, _ := api.GetJobTaskByName(jobName, taskName)
	task := jobTask["task"]
	if task == nil {
		lib.Error.Printf("job:%s,task:%s can not get task url", jobName, taskName)
		panic("exit")
	}
	if withParameters == true {
		url = task.(TaskDetail).Url + "buildWithParameters"
	} else {
		url = task.(TaskDetail).Url + "build"

	}
	return
}

// warp build url when you call as: http://snake.ops.dragonest.com:8080/job/snake/job/fuck_args/buildWithParameters
func (api *API) WrapBuildWithParametersUrl(jobName string, taskName string, parameters map[string]string) (url string) {
	url = api.GetBuildUrl(jobName, taskName, true)
	var argsString string
	argsString += "/?"
	for k, v := range parameters {
		argsString += k + "=" + v + "&"
	}
	url += argsString
	return
}

// 开始构建
func (api *API) StartBuild(jobName string, taskName string, args ...interface{}) {
	url := api.GetBuildUrl(jobName, taskName, false)
	lib.Info.Printf("request build url: %s", url)
	header := map[string]string{}
	api.ApiCall(url, "POST", header, args)
}

// 开始构建（带参数）
func (api *API) StartBuildWithParameters(jobName string, taskName string, parameters map[string]string) {
	// StartBuild()是实现带参数build的一种方式，这是另一种方式
	url := api.WrapBuildWithParametersUrl(jobName, taskName, parameters)
	lib.Info.Printf("request build with parameters url: %s", url)
	header := map[string]string{}
	api.ApiCall(url, "POST", header)
}

func (api *API) GetStopBuildUrl(jobName string, taskName string, taskNumber int) (url string) {
	jobTask, _ := api.GetJobTaskByName(jobName, taskName)
	task := jobTask["task"]
	url = task.(TaskDetail).Url + strconv.Itoa(taskNumber) + "/stop"
	return
}

func (api *API) GetBuildStatusUrl(jobName string, taskName string, taskNumber int) (url string) {
	jobTask, _ := api.GetJobTaskByName(jobName, taskName)
	task := jobTask["task"]
	url = task.(TaskDetail).Url + strconv.Itoa(taskNumber) + "/api/json"
	return
}

// 终止构建
func (api *API) StopBuild(jobName string, taskName string, taskNumber int) {
	url := api.GetStopBuildUrl(jobName, taskName, taskNumber)
	lib.Info.Printf("request stop build url: %s", url)
	header := map[string]string{}
	api.ApiCall(url, "POST", header)
}

// 获取构建状态
func (api *API) GetBuildByNumber(jobName string, taskName string, taskNumber int) {
	url := api.GetBuildStatusUrl(jobName, taskName, taskNumber)
	lib.Info.Printf("request build status url: %s", url)
	header := map[string]string{}
	api.ApiCall(url, "GET", header)
}

