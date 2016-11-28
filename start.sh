#!/usr/bin/env bash

make

rm -rf oss_callback.tar

docker save -o oss_callback.tar index.boxlinker.com/boxlinker/oss_callback:1.0.1

scp oss_callback.tar root@123.56.9.18:/root/
scp oss_callback.tar root@192.168.1.5:/root/
rm -rf oss_callback.tar
rm -rf OssServer