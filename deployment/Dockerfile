FROM golang:1.17.10 AS builder

ENV GO111MODULE on

ENV GOPROXY https://goproxy.io

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .

FROM centos:centos7
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/ .
RUN chmod +x  main && yum -y install vim net-tools telnet wget curl && yum clean all
ENTRYPOINT [ "/bin/sh", "/app/docker-entrypoint.sh" ]