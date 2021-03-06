# Jenkins Golang SDK


## Environment
    JENKINS_API_USER=snake
    JENKINS_API_TOKEN=19505d43c11ebc773b5689d9029f54d3
    JENKINS_HOST=127.0.0.1:8080
    
## Usage

#### notice    
    Pls. decare two variables JENKINS_API_TOKEN, JENKINS_HOST

    JENKINS_API_USER, username
    JENKINS_API_TOKEN, user token
    JENKINS_HOST, "[HOST]:[PORT]"
    
#### how to get user token from jenkins ?
    http://127.0.0.1:8080/securityRealm/user/{YOU_NAME}/configure
    
    YOU_NAME replace your real name.
    
    
## Document

#### prescriptive
    when you read document from godoc.
    job, means user view job list;
    task, means subjob;
    
**So, this SDK just support two layer job layout. if you have more layer, this sdk not work well**
       
#### install godoc
    go get -u golang.org/x/tools/cmd/godoc
    
#### run godoc 
    ./bin/godoc.exe -http 127.0.0.1:9999 -play -goroot ./src/

#### read document
    if you run godoc, you can view document of JenkinsGo via:

    http://127.0.0.1:9999/pkg/jenkins
    
    
# Use case

##### get job list, 获取job列表
```test

api.GetJobs()
```
##### get job by name, 根据名称获取job
```test
api.GetJobs("snake")
```
##### get task list, 获取task列表
```test
api.GetJobTasks()
```
##### get task by name, 根据名称获取task
```test
api.GetJobTaskByName()
```
##### build, 构建task
```test
	api := jenkins.API{
		JenkinsUser:  os.Getenv("JENKINS_API_USER"),
		JenkinsToken: os.Getenv("JENKINS_API_TOKEN"),
		JenkinsHost:  os.Getenv("JENKINS_HOST"),
	}
	api.StartBuild("snake", "fuck")
```
##### build with parameters, 带参数构建

```test
	api := jenkins.API{
		JenkinsUser:  os.Getenv("JENKINS_API_USER"),
		JenkinsToken: os.Getenv("JENKINS_API_TOKEN"),
		JenkinsHost:  os.Getenv("JENKINS_HOST"),
	}
	para := map[string]string{
		"name": "chenlin",
		"age":  "18",
		"sex":  "male",
	}
	api.StartBuild("snake", "fuck")
	api.StartBuildWithParameters("snake", "fuck_args", para)
```

##### start build, 开始构建 

```
api.StartBuild()
```

##### start build with parameters, 开始构建 

```
api.StartBuildWithParameters()
```

#### stop build, 终止构建

```
api.StopBuild()
```

#### build status, 获取构建状态

```
api.GetBuildByNumber()

```