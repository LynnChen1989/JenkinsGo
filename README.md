# Jenkins Golang SDK


## 环境变量

    JENKINS_API_TOKEN="snake:19505d43c11ebc773b5689d9029f54d3"
    JENKINS_HOST="127.0.0.1:8080"
    
## 使用说明
    
    使用前请先申明JENKINS_API_TOKEN和JENKINS_HOST两个环境变量, 用于提供jenkins的url和用户认证信息
    
    获取jenkins的token使用如下链接：
    http://127.0.0.1:8080/securityRealm/user/{YOU_NAME}/configure
    
    YOU_NAME替换为你的用户名
    
    
## 方法说明
    
#### GetJobs
    获取到当前用户视图下的所有配置的job列表
    
#### GetJobByName
    根据提供的job名，获取任务的信息
    
#### GetJobs
    获取某个job下所有的task信息
    
#### GetJobByName
    获取某个job下某个task的信息