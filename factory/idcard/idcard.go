package idcard

import (
	"fmt"
	"github.com/work/design-pattern/factory/framework"
)

type IDCard struct {
	Owner string
}

func (idCard *IDCard) Use() {
	fmt.Printf("%s のカードを使います。\n", idCard.Owner)
}

func (idCard * IDCard) GetOwner() string {
	return idCard.Owner
}

type IDCardFactory struct {
	Owners []framework.Product
}

func (factory *IDCardFactory) CreateProduct(owner string) (framework.Product) {
	fmt.Printf("%s のカードを作ります。\n", owner)
	return &IDCard{
		Owner: owner,
	}
}

func (factory *IDCardFactory) RegisterProduct(product framework.Product) {
	factory.Owners = append(factory.Owners, product)
}

func (factory *IDCardFactory) GetOwners() []framework.Product {
	return factory.Owners
}

func (factory *IDCardFactory) Create(owner string) (framework.Product){
	var p framework.Product = factory.CreateProduct(owner)
	factory.RegisterProduct(p)
	return p
}