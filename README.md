# Jenkins Golang SDK


## Environment

    JENKINS_API_TOKEN="snake:19505d43c11ebc773b5689d9029f54d3"
    JENKINS_HOST="127.0.0.1:8080"
    
## Usage

#### notice    
    Pls. decare two variables JENKINS_API_TOKEN, JENKINS_HOST

    JENKINS_API_TOKEN format,  "[username]:[token]"
    JENKINS_HOST format, "[HOST]:[PORT]"
    
#### how to get user token from jenkins ?
    http://127.0.0.1:8080/securityRealm/user/{YOU_NAME}/configure
    
    YOU_NAME replace your real name.
    
    
## Document

#### prescriptive
    when you read document from godoc.
    job, means user view job list;
    task, means when you click one job, you can get a list of current job;
       
#### install godoc
    go get -u golang.org/x/tools/cmd/godoc
    
#### run godoc 
    ./bin/godoc.exe -http 127.0.0.1:9999 -play -goroot ./src/

#### read document
    if you run godoc, you can view document of JenkinsGo via:

    http://127.0.0.1:9999/pkg/jenkins