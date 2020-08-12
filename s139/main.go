package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	// 路由组
	/*
		用户模块users
		xxx/users/register  注册
		xxx/users/login   登陆
		xxx/user/info   获取用户信息
	*/

	// localhost:8000/users/{register, login, info}
	// 路由组的请求
	app := iris.New()
	userParty := app.Party("/users", func(context context.Context) {
		// 处理下一级的请求(register, login, info)
		context.Next()
	})

	// 路由组下面的下一级请求
	// register
	userParty.Get("/register", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
		//c.WriteString("path:" + path)  // path:/users/register
		c.WriteString("user register ...")
	})

	// login
	userParty.Get("/login", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
		//c.WriteString("path:" + path)  // path:/users/register
		c.WriteString("user login ...")
	})

	// info
	userParty.Get("/info", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
		//c.WriteString("path:" + path)  // path:/users/register
		c.WriteString("info login ...")
	})



	// 另外一种写法
	//  定义了一个路由组 localhost:8000/admin
	usersRouter := app.Party("/admin", userMiddleware)

	// Done 这个要写在userRouter.Get之上面
	usersRouter.Done(func(c context.Context) {
		c.Application().Logger().Info("response send to " + c.Path())  // [INFO] 2020/08/12 22:13 response send to /admin/info
	})

	// /admin/ino
	usersRouter.Get("/info", func(c context.Context) {
		c.HTML("<h1>/admin/info</h1>")
		c.Next()  // 手动的显示调用， 这个会执行userRouter.Done()
	})


	// query
	// admin/query 没有c.Next()
	usersRouter.Get("/query", func(c context.Context) {
		c.HTML("<h1>/admin/query</h1>")
	})

	app.Run(iris.Addr(":8000"))
}

func userMiddleware(context iris.Context) {
	context.Next()
}
