package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func main() {

	data := []byte("i'm kk")
	var crypted []byte

	// 读取公钥文件进行加密
	cxt, err := ioutil.ReadFile("public.key.pem")
	if err == nil {
		// 解析公钥内容
		block, _ := pem.Decode(cxt)
		if block != nil {
			// 获取公钥
			publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
			if err == nil {
				// 加密数据
				crypted, _ = rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(err)
	}
	fmt.Println(crypted)

	//读取私钥文件进行解密
	cxt, err = ioutil.ReadFile("private.key.pem")
	if err == nil {
		// 解析私钥内容
		block, _ := pem.Decode(cxt)
		if block != nil {
			// 获取私钥
			privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
			if err == nil {
				// 解密数据
				data, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, crypted)
				if err == nil {
					fmt.Println(string(data))
				}
			} else {
				fmt.Println(err)
			}
		}
	}

	var signature []byte

	//读取私钥文件对数据进行签名
	cxt, err = ioutil.ReadFile("private.key.pem")
	if err == nil {
		// 解析私钥内容
		block, _ := pem.Decode(cxt)
		if block != nil {
			// 获取私钥
			privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
			if err == nil {
				// 计算数据hash值
				hash := sha256.Sum256(data)
				// 对数据进行签名
				signature, _ = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
			} else {
				fmt.Println(err)
			}
		}
	}

	fmt.Println(signature)

	// 读取公钥文件对数据进行验证
	cxt, err = ioutil.ReadFile("public.key.pem")
	if err == nil {
		// 解析公钥内容
		block, _ := pem.Decode(cxt)
		if block != nil {
			// 获取公钥
			publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
			if err == nil {
				// 计算数据hash值
				err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
				fmt.Println(err)
			} else {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println(err)
	}

}
