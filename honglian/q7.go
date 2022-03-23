package main

import "fmt"

/*
已作答：编译错误

正确答案：12

知识点：
Go里面 switch 默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。
fallthrough不能用在switch的最后一个分支
fallthrough到下一个case块时，不执行case匹配检查！不执行case匹配检查！不执行case匹配检查！
*/

func main() {
	var object interface{}
	object = "string"
	//object = true
	object = 1
	switch object {
	case 1:
		fmt.Printf("1")
		fallthrough
	case "string":
		fmt.Printf("2")
	//	fallthrough
	//case "bool":
	//	fmt.Printf("bool")
	case 2:
		fmt.Printf("3")
		break
	case 3:
		fmt.Printf("4")
	default:
		fmt.Printf("5 ")
	}
}
