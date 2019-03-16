protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. order.proto

protoc --go_out=plugins=grpc:. order.proto



abigen --abi hash.abi --pkg main --type HashDice --out hashdice.go



go mod init example.com/m

go mod tidy

其实还新增了 GOPROXY 环境变量。如果设置了该变量，下载源代码时将会通过这个环境变量设置的代理地址，而不再是以前的直接从代码库下载。这无疑对我等无法科学上网的开发良民来说是最大的福音。

更可喜的是，goproxy.io 这个开源项目帮我们实现好了我们想要的。该项目允许开发者一键构建自己的 GOPROXY 代理服务。同时，也提供了公用的代理服务 https://goproxy.io，我们只需设置该环境变量即可正常下载被墙的源码包了：

 set GOPROXY https://goproxy.io
 export GOPROXY=https://goproxy.io 
 
 
 
不过，需要依赖于 go module 功能。可通过 export GO111MODULE=on 开启 MODULE。

如果项目不在 GOPATH 中，则无法使用 go get ...，但可以使用 go mod ... 相关命令。

也可以通过置空这个环境变量来关闭，export GOPROXY=。




# Hello World

This is hello world using micro

## Contents

- main.go - is the main definition of the service, handler and client
- proto - contains the protobuf definition of the API

## Dependencies

Install the following

- [micro](https://github.com/micro/micro)
- [protoc-gen-micro](https://github.com/micro/protoc-gen-micro)

## Run Service

```shell
go run main.go
```

## Query Service

```
micro call greeter Greeter.Hello '{"name": "John"}'
```
