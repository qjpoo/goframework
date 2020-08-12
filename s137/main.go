package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {

	app := iris.New()


	// type Get request
	// http://localhost:8080/hello
	// 拦截的请求  匿名函数
	// conext里面封闭了request, response请求
	// restful api
	app.Get("/hello", func(context context.Context) {
		// 处理get请求， 请求的Url为： /hello
		path := context.Path()
		// 日志输出
		app.Logger().Info(path)

		// 写入返回的数据： string
		//context.WriteString("request path: " + path)
		//context.Write([]byte("hello, world ..."))
		context.HTML("<h1>hello, world<h1>")
	})

	app.Get("/json", func(context context.Context) {
		context.JSON(&Person{"chiling", "123456"})

	})

	// curl -XPOST http://localhost:8000/
	app.Post("/", func(context context.Context) {
		context.WriteString("post request ...")

	})

	// 另外一种写法
	app.Handle("GET", "/user", func(context context.Context) {
		context.WriteString("/user")
		
	})

	// 参数是动态的
	// curl http://localhost:8000/user/10
	app.Handle("GET", "/user/{id:int}", func(context context.Context) {
		// 获取参数
		id := context.Params().Get("id")
		context.WriteString("user: " + id )

	})


	// curl http://localhost:8000/info?id=1
	app.Handle("GET", "/info", func(context context.Context) {
		// 获取参数
		//id := context.Params().Get("id")
		id  := context.URLParam("id")
		context.WriteString("info: " + id )

	})

	// curl http://localhost:8000/username/john
	app.Get("/username/{name}", func(c iris.Context) {
		name := c.Params().Get("name")
		c.Writef("hello, %s", name)
	})

	// http://localhost:8000/good/chiling/send
	app.Post("/good/{name:string}/{action:path}", func(c context.Context) {
		name := c.Params().Get("name")
		action  := c.Params().Get("action")
		msg := name + " is " + action
		c.WriteString(msg)
	})

	// form-data   x-www-form-urlencoded
	app.Post("/post", func(c context.Context) {
		c.WriteString("post request: " + c.FormValue("name"))
	})


	// raw  json




	// 端口监听
	//app.Run(iris.Addr("0.0.0.0:8080"))
	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))


}

type Person struct {
	Name string `json:"username"`
	Password string `json:"passwd"`
}
