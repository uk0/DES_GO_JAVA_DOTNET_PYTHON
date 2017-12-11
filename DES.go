package main

import "crypto/cipher"
import "crypto/des"
import "bytes"
import "fmt"
import COM  "github.com/bitherhq/go-bither/common"

func main(){
	key := []byte("aFefbFca")
	IV := []byte("Passw0rd")
    Encryptresult, err := DesEncrypt([]byte("Hello World!"), key,IV)
    if err != nil {
        panic(err)
	}
	fmt.Println(COM.Bytes2Hex(Encryptresult))
	
	DesDecryptresult, err := DesDecrypt([]byte("e529b5bbd2e0e3fb62eeb88ba6a22367"), key,IV)
    if err != nil {
        panic(err)
	}
	fmt.Println((string(DesDecryptresult)))


}

func DesEncrypt(origData, key []byte,IV []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		 return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, IV)
	crypted := make([]byte, len(origData))
	 // 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}


func DesDecrypt(crypted, key []byte ,IV []byte) ([]byte, error) {
	crypted = COM.Hex2Bytes(string(crypted))
	block, err := des.NewCipher(key)
	if err != nil {
		 return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, IV)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    // 去掉最后一个字节 unpadding 次
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}