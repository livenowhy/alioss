#!/usr/bin/env python
# encoding: utf-8

"""
@version: 0.1
@author: liuzhangpei
@contact: liuzhangpei@126.com
@site: http://www.livenowhy.com
@time: 16/11/18 10:15
"""

import md5
import BIO

if __name__ == '__main__':
    auth_str = "ssssss"
    auth_md5 = md5.new(auth_str).digest()
    bio = BIO.MemoryBuffer(pub_key)
    # print "--> %x" % (auth_md5)
