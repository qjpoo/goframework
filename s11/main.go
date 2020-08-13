package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 不要忘记了
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func main() {
	// xorm

	// 1. 创建一个数据库的引擎对象
	engine, err := xorm.NewEngine("mysql", "root:qujian123@tcp(47.98.179.41:13360)/mysql?charset=utf8mb4")
	if err != nil {
		panic(err.Error())
	}

	// 2. 数据引擎关闭]
	defer engine.Close()

	// 数据库引擎设置
	engine.ShowSQL(true)                     // 显示sql语句
	engine.Logger().SetLevel(core.LOG_DEBUG) // 日志等级
	engine.SetMaxOpenConns(10)               // 设置最大的连接数
	engine.SetMaxIdleConns(2)

	// 查询表的所有数据
	session := engine.Table("user")
	count, err := session.Count()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("count: ", count)

	// 使用原生的sql语句来查询
	result, err := engine.Query("select host, user from user")
	if err != nil {
		panic(err.Error())
	}
	for _, value := range result {
		//fmt.Println(key, "===", value)
		for k, v := range value {
			fmt.Println(k, "-----", string(v[:]))
		}
	}

	// 设置自动同步结构体到数据库
	//err = engine.Sync2(new(Person))
	//fmt.Println(err)


}

type Person struct {
	Name string
	age  int
}
