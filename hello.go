package main


import "fmt"

func main() {
	fmt.Println("hello, world ..")
	var xxoo string = ""
	fmt.Println(xxoo)
	xxoo = "true"
	fmt.Println(xxoo)

	n := true
	fmt.Println(n)

	fmt.Println("---------------------")
	v1 := "hello"
	v2 := []byte("abc")  // 字符串转化为byte切片
	fmt.Println(string(v2))  // 切片转化为字符串
	fmt.Println([]byte(v1))
	fmt.Println(string(v2[:]))

	fmt.Println("---------------------")

}
