#!/usr/bin/python
# -*- coding: UTF-8 -*-
import sys
reload(sys)
sys.setdefaultencoding('utf-8')
from pyDes import *
from binascii import b2a_hex, a2b_hex

# For Python3, you'll need to use bytes, i.e.:
#   data = b"Please encrypt my data"
#   k = des(b"KEY", CBC, b"VI", pad=None, padmode=PAD_PKCS5)
class DES():
    def encrypt(self,key,iv,data):
        k = des(key, CBC, iv, pad=None, padmode=PAD_PKCS5)
        d = k.encrypt(data)
        return b2a_hex(d)

    def decrypt(self,key,iv,data):
        k = des(key, CBC, iv, pad=None, padmode=PAD_PKCS5)
        d = a2b_hex(data)
        return k.decrypt(d)

       # print encrypt("test_KEY",Des_IV,data)
       #
       # print decrypt("test_KEY",Des_IV,"82b40012c204887f")



