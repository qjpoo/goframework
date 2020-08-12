package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"os"
)

func main() {
	// 读取json格式的配置文件里面的值

	app := iris.Default()

	// 通过json配置文件进行应用配置
	file, _ := os.Open("./configs/config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err := decoder.Decode(&conf)  // 解码放入到conf结构体中
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(conf.Name, conf.Age)  // chiling 48


	app.Run(iris.Addr(":8000"))
}

// 定义一个跟config.json文件一样字段的struct
type Configuration struct {
	Name string `json:"name"`
	Age int `json:age`
}