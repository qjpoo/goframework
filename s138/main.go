package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	// 路由功能处理方式

	app := iris.New()

	//统一处理get方法
	// localhost:8000/userInfo
	app.Handle("GET", "/userInfo", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)

		context.HTML("<h1>hello</h1>")
	})

	// date是一个变量
	app.Get("/weatherInfo/{date}", func(c context.Context) {
		// 正则表达式变量内容的值
		date := c.Params().Get("date")
		c.WriteString(date) // 2019
	})
/*
	// {userId:int}对变量的类型限制
	app.Get("/api/users/{userId:int}", func(c context.Context) {
		userId, err := c.Params().GetInt("userId")
		app.Logger().Info(err)
		if err != nil {
			// 设置请求的状态码， 状态码可以自定义
			c.JSON(map[string]interface{}{
				"requestCode": 500,
				"message":     "bad request",
			})
			return
		}
		app.Logger().Info(userId)
		//c.Writef("userId: %d", userId) // userId: 20
		c.JSON(map[string]interface{}{
			"requestCode": 200,
			"message":     "ok",
			"userId":      userId,
		})
	})
*/

	// bool
	app.Get("/api/user/{isLogin:bool}", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
		isLogin, err := c.Params().GetBool("isLogin")
		if err != nil {
			c.StatusCode(iris.StatusNonAuthoritativeInfo)
			return
		}
		if isLogin {
			c.WriteString("Login sucessfull ...\n")
		} else {
			c.WriteString("Login failure ...\n")
		}

		c.JSON(map[string]interface{}{
			"isLogin": isLogin,
			"code":    "200",
		})
	})

	app.Run(iris.Addr(":8000"))

}
