package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	/** 定义处理器函数, 用于打印提交请求到/的request信息
	* 在控制台打印请求方式、请求URL、请求Header信息及请求body信息
	* 向客户端返回当前时间
	**/
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 请求行
		fmt.Println(request.Method, request.URL, request.Proto)

		// 请求头
		fmt.Println(request.Header)

		// 请求体字节数量
		fmt.Println(request.ContentLength)

		// 请求体
		io.Copy(os.Stdout, request.Body)

		fmt.Println()

		// 响应体
		fmt.Fprintf(writer, "%d", time.Now().Unix())
	})

	/** 定义处理器函数, 用于获取用URL(body)中的参数
	* 在控制台打印请求方式、请求URL及从URL(body)中读取的参数
	* 向客户端返回当前时间
	**/
	http.HandleFunc("/url/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method, request.URL)

		// 解析url(querystring)和body内容
		request.ParseForm()

		// 获取提交的所有参数
		fmt.Println(request.Form)

		// 根据指定名称获取参数值
		fmt.Println(request.Form.Get("id"))

		// 根据名称获取参数值(自定调用ParseForm)
		fmt.Println(request.FormValue("id"))
		// fmt.Println(request.Form)

		// 响应体
		fmt.Fprintf(writer, "%d", time.Now().Unix())
	})

	/** 定义处理器函数, 用于获取使用application/x-www-form-urlencoded编码的body参数
	* 在控制台打印请求方式、请求URL及从body中读取的参数
	* 向客户端返回当前时间
	**/
	http.HandleFunc("/urlencode/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method, request.URL)

		// 解析url(querystring)和body内容
		request.ParseForm()

		// 获取提交的所有参数
		fmt.Println(request.Form)

		// 根据指定名称获取参数值
		fmt.Println(request.Form.Get("id"))

		// 根据名称获取参数值(自定调用ParseForm)
		fmt.Println(request.FormValue("id"))

		// 获取body中提交的所有参数
		fmt.Println(request.PostForm)

		// 根据指定名称从body中获取参数值
		fmt.Println(request.PostForm.Get("name"))

		// 根据名称从body中获取参数值(自定调用ParseForm)
		fmt.Println(request.PostFormValue("name"))

		// 响应体
		fmt.Fprintf(writer, "%d", time.Now().Unix())
	})

	/** 定义处理器函数, 用于获取使用multipart/form-data编码的body参数
	* 在控制台打印请求方式、请求URL及从body中读取的参数
	* 向客户端返回当前时间
	**/
	http.HandleFunc("/multipart/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method, request.URL)

		// 解析url(querystring)和body内容
		request.ParseMultipartForm(1024)

		// 获取提交的所有参数
		fmt.Println(request.Form)
		// 根据指定名称获取参数值
		fmt.Println(request.Form.Get("id"))
		// 根据名称获取参数值(自定调用ParseForm)
		fmt.Println(request.FormValue("id"))

		// 获取body中提交的所有参数
		fmt.Println(request.PostForm)
		// 根据指定名称从body中获取参数值
		fmt.Println(request.PostForm.Get("name"))
		// 根据名称从body中获取参数值(自定调用ParseForm)
		fmt.Println(request.PostFormValue("name"))

		// 获取body中提交的所有参数和文件
		fmt.Println(request.MultipartForm)
		// 获取body中提交的所有参数
		fmt.Println(request.MultipartForm.Value)
		// 获取body中提交的所有文件
		fmt.Println(request.MultipartForm.File)
		// 根据名称从body中获取文件流(自定调用ParseMultipartForm)
		if file, _, err := request.FormFile("test"); err == nil {
			defer file.Close()
			io.Copy(os.Stdout, file)
		}
		// 响应体
		fmt.Fprintf(writer, "%d", time.Now().Unix())
	})

	/** 定义处理器函数, 用于获取使用application/json编码的body参数
	* 在控制台打印请求方式、请求URL及从body中读取的参数
	* 向客户端返回当前时间
	**/
	http.HandleFunc("/json/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method, request.URL)

		// 解析url(querystring)和body内容
		request.ParseForm()

		// 获取提交的所有参数
		fmt.Println(request.Form)
		// 根据指定名称获取参数值
		fmt.Println(request.Form.Get("id"))
		// 根据名称获取参数值(自定调用ParseForm)
		fmt.Println(request.FormValue("id"))

		// 反序列化body中json字符串
		info := map[string]interface{}{}

		decoder := json.NewDecoder(request.Body)
		if err := decoder.Decode(&info); err == nil {
			fmt.Println(info)
			// 将对象序列化为字符串响应给客户端
			encoder := json.NewEncoder(writer)
			encoder.Encode(info)
		} else {
			fmt.Println(err)
		}

		// 响应体
		fmt.Fprintf(writer, "%d", time.Now().Unix())
	})

	/** 定义处理器函数, 用于获取使用获取和设置cookie
	* 向客户端返回当前时间
	**/
	http.HandleFunc("/cookie/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method, request.URL)

		// 获取请求头中的Cookie信息
		fmt.Println(request.Header.Get("Cookie"))

		// 设置响应头中的Header信息
		writer.Header().Set("ServerName", "goserver")

		// 定义cookie
		token := http.Cookie{
			Name:     "token",
			Value:    "123456789",
			HttpOnly: true,
		}
		// 在响应头中增加增加cookie信息
		writer.Header().Add("Set-Cookie", token.String())
		http.SetCookie(writer, &token)

		// 响应体
		fmt.Fprintf(writer, "%d", time.Now().Unix())
	})

	// 重定向
	http.HandleFunc("/redirect/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "http://www.baidu.com", http.StatusMovedPermanently)
	})

	// 文件服务器
	http.Handle("/static/", http.FileServer(http.Dir(".")))

	// 启动http服务
	log.Fatal(http.ListenAndServe(":8888", nil))

	//启动https服务
	// log.Fatal(http.ListenAndServeTLS(":8888", "ssl/server.crt", "ssl/server.key", nil))
}
