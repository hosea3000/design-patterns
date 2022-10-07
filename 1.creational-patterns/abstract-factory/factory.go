package abstract_factory

import "fmt"

// 抽象

type AbstractProductA interface {
	Show()
}
type AbstractProductB interface {
	Show()
}


type AbstractFactoryA interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

type AbstractFactoryB interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

// 实现

type FactoryAProductA struct {}

func (p FactoryAProductA) Show() {
	fmt.Println("FactoryAProductA Show")
}

type FactoryAProductB struct {}

func (p FactoryAProductB) Show() {
	fmt.Println("FactoryAProductB Show")
}

type FactoryBProductA struct {}

func (p FactoryBProductA) Show() {
	fmt.Println("FactoryBProductA Show")
}

type FactoryBProductB struct {}

func (p FactoryBProductB) Show() {
	fmt.Println("FactoryBProductB Show")
}

type FactoryA struct {}

func (f FactoryA) CreateProductA() AbstractProductA {
	return &FactoryAProductA{}
}

func (f FactoryA) CreateProductB() AbstractProductB {
	return &FactoryAProductB{}
}

type FactoryB struct {}

func (f FactoryB) CreateProductA() AbstractProductA {
	return &FactoryBProductA{}
}

func (f FactoryB) CreateProductB() AbstractProductB {
	return &FactoryBProductB{}
}

// 业务

/*
	factoryAProductA := new(FactoryA).CreateProductA()
	factoryAProductB := new(FactoryA).CreateProductB()

    factoryBProductA := new(FactoryA).CreateProductA()
	factoryBProductB := new(FactoryA).CreateProductB()
 */