FROM busybox
ADD OssServer /oss_callback/

WORKDIR /oss_callback
ENTRYPOINT ["/oss_callback/OssServer"]
CMD ["--v=2","--alsologtostderr","/oss_callback/conf/key.yml"]
EXPOSE 8765
