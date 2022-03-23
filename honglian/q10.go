package main

import (
	"encoding/json"
	"fmt"
)

/*
已作答：编译错误

正确答案：,0

知识点：
https://sanyuesha.com/2018/05/07/go-json/
json序列化与反序列化，struct大小写访问权限问题

https://studygolang.com/articles/33878
我们调用json.Marshal()方法的时候，由于这个方法是json这个包的，如果我们结构体字段小写，json包的方法是无法使用本包结构体的首字母小写的字段。因为是私有的，不能跨包使用。
*/

/*type People10 struct {
	name string
	age int
}*/

type People10 struct {
	Name string
	Age  int
}

var data string = `{
						"name": "Jane",
						"age": 22,
						"Name": "Ella",
						"Age": 20
					}`

//var data string = `{"name": "Jane","age": 22}`

func main() {
	var people People10
	str := []byte(data)
	err := json.Unmarshal(str, &people)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%s,%d",people.name,people.age) // ,0
	fmt.Printf("%s,%d", people.Name, people.Age) // Ella,20
}
