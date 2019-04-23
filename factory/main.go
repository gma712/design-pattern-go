package main

import (
	"github.com/work/design-pattern/factory/framework"
	"github.com/work/design-pattern/factory/idcard"
)



func main(){
	var factory framework.Factory = new(idcard.IDCardFactory)
	var card1 framework.Product = factory.Create("hogehoge")
	var card2 framework.Product =factory.Create("fugafuga")
	card1.Use()
	card2.Use()
}