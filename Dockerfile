FROM golang:latest

ENV GOPROXY=https://goproxy.cn,direct \
    profile=prod

WORKDIR $GOPATH/src/github.com/ruomu/go-gin-template
COPY . $GOPATH/src/github.com/ruomu/go-gin-template
RUN go build .

EXPOSE 8001
ENTRYPOINT ["./go-gin-template"]