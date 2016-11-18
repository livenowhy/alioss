
#!/usr/bin/env python
# encoding: utf-8

"""
@version: 0.0.1
@author: lzp
@license: Apache Licence 
@contact: livenowhy@gmail.com
@file: callServer.py
@time: 2016/3/9 13:53
"""


import hashlib
import time
import os
import re
import sys


import socket
import httplib
import base64
import md5
import urllib2
from M2Crypto import RSA
from M2Crypto import BIO
import json

from flask import Flask, url_for, request, render_template,make_response, jsonify

app = Flask(__name__)

app.debug = True  # apache 中 main 的设置无效


@app.route('/callback', methods=['POST'])
def do_post():
    pub_key_url = ''
    try:
        # user_agent = request.headers['User-Agent']
        # connection = request.headers['Connection']
        # host = request.headers['Host']
        # content_type = request.headers['Content-Type']
        # 暂时不用

        pub_key_url_base64 = request.headers['X-Oss-Pub-Key-Url']
        pub_key_url = pub_key_url_base64.decode('base64')
        url_reader = urllib2.urlopen(pub_key_url)
        pub_key = url_reader.read()
    except:
        resp = make_response(jsonify({'error': 'Not found'}), 400)
        return resp

    #get authorization
    try:
        authorization_base64 = request.headers['Authorization']
    except Exception as msg:
        resp = make_response(jsonify({'error': 'Not Found Authorization'}), 400)
        return resp

    try:
        authorization = authorization_base64.decode('base64')

        #get callback body
        content_length = request.headers['Content-Length']
        callback_body = request.stream.read()
    except Exception as msg:
        print(msg.args[0])

    #compose authorization string
    auth_str = ''
    pos = request.path.find('?')
    if -1 == pos:
        auth_str = request.path + '\n' + callback_body
    else:
        auth_str = urllib2.unquote(request.path[0:pos]) + request.path[pos:] + '\n' + callback_body

    print("lzp auth_str --> ")
    print auth_str
    print("lzp auth_str 11--> ")
    print(callback_body)

    body_str = urllib2.unquote(callback_body)

    #verify authorization
    auth_md5 = md5.new(auth_str).digest()
    bio = BIO.MemoryBuffer(pub_key)
    rsa_pub = RSA.load_pub_key_bio(bio)
    try:
        result = rsa_pub.verify(auth_md5, authorization, 'md5')
    except Exception as msg:
        result = False
        print msg.message

    if not result:

        resp = make_response(jsonify({'error': 'Not found'}), 400)
        return resp

    #do something accoding to callback_body

    resp_body = '{"Status":"OK"}'


    resp_body_dict = json.loads(resp_body)
    resp_body_json = jsonify(resp_body_dict)


    try:
        resp = make_response(resp_body_json, 200)
    except Exception as msg:
        print(msg.args[0])
        resp = make_response(jsonify({'error': 'hava a error'}), 400)
    finally:
        return resp


def run_server():
    app.run(debug=True, port=8000, host='0.0.0.0', threaded=True)


if __name__ == '__main__':
    run_server()