package codeTop

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "1"

	v, ok := m[2]
	fmt.Println(v, ok)

}
