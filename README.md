# Beego-sample
beego框架的学习
#如何安装[beego](https://beego.me/quickstart)

1. 设置环境变量(windows)
    
   `set GOPATH=XXX(path路径)`

2. 安装beego
 
   **请安装git**
 
    `go get github.com/astaxie/beego`

3. 安装bee工具
  
    `go get github.com/beego/bee`

4. 将bee工具添加到环境中
    
    `set path=%PATH%;%GOPATH%/bin`

5. 创建beego项目
  
    + `bee api 项目名` 为beego的api项目
    + `bee new 项目名` 为beego的web项目

6. 运行项目
 
     `bee run `