#!/usr/bin/env bash

make

rm -rf oss_callback.tar

docker save -o oss_callback.tar index.boxlinker.com/boxlinker/oss_callback:latest


scp oss_callback.tar root@192.168.1.5:/root/
rm -rf oss_callback.tar
rm -rf OssServer