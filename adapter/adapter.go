package main

import "fmt"

type Banner struct {
	String string
}

func (this *Banner) showWithPattern() {
	fmt.Println("(" + this.String + ")")
}

func (this *Banner) showWithAster() {
	fmt.Println("*" + this.String + "*")
}

type Print interface {
	printWeak()
	printStrong()
}


type PrintBanner struct {
	banner *Banner
}

func (this *PrintBanner) printWeak() {
	this.banner.showWithPattern()
}

func (this *PrintBanner) printStrong() {
	this.banner.showWithAster()
}


func makePrintBanner(String string) *PrintBanner {
	return &PrintBanner{&Banner{String}}
}

func main(){
	var p Print = makePrintBanner("Hello")
	p.printWeak()
	p.printStrong()
}