package lib

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func httpGet(url string, header map[string]string) (content string) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Fatal Error:", err.Error())
		os.Exit(0)
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}

	response, err := client.Do(req)
	defer response.Body.Close()

	for _, cookie := range response.Cookies() {
		fmt.Println("cookie:", cookie)
	}

	body, err:=ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Get response error", err)
		os.Exit(0)
	}
	content = string(body)
	return
}
