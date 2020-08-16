package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request)  {
	//str := `<h1 style="color:blue">hello</h1>`
	b, err := ioutil.ReadFile("./http/demo1-server/1.txt")
	if err !=nil {
		//fmt.Println(err.Error())
		//fmt.Println(fmt.Sprintf("%v", err))
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	//w.Write([]byte(str))
	w.Write(b)
}



func main() {

	//设置访问的路由
	http.HandleFunc("/hello", f1)


	h1 := func(w http.ResponseWriter, _ *http.Request) {

		io.WriteString(w, "i'am a h1 ..\n")
	}
	http.HandleFunc("/h1", h1)

	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "i'am a h2 ...\n")
	}
	http.HandleFunc("/h2", h2)

	h := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "i'am a h ...\n")
	}
	http.HandleFunc("/", h)


	// query
	http.HandleFunc("/query", func(writer http.ResponseWriter, request *http.Request) {
		method := request.Method
		host := request.Host
		// 对于get请求，参数都是放在url里面的比如 query param   localhsot:8000/query?name=xx&name=18,请求体是没有数据的
		url := request.RequestURI
		parse := request.URL.Query()
		nameValue := parse.Get("name")
		ageValue:= parse.Get("age")
		userAgent := request.UserAgent()
		body, _ := ioutil.ReadAll(request.Body)
		// 打印客户端发来的请求信息
		fmt.Printf("method: %s, host: %s, url: %s, userAgent: %s, body: %s, parse[name]: %s, parse[age]:%s\n",method, host, url, userAgent, body, nameValue, ageValue)
		writer.WriteHeader(200)
		writer.Write([]byte("request sucesfull ..."))
		writer.Write(body)

	})
	// 开启服务
	http.ListenAndServe("0.0.0.0:8000", nil)
}
