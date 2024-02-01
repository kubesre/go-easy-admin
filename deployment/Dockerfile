FROM golang:1.17.10 AS builder

ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .

FROM alpine:3.17.3
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8899

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --no-cache add tzdata ca-certificates curl busybox-extras && \
    cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

CMD ["./main"]