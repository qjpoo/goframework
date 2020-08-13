package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/sessions/sessiondb/boltdb"
)

var (
	USERNAME = "username"
	ISLOGIN = "islogin"
)
func main() {
	// session 使用

	app := iris.New()

	sessionID := "mySession"

	// 创建session并进行使用
	sess := sessions.New(sessions.Config{
		Cookie: sessionID,
	})

	// 用户登陆
	app.Post("/login", func(context context.Context) {
		path := context.Path()
		app.Logger().Info("请求Path: ",path)
		userName := context.PostValue("name")
		passwd := context.PostValue("pwd")

		if userName == "chiling" && passwd == "123" {
			session := sess.Start(context)

			// 用户名
			session.Set(USERNAME, userName)
			// 登陆状态
			session.Set(ISLOGIN, true)
			//all := session.GetAll()
			all := session.GetString(ISLOGIN)
			app.Logger().Info("all: ",all)

			context.WriteString("帐户登陆成功....")
		} else {
			session := sess.Start(context)
			session.Set(ISLOGIN, false)
			context.WriteString("帐户登陆失败,请重新登陆 ...")
		}
	})

	app.Get("/logout", func(c context.Context) {
		path := c.Path()
		app.Logger().Info("退出登陆 path: ", path)
		session := sess.Start(c)
		// 删除session
		session.Delete(ISLOGIN)
		session.Delete(USERNAME)
		c.WriteString("退出登陆 ...")
	})

	app.Get("query", func(c context.Context) {
		path := c.Path()
		app.Logger().Info("查询登陆 path: ", path)

		session := sess.Start(c)
		islogin, err := session.GetBoolean(ISLOGIN)
		// session为空的，就会报错
		if err !=nil {
			c.WriteString("用户没有登陆, 请先登陆 ...")
			return
		}

		if islogin {
			app.Logger().Info("islogin: ", islogin)
			c.WriteString("用户已登陆")
		}else {
			app.Logger().Info("islogin: ", islogin)
			c.WriteString("用户没有登陆,请先登陆")

		}
	})


	// session和DB绑定
	db, err := boltdb.New("./s141/session.db", 0600)
	if err != nil {
		panic(err.Error())
	}

	// 程序中断,将数据库关系
	iris.RegisterOnInterrupt(func() {
		defer db.Close()
	})

	// sessiont db绑定
	sess.UseDatabase(db)

	app.Run(iris.Addr(":8000"))


	
}
