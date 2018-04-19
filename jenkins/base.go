// Jenkins Golang SDK
package jenkins

import (
	"log"
	"os"
	"lib"
)

type API struct {
	Logger       *log.Logger
	JenkinsToken string
	JenkinsHost  string
}

func (api *API) Printf(format string, v ...interface{}) {
	api.Logger = log.New(os.Stderr, "[Jenkins SDK] "+ lib.Now()+ " ", 0)
	if api.Logger != nil {
		api.Logger.Printf(format, v...)
	}
}
