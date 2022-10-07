package factory_method

import "fmt"

// 抽象

type AbstractProduct interface {
	Show()
}

type AbstractFactory interface {
	CreateProduct() *AbstractProduct
}

// 实现

type ProductA struct {}

func (p ProductA) Show() {
	fmt.Println("ProductA Show")
}

type ProductB struct {}

func (p ProductB) Show() {
	fmt.Println("ProductB Show")
}

type FactoryA struct {}

func (f FactoryA) CreateProduct() AbstractProduct {
	return  &ProductA{}
}

type FactoryB struct {}

func (f FactoryB) CreateProduct() AbstractProduct {
	return  &ProductB{}
}

// 业务

/*
	productA := new(FactoryA).CreateProduct()
	productA := new(FactoryB).CreateProduct()
 */


