#!/usr/bin/python
# -*- coding: UTF-8 -*-

from DES import *;



def main():
    des = DES()
    Des_IV = "Passw0rd"
    encode = des.encrypt("Passw0rd",Des_IV,str("Hello World!"))
    encode2 = des.decrypt("Passw0rd",Des_IV,str("53c8927a4dd1e9b6f44c9d53af92c794"))
    print(encode)

    print(encode2)

main()