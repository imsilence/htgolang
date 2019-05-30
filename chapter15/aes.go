package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func AESEncrypt(cxt, key []byte) ([]byte, error) {

	// 创建加密算法块对象
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	// 按照PKCS7填充字节
	padding := blockSize - len(cxt)%blockSize
	cxt = append(cxt, bytes.Repeat([]byte{byte(padding)}, padding)...)

	// 设置初始化向量
	crypted := make([]byte, len(cxt)+aes.BlockSize)
	iv := crypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// 获取CBC加密对象
	blockMode := cipher.NewCBCEncrypter(block, iv)

	// 加密
	blockMode.CryptBlocks(crypted[aes.BlockSize:], cxt)

	return crypted, nil
}

func AESDecryprt(crypted, key []byte) ([]byte, error) {

	// 创建加密算法块对象
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 获取CBC解密对象
	blockMode := cipher.NewCBCDecrypter(block, crypted[:aes.BlockSize])

	// 解密
	cxt := make([]byte, len(crypted)-aes.BlockSize)
	blockMode.CryptBlocks(cxt, crypted[aes.BlockSize:])

	// 按照PKCS7规范移除填充的1字节
	return cxt[:len(cxt)-int(cxt[len(cxt)-1])], nil
}

func main() {
	// key字节长度可选16, 24, 32
	key := []byte("12345678901234561234567890123456")
	cxt := []byte("1234567890123456")
	crypted, err := AESEncrypt(cxt, key)
	if err == nil {
		fmt.Printf("%x\n", crypted)

		cxt, err := AESDecryprt(crypted, key)
		if err == nil {
			fmt.Println(string(cxt))
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
