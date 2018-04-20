package lib

import (
	"net/http"
	"os"
	"io/ioutil"
	"bytes"
	"encoding/json"
)

type Call struct {
	Url          string
	Header       map[string]string
	ReturnData   string
	Username     string
	Password     string
	PostData     interface{}
	c            http.Client
	ResponseCode int
}

func (call *Call) HttpGet() {
	client := call.c
	req, err := http.NewRequest("GET", call.Url, nil)
	if err != nil {
		Error.Println("Fatal Error:", err.Error())
		os.Exit(0)
	}
	for key, value := range call.Header {
		req.Header.Add(key, value)
	}

	if len(call.Username) > 0 && len(call.Password) > 0 {
		Info.Println("you have specified username and password, seting basic auth")
		req.SetBasicAuth(call.Username, call.Password)
	}
	response, err := client.Do(req)
	defer response.Body.Close()

	for _, cookie := range response.Cookies() {
		Info.Println("get cookie:", cookie)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Error.Println("Get response error", err)
		os.Exit(0)
	}
	call.ReturnData = string(body)
}

func (call *Call) HttpPost() {
	client := call.c
	var jsonStr []byte
	if call.PostData == nil {
		Info.Println("you post without any data")
		jsonStr = []byte(`{}`)
	} else {
		postData, _ := json.Marshal(call.PostData)
		jsonStr = []byte(postData)
	}

	Info.Println("post data: ", bytes.NewBuffer(jsonStr))
	req, err := http.NewRequest("POST", call.Url, bytes.NewBuffer(jsonStr))

	if err != nil {
		Error.Println("Fatal Error:", err.Error())
		os.Exit(0)
	}
	for key, value := range call.Header {
		req.Header.Add(key, value)
	}
	req.Header.Add("Content-Type", "application/json")
	if len(call.Username) > 0 && len(call.Password) > 0 {
		Info.Println("you have specified username and password, seting basic auth")
		req.SetBasicAuth(call.Username, call.Password)
	}
	response, err := client.Do(req)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Error.Println("Get response error", err)
		os.Exit(0)
	}
	Info.Println("response status:", response.Status)
	Info.Println("response header:", response.Header)
	call.ReturnData = string(body)
}
