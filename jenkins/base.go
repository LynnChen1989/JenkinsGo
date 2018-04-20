// Jenkins Golang SDK
package jenkins

import "lib"

type API struct {
	JenkinsUser  string
	JenkinsToken string
	JenkinsHost  string
}

func (api *API) ApiCall(url string, method string, header map[string]string, args ...interface{}) (content string) {
	call := lib.Call{
		Url:      url,
		Header:   header,
		Username: api.JenkinsUser,
		Password: api.JenkinsToken,
	}
	if method == "GET" {
		call.HttpGet()
	} else if method == "POST" {
		for _, arg := range args {
			switch arg.(type) {
			case Parameters:
				call.PostData = args
			default:
				lib.Error.Printf("I dont kown what fuck type.\n")
			}
		}
		call.HttpPost()
	} else {
		panic("not supported http request method")
	}
	content = call.ReturnData
	return
}
