package main

import (
	"fmt"
	"math"
)

const  (
	//定义每分钟的秒数
	SecondsPerMinute = 60

	//定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60

	// 定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

var imp = [3]int{1,2,3}

//函数 func 函数名(形式参数列表)(返回值列表){函数体}
func hypot(x,y float64)  float64 {
	return  math.Sqrt(x*x+y*y)

}



func Multiply(a [3]int) int  {
	return a[0]*a[1]*a[2]
}

func resolve(seconds int)(day int,hour int, minute int){
	day=seconds/ SecondsPerDay
	hour=(seconds-day*SecondsPerDay) /SecondsPerHour
	minute=(seconds-day*SecondsPerDay-hour*SecondsPerHour)/SecondsPerMinute
	return
}

func main() {
	fmt.Println(hypot(3,4))
	fmt.Println(Multiply(imp))
	day, hour, minute := resolve(10000);
	fmt.Printf("day= [%d],hour=[%d],minute=[%d]",day,hour,minute)

}
