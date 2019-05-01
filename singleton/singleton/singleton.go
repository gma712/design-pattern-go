package singleton


type Singleton struct {}

var singleton *Singleton

func (s *Singleton) String() string {
	return "インスタンスが生成されました。"
}

func GetInstance() (*Singleton) {
	singleton = &Singleton{}
	return singleton
}
