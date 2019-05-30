package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {

	addr := "http://localhost:8888/"

	params := url.Values{}
	params.Add("name", "kk")

	// 发起Head请求
	response, err := http.Head(fmt.Sprintf("%s?%s", addr, params.Encode()))
	if err == nil {
		response.Write(os.Stdout)
		fmt.Println()
		fmt.Println(response.Proto, response.Status)
		fmt.Println(response.Header)
	} else {
		fmt.Println(err)
	}

	// 发起GET请求
	response, err = http.Get(fmt.Sprintf("%s?%s", addr, params.Encode()))
	if err == nil {
		io.Copy(os.Stdout, response.Body)
		fmt.Println()
	} else {
		fmt.Println(err)
	}

	// POST提交application/json格式数据
	builder := strings.NewReader(`{"name" : "kk"}`)
	response, err = http.Post(fmt.Sprintf("%s", addr), "application/json", builder)
	if err == nil {
		io.Copy(os.Stdout, response.Body)
		fmt.Println()
	} else {
		fmt.Println(err)
	}

	// POST提交application/x-www-form-urlencoded格式数据
	response, err = http.PostForm(fmt.Sprintf("%s", addr), params)
	if err == nil {
		io.Copy(os.Stdout, response.Body)
		fmt.Println()
	} else {
		fmt.Println(err)
	}

	// 使用client提交数据

	// 创建request对象
	request, _ := http.NewRequest("POST", fmt.Sprintf("%s/cookie/", addr), nil)

	// 添加cookie
	request.AddCookie(&http.Cookie{
		Name:  "token",
		Value: "123123123xxxxxxx",
	})

	// 创建传输对象, 配置跳过ssl验证
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	// 创建client对象
	client := &http.Client{
		Transport: transport,
	}

	// 发起请求
	response, err = client.Do(request)
	if err == nil {
		response.Write(os.Stdout)
	}

}
