package main

import (
	"fmt"
	"github.com/work/design-pattern/singleton/singleton"
)


func main() {
	fmt.Println("Start.")
	obj1 := singleton.GetInstance()
	obj2 := singleton.GetInstance()
	fmt.Println(obj1)
	fmt.Println(obj2)
	if obj1 == obj2 {
		fmt.Println("同じインスタンスです。")
	} else {
		fmt.Println("異なるインスタンスです。")
	}
}