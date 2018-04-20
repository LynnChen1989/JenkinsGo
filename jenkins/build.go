package jenkins

import (
	"lib"
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
func (api *API) GetBuildUrl(jobName string, taskName string) (url string) {
	jobTask, _ := api.GetJobTaskByName(jobName, taskName)
	task := jobTask["task"]
	if task == nil {
		lib.Error.Printf("job:%s,task:%s can not get task url", jobName, taskName)
		panic("exit")
	}
	url = task.(TaskDetail).Url + "build"
	return
}

// start build a job
func (api *API) StartBuild(jobName string, taskName string, args ...interface{}) (content string) {
	url := api.GetBuildUrl(jobName, taskName)
	lib.Info.Printf("request build url: %s", url)
	header := map[string]string{}
	call := lib.Call{
		Url:      url,
		Header:   header,
		Username: api.JenkinsUser,
		Password: api.JenkinsToken,
	}
	/*
		如果是带参数的build， 那么需要传递对应的参数进去， jenkins build的参数格式很固定
		{"parameter":[{"name":"name", "value":"chenlin"}, {"name":"age", "value":28}, {"name":"sex", "value":"male"}]}
	*/
	for _, arg := range args {
		switch arg.(type) {
		case Parameters:
			call.PostData = args
		default:
			lib.Error.Printf("I dont kown what fuck type.\n")
		}
	}
	call.HttpPost()
	content = call.ReturnData
	return
}

func (api *API) StopBuild() {

}

func (api *API) GetBuildsFromQueue() {

}

func (api *API) CancelBuildFromQueue() {

}

func (api *API) GetBuildByNumber() {

}
