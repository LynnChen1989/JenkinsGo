// Jenkins Golang SDK
package jenkins

import "lib"

type API struct {
	JenkinsUser  string
	JenkinsToken string
	JenkinsHost  string
	C            lib.Call
}

func (api *API) ApiCall(url string, method string, header map[string]string, args ...interface{}) {
	api.C = lib.Call{
		Url:      url,
		Header:   header,
		Username: api.JenkinsUser,
		Password: api.JenkinsToken,
	}
	if method == "GET" {
		api.C.HttpGet()
	} else if method == "POST" {
		for _, arg := range args {
			switch arg.(type) {
			case Parameters:
				api.C.PostData = args
			default:
				lib.Error.Printf("I dont kown what fuck type.\n")
			}
		}
		api.C.HttpPost()
	} else {
		panic("not supported http request method")
	}
}
