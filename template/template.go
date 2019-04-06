package main

import "fmt"

type AbstructDisplay interface{
	open()
	print()
	close()
}

func display(this AbstructDisplay) {
	this.open()
	for i := 0; i < 5; i++ {
		this.print()
	}
	this.close()
}

type CharDisplay struct {
	ch string
}


func (this *CharDisplay) open() {
	fmt.Println("<<")
}

func (this *CharDisplay) print() {
	fmt.Println(this.ch)
}

func (this *CharDisplay) close() {
	fmt.Println(">>")
}

type StringDisplay struct {
	str string
	width int
}

func (this *StringDisplay) open() {
	this.printLine()
}

func (this *StringDisplay) print() {
	fmt.Println("|" + this.str + "|")
}

func (this *StringDisplay) close() {
	this.printLine()
}

func (this *StringDisplay) printLine() {
	bar := ""
	for i := 0; i < this.width; i++ {
		bar += "-"
	}
	fmt.Println("+" + bar + "+")
}

func setStringDisplay(str string) *StringDisplay {
	w := len(str)
	return &StringDisplay{
		str: str,
		width: w,
	}
}


func main() {
	var d1 AbstructDisplay = &CharDisplay{"a"}
	var d2 AbstructDisplay = setStringDisplay("Hello, world.")

	display(d1)
	display(d2)
}