package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	// 客户端发送请求
	resp, err := http.Get("http://127.0.0.1:8000/query?name=chiling&age=48")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 把从服务端的数据resp读出来
	b, err := ioutil.ReadAll(resp.Body) // 在客户端打印从服务端发来的响应体

	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(b))

	// 带有参数的get请求
	apiUrl := "http://127.0.0.1:8000/query"
	data := url.Values{}
	data.Set("name", "chiling")
	data.Set("age", "20")
	u, err := url.ParseRequestURI(apiUrl) // u  *URL
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(data, data.Encode()) // map[age:[20] name:[chiling]] age=20&name=chiling
	//fmt.Println(u)  // http://127.0.0.1:8080/query

	u.RawQuery = data.Encode()
	//fmt.Println(u.String())  // http://127.0.0.1:8080/query?age=20&name=chiling

	resp1, err := http.Get(u.String())
	defer resp1.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	b1, err := ioutil.ReadAll(resp.Body)
	if err !=nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("xxxoo",string(b1))






}
