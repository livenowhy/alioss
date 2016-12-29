FROM alpine:3.4

# ensure local python is preferred over distribution python
ENV PATH /usr/local/bin:$PATH

# http://bugs.python.org/issue19846
# > At the moment, setting "LANG=C" on a Linux system *fundamentally breaks Python 3*, and that's not OK.
ENV LANG C.UTF-8

# install ca-certificates so that HTTPS works consistently
# the other runtime dependencies for Python are installed later
RUN apk add --no-cache ca-certificates

ADD OssServer /oss_callback/

WORKDIR /oss_callback
ADD ./conf/key.yml /oss_callback/conf/key.yml
# 测试代码
#ADD ./test/test /oss_callback/


ENTRYPOINT ["/oss_callback/OssServer"]
CMD ["--v=2","--alsologtostderr","/oss_callback/conf/key.yml"]
EXPOSE 8765


# docker run -d -p 8765:8765 index.boxlinker.com/boxlinker/oss_callback:latest
