package main
import (
"time"
)

// 定义time.Duration的别名为MyDuration
type MyDuration1  =time.Duration


// 定义 MyDuration 为time.Duration 类型
type MyDuration2 time.Duration
// 为MyDuration添加一个函数
func (m MyDuration2) EasySet(a string) {
}
func main() {

}
