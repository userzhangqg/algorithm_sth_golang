package main

import "fmt"

/*
已作答：编译报错

正确答案：,0

知识点，引用传递，值传递
*/

type People struct {
	age  *int
	name string
}

func NewPeople(name string, age int) (p *People) {
	p = new(People)
	p.age = new(int)
	p.SetName(name)
	p.SetAge(age)
	return
}

// SetAge 值传递，非指针传递
func (p People) SetAge(age int) {
	p.age = &age
}

func (p People) GetAge() int {
	return *p.age
}

func (p People) SetName(name string) {
	p.name = name
}

func (p People) GetName() string {
	return p.name
}

func main() {
	var people *People = NewPeople("John", 22)
	people.SetName("Grace")
	people.SetAge(45)
	fmt.Printf("%s,%d", people.GetName(), people.GetAge())
}
