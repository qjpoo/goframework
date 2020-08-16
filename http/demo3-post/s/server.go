package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {


	http.HandleFunc("/post", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()

		// 请求类型是application/x-www-form-urlencoded时解析form数据
		request.ParseForm()
		fmt.Println(request.PostForm)
		fmt.Println(request.PostForm.Get("name"), request.PostForm.Get("age"))
		fmt.Println(request.PostFormValue("name"))
		b, err := ioutil.ReadAll(request.Body)
		if err !=nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(b))  // {"name": "chiling", "age": 20}
		answer := `{"status": "ok"}`
		writer.Write([]byte(answer))

	})

	http.ListenAndServe("0.0.0.0:8080", nil)
}