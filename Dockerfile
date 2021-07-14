FROM golang:latest

ENV GOPROXY=https://goproxy.cn,direct \
    profile=prod

WORKDIR $GOPATH/src/github.com/ruomu/ticket-crawler
COPY . $GOPATH/src/github.com/ruomu/ticket-crawler
RUN go build .

EXPOSE 8001
ENTRYPOINT ["./ticket-crawler"]