package framework

type Product interface{
	Use()
}

type Factory interface{
	CreateProduct(owner string) Product
	RegisterProduct(product Product)
	Create(owner string) Product
}
