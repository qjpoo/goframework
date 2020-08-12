package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	// request

	app := iris.New()

	// get请求
	app.Get("/getrequest", func(context context.Context) {
		// 得到路径
		path := context.Path()
		// 打印日志
		app.Logger().Info(path)
		// 写出返回的数据 以字符串
		context.WriteString("request path: " + path)
	})

	// 处理get请求，接受参数
	// localhost:8000/userinfo?username=qujian&passwd=123456
	// c.URLParam
	app.Get("/userinfo", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)

		username := c.URLParam("username")
		app.Logger().Info(username)

		passwd := c.URLParam("passwd")
		app.Logger().Info(passwd)

		// html方式返回
		c.HTML("<h1>" + "username: " + username + "passwd: " + passwd + "</h1>")

	})

	fmt.Println("------------------post---------------------")
	// PostValue获取post里面提交的参数
	app.Post("/postlogin", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)

		name := c.PostValue("name")
		pwd := c.PostValue("pwd")
		app.Logger().Info(name, pwd)
		c.HTML("<h1>" + name + pwd + "</h1>")

	})



	// json

	/*
	postman: json格式传入
	{
		"name": "chiling",
		"password": "1223"
	}

	 */
	app.Post("/postjson", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)

		// json数据解析  ReadJSON
		//app.Logger().Info(c.PostValue("name"))
		//var person Person = Person{"chiling","123"}
		var person Person
		// 读入前端传入的json格式的数据
		if err := c.ReadJSON(&person); err != nil {
			panic(err.Error())
		}
		app.Logger().Info(person)
		c.Writef("Received: %#+v\n", person)  // Received: main.Person{Name:"chiling", Pwd:"1223"}
		c.Writef("Received: %v\n", person)  //  Received: {chiling 1223}
	})


	// xml readXml
	//postman 传入xml格式
	/*
	<student>
		<stu_name>qujian</stu_name>
		<stu_age>20</stu_age>
	</student>

	 */
	app.Post("/postXml", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)

		// xml解析
		var student Student
		if err := c.ReadXML(&student);err !=nil {
			panic(err.Error())
		}

		// 输出
		app.Logger().Info(student)
		c.Writef("Received: %#+v\n", student)  // Received: main.Student{Name:"qujian", Age:20}
	})


	// 返回json格式
	app.Get("/getJson", func(c context.Context) {
		c.JSON(iris.Map{"msg: ": "hello, world ...", "response code: ": 200})
		/*
		{
		    "msg: ": "hello, world ...",
		    "response code: ": 200
		}
		 */
	})

	// 返回text
	app.Get("getText", func(c context.Context) {
		c.Text("text  ...")
	})

	// put请求
	app.Put("/putHello", func(c context.Context) {
		c.WriteString("put method ...")
	})

	// delete
	app.Delete("/deleteHello", func(c context.Context) {
		c.WriteString("delete method ...")
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}

type Person struct {
	Name string `json:"name"`
	Pwd string `json:"password"`
}

type Student struct {
	Name string `xml:"stu_name"`
	Age int `xml:"stu_age"`
}
