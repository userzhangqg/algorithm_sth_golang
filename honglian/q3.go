package main

import "fmt"

/*
已作答：编译报错

正确答案：100

知识点：interface接口实现
*/

type Callback interface {
	OnProgress(progress int)
}

type Download struct {
}

func (d *Download) OnProgress(progress int) {
	fmt.Println(progress)
}

func GetDownload(role string) Callback {
	var d *Download
	if role == "admin" {
		d = new(Download)
	}
	return d
}

func main() {
	c := GetDownload("visitor")
	b := GetDownload("admin")
	fmt.Println(c, b)
	// 使用 GetDownload() 的返回值与 nil 判断时，虽然接口里的 value 为 nil，但 type 带有 *Callback 信息，使用 == 判断相等时，依然不为 nil。
	if c != nil {
		c.OnProgress(100)
	} else {
		c.OnProgress(0)
	}
}
