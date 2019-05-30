package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	//生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err == nil {
		//将私钥转换为pem格式
		block := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		}
		// 写入文件
		file, err := os.Create("private.key.pem")
		if err == nil {
			defer file.Close()
			if err := pem.Encode(file, block); err != nil {
				fmt.Println(err)
			}
		}

		//获取公钥并转化为pem格式
		block = &pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
		}
		// 写入文件
		file, err = os.Create("public.key.pem")
		if err == nil {
			defer file.Close()
			if err := pem.Encode(file, block); err != nil {
				fmt.Println(err)
			}
		}
	}
}
