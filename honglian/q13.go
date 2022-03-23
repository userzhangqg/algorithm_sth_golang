package main

import "fmt"

/*
13. 请写一段简短的代码实现C语言中的char*转换成Golang中的[]byte？

已作答：
func getChar(s String) {
    return byte(s)
}

正确答案：
func getChar(s String) {
    return []byte(s)
}

*/

func getChar(s string) []byte {
	return []byte(s)
}

func main() {
	fmt.Println(getChar("0aa"))
}
