FROM golang:1.13 as builder

MAINTAINER Jack <jack@163.com>

WORKDIR /build
COPY . /build

ENV GOPROXY=https://goproxy.cn,direct

RUN go build -a -o basic

FROM centos:latest

WORKDIR /data

COPY --from=builder /build/basic .

VOLUME ["/data/log"]

EXPOSE 80
ENTRYPOINT ["./basic"]
