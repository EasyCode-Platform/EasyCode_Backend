# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.18 as builder

# 启用go module
# ENV GO111MODULE=off \
ENV  GOPROXY=https://goproxy.cn,direct

COPY . .


WORKDIR /app


# 指定OS等，并go build
# RUN GOOS=linux GOARCH=x86_64 go build  ./src/cmd/easycode-authenticator-backend/main.go \
#     &&  GOOS=linux GOARCH=x86_64  go build  ./src/cmd/easycode-authenticator-backend-internal/main.go
RUN make build

# # 由于我不止依赖二进制文件，还依赖views文件夹下的html文件还有assets文件夹下的一些静态文件
# # 所以我将这些文件放到了publish文件夹
# RUN mkdir publish && cp toc-generator publish && \
#     cp -r views publish && cp -r assets publish

# 运行阶段指定scratch作为基础镜像
FROM alpine

WORKDIR /app

# 将上一个阶段publish文件夹下的所有文件复制进来
# COPY --from=builder /app/publish .

# 指定运行时环境变量
ENV GIN_MODE=release 
