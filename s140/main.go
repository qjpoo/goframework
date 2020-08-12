package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	// mvc包的使用

	app := iris.New()

	// 设置自定义的控制器
	mvc.New(app).Handle(new(UserController))

	// mvc.Configure来配置路由组和控制器对象
	// localhost:8000/user/xxoo
	mvc.Configure(app.Party("/user", func(context *mvc.Application) {
		context.Handle(new(UserController))

	}))

	app.Run(iris.Addr(":8000"))

}

/*
url: http://localhost:8000
type: get
*/
func (uc *UserController) Get() string {
	iris.New().Logger().Info("Get 请求")
	return "hello, Get"
}

/*
url: http://localhost:8000
type: Post
*/
func (uc *UserController) Post() string {
	iris.New().Logger().Info("Post 请求")
	return "hello, Post"
}

/*
url: http://localhost:8000
type: Put
*/
func (uc *UserController) Put() string {
	iris.New().Logger().Info("Put 请求")
	return "hello, Put"
}

/*
url: http://localhost:8000/info
type: get
*/

// 请求方法 Get  uri Info
func (uc *UserController) GetInfo() mvc.Result {
	iris.New().Logger().Info("Get 请求路径为info") // [INFO] 2020/08/12 23:11 Get 请求路径为info
	return mvc.Response{
		Object: map[string]interface{}{
			"code":    1,
			"message": "请求成功",
		},
	}
}

/*
url: http://localhost:8000/login
type: post
*/
// 请求方法 Post  uri login
func (uc *UserController) PostLogin() mvc.Result {
	iris.New().Logger().Info("Post 请求路径为login") // [INFO] 2020/08/12 23:16 Post 请求路径为login
	return mvc.Response{
		Object: map[string]interface{}{
			"code":    1,
			"message": "请求成功",
		},
	}
}

type UserController struct {
}

func (uc *UserController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/query", "UserInfo")
}

// 自定义匹配别的请求
/*
url: http://localhost:8000/query
type: get
*/
func (uc *UserController) UserInfo() mvc.Result {
	iris.New().Logger().Info("user info query ...")
	return mvc.Response{}

}
