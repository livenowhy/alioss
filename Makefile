all: container

build:
	rm OssServer ||  true
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o OssServer

container: build
	docker build -t index.boxlinker.com/boxlinker/oss_callback:latest .
	docker push index.boxlinker.com/boxlinker/oss_callback:latest

#docker build --no-cache -t index.boxlinker.com/boxlinker/oss_callback:latest .