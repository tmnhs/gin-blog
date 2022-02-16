FROM golang:1.16 as builder

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

#FROM golang:latest
#
#MAINTAINER tmnhs
#
#RUN go env -w GO111MODULE=on
#RUN go env -w GOPROXY=https://goproxy.cn,direct
#
#WORKDIR $GOPATH/src/myblog
#COPY . $GOPATH/src/myblog
#
#RUN go build -o myblog .
#
#EXPOSE 80
#
#CMD ["./myblog"]
