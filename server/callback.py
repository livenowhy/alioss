#!/usr/bin/env python
# encoding: utf-8

"""
@version: 0.1
@author: liuzhangpei
@contact: liuzhangpei@126.com
@site: http://www.livenowhy.com
@time: 16/11/17 16:26
"""

import socket
import httplib
import base64
import md5
import urllib2
from BaseHTTPServer import BaseHTTPRequestHandler, HTTPServer


class MyHTTPRequestHandler(BaseHTTPRequestHandler):

    def do_POST(self):
        #get public key
        pub_key_url = ''
        try:
            pub_key_url_base64 = self.headers['x-oss-pub-key-url']
            pub_key_url = pub_key_url_base64.decode('base64')
            url_reader = urllib2.urlopen(pub_key_url)
            pub_key = url_reader.read()
        except:
            print 'pub_key_url : ' + pub_key_url
            print 'Get pub key failed!'
            self.send_response(400)
            self.end_headers()
            return

        #get authorization
        authorization_base64 = self.headers['authorization']
        authorization = authorization_base64.decode('base64')

        #get callback body
        content_length = self.headers['content-length']
        callback_body = self.rfile.read(int(content_length))

        #compose authorization string
        auth_str = ''

        // filename = user - dir % 2
        F537078.gif & size = 7005 & mimeType = image % 2
        Fgif & height = 64 & width = 64
        // ?    ??  # E?J??64P

        print "self.path" # self.path
        print self.path   # /callback
        pos = self.path.find('?')
        if -1 == pos:
            auth_str = self.path + '\n' + callback_body
        else:
            auth_str = urllib2.unquote(self.path[0:pos]) + self.path[pos:] + '\n' + callback_body
        print auth_str  # /callback
                        # filename=user-dir%2F537078.gif&size=7005&mimeType=image%2Fgif&height=64&width=64

        #verify authorization
        auth_md5 = md5.new(auth_str).digest()
        print auth_md5  # 乱码
        # bio = BIO.MemoryBuffer(pub_key)
        # rsa_pub = RSA.load_pub_key_bio(bio)
        # try:
        #     result = rsa_pub.verify(auth_md5, authorization, 'md5')
        # except e:
        #     result = False
        #
        # if not result:
        #     print 'Authorization verify failed!'
        #     print 'Public key : %s' % (pub_key)
        #     print 'Auth string : %s' % (auth_str)
        #     self.send_response(400)
        #     self.end_headers()
        #     return

        #do something accoding to callback_body

        #response to OSS
        resp_body = '{"Status":"OK"}'
        self.send_response(200)
        self.send_header('Content-Type', 'application/json')
        self.send_header('Content-Length', str(len(resp_body)))
        self.end_headers()
        self.wfile.write(resp_body)

class MyHTTPServer(HTTPServer):
    def __init__(self, host, port):
        HTTPServer.__init__(self, (host, port), MyHTTPRequestHandler)


if '__main__' == __name__:
    server_ip = '0.0.0.0'
    server_port = 8765

    server = MyHTTPServer(server_ip, server_port)
    server.serve_forever()


