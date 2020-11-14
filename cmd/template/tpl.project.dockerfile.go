package template

func init() {
	ProjectTpl["Dockerfile"] = `# Created by template.
# Builder
FROM golang:1.15-alpine as builder
COPY / /go/src/github.com/dantin-s/{{ .ProjectName }}
WORKDIR /go/src/github.com/dantin-s/{{ .ProjectName }}
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/{{ .ProjectName }} .

# Package
FROM alpine:3.9
RUN echo "http://mirrors.aliyun.com/alpine/v3.9/main" > /etc/apk/repositories
RUN echo "http://mirrors.aliyun.com/alpine/v3.9/community" >> /etc/apk/repositories
RUN apk add -U --update --no-cache curl tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /go/bin/{{ .ProjectName }} .
#COPY /resources /resources
EXPOSE 80
CMD ["/{{ .ProjectName }}", "server"]
`
}
