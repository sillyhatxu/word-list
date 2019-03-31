go clean
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
docker build -t word-list .
docker tag word-list:latest xushikuan/word-list:1.3
docker push xushikuan/word-list:1.3


