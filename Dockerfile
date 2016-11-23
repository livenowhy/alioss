FROM busybox
ADD OssServer /oss_callback/

WORKDIR /oss_callback
ENTRYPOINT ["/oss_callback/OssServer"]
EXPOSE 8765
