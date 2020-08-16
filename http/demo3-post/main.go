package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// post

	url := "http://127.0.0.1:8080/post"
	contentType := "application/json"
	data := `{"name": "chiling", "age": 20}`
	resp, err := http.Post(url, contentType,strings.NewReader(data)) // strings.NewReader 把字符串转为 Reader对象
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(b))

}
