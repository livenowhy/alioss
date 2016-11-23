FROM busybox
ADD OssServer /oss_callback/

WORKDIR /oss_callback
ENTRYPOINT ["/oss_callback/OssServer"]
CMD ["--v=2","--alsologtostderr"]
EXPOSE 5001
