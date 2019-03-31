FROM alpine
LABEL MAINTAINER="heixiushamao@gmail.com"

ENV DEFAULT_CONFIG config-production

ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone
RUN apk add --no-cache tzdata

##CA证书，https请求
RUN apk add --no-cache ca-certificates

WORKDIR /go
COPY . /go

ENTRYPOINT ./main -c=$DEFAULT_CONFIG