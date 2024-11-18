# 使用官方的 Go 语言镜像作为基础镜像
FROM golang:1.23.3-alpine AS builder

# 设置工作目录为 /app
WORKDIR /app

# 将当前项目目录的所有文件拷贝到容器的 /app 目录中
COPY . . 

# 设置 Go 模块代理为 https://goproxy.cn（在中国加速模块下载），并下载项目的依赖
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod tidy && go mod download

# 编译 Go 项目，生成可执行文件 gin-example
RUN GOOS=linux GOARCH=amd64 go build -o gin-example

# 使用一个更小的基础镜像（Alpine）来运行应用程序
FROM alpine:latest

# 安装 tzdata 包，确保支持时区的配置
RUN apk add --no-cache tzdata

# 创建一个非 root 用户（提升安全性）
RUN adduser -D myuser

# 设置工作目录为 /app
WORKDIR /app

# 从构建阶段的镜像中拷贝编译后的二进制文件到运行镜像中
COPY --from=builder /app/gin-example /app/gin-example

# 给非 root 用户执行权限
RUN chmod +x /app/gin-example

# 暴露容器的 8080 端口，用于外部访问
EXPOSE 8080

# 切换到非 root 用户
USER myuser

# 设置容器启动时运行的命令
CMD ["/app/gin-example"]
