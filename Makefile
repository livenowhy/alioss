all: container

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o OssServer

container: build
	docker build -t index.boxlinker.com/boxlinker/oss_callback:1.0.1 .

#docker build --no-cache -t index.boxlinker.com/boxlinker/oss_callback:1.0.1 .