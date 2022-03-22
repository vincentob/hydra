package template

func init() {
	CommonProjectFiles["Dockerfile"] = `# Created by template.
# Builder
FROM golang:1.17-alpine as builder
COPY / /go/src/github.com/vincentob/{{ .ProjectName }}
WORKDIR /go/src/github.com/vincentob/{{ .ProjectName }}
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/{{ .ProjectName }} .

# Package
FROM alpine:3.13
RUN echo "http://mirrors.aliyun.com/alpine/v3.13/main" > /etc/apk/repositories
RUN echo "http://mirrors.aliyun.com/alpine/v3.13/community" >> /etc/apk/repositories
RUN apk add -U --update --no-cache curl tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /go/bin/{{ .ProjectName }} .
#COPY /resources /resources
EXPOSE 80
CMD ["/{{ .ProjectName }}", "server"]
`
}
