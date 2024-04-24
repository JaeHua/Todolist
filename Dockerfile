FROM golang:1.21.8

# 维护人信息
MAINTAINER jianglele

# 工作目录，即执行go命令的目录
WORKDIR $GOPATH/src/gin

# 将本地内容添加到镜像指定目录
ADD . $GOPATH/src/gin

# 设置开启go mod
RUN go env -w GO111MODULE=auto
# 设置go代理
RUN go env -w GOPROXY=https://goproxy.cn,direct
# 构建go应用
RUN go build -mod=mod main.go

# 指定镜像内部服务监听的端口
EXPOSE 3030

# 镜像默认入口命令，即go编译后的可执行文件
ENTRYPOINT ["./main"]