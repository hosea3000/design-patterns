package simple_factory

import "fmt"

// 抽象

type AbstractProduct interface {
	Show()
}

// 实现

type ProductA struct{}

func (p ProductA) Show() {
	fmt.Println("ProductA Show")
}

type ProductB struct{}

func (p ProductB) Show() {
	fmt.Println("ProductB Show")
}

// 逻辑

type Factory struct{}

func (f *Factory) CreateProduct(productType string) AbstractProduct {
	var product AbstractProduct
	switch productType {
	case "ProductA":
		product = &ProductA{}
	case "ProductB":
		product = &ProductB{}
	}

	return product
}

/*
* new(Factory).CreateProduct("ProductA")
 */
