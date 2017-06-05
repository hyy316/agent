#agent
##启动服务端服务
    go build server.go
    go run server.go
    
##启动agent端
    go build cyclereport.go
    go run cyclereport.go
    
可以在服务端看到客户端周期性上报的信息，可以在agent端看到服务端反馈的信息。