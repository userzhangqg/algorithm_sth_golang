package main

import (
	"fmt"
	"time"
)

/*
已作答：0,0,1,1,2,10,

正确答案：0,1,0,2,1,2,10,

知识点：
channel的使用
*/

var receive = make(chan int, 1)
var wait = make(chan bool)

func main() {
	go func() {
		i := 0
		for i < 3 {
			receive <- i
			// 先发送a0，a1，被阻塞，5秒后同道中a1释放，a2塞入，切换到b执行。注意channel长度为1
			fmt.Printf("a%d,", i)
			i++
		}
	}()
	go func() {
		for {
			//v := <-receive
			//time.Sleep(5 * time.Second)
			//fmt.Printf("b%d,", v)
			//if v == 2 {
			//	wait <- true
			//}
			select {
			case v := <-receive:
				// 阻塞5秒，这时，a0吐出来，a1通道中
				time.Sleep(5 * time.Second)
				fmt.Printf("b%d,", v)
				if v == 2 {
					wait <- true
				}
			}
		}
	}()
	<-wait
	fmt.Printf("10,")
	time.Sleep(10 * time.Second)
}
