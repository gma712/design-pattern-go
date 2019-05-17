package singleton

import "fmt"


type Singleton struct {}

var singleton *Singleton

func init() {
	fmt.Println("Initialized.")
	singleton = &Singleton{}
}

func (s *Singleton) String() string {
	return "インスタンスが生成されました。"
}

func GetInstance() (*Singleton) {
	return singleton
}
