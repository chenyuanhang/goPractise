package main

import "fmt"

func main() {
	//自动判别类型
	var a  = "initial"
	fmt.Println(a)

	//定义数据类型
	var b, c int=1,2
	fmt.Println(b,c)

	var d =true
	fmt.Println(d)

	//默认赋值
	var e int
	fmt.Println(e)

	//声明并赋值的另一种方式

	f :="apple"
	fmt.Println(f)
}
