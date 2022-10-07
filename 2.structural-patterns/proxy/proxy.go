package main

import "fmt"

// 抽象

type AbstractShopper interface {
	Shopping()
}

// 实现

type KrShopping struct {}

func (k KrShopping) Shopping() {
	fmt.Println("韩国购物")
}


type JaShopping struct {}

func (k JaShopping) Shopping() {
	fmt.Println("日本购物")
}


type AmShopping struct {}

func (k AmShopping) Shopping() {
	fmt.Println("美国购物")
}


type ShoppingProxy struct {
	shopper AbstractShopper
}

func NewShoppingProxy(shopper AbstractShopper) *ShoppingProxy {
	return &ShoppingProxy{shopper: shopper}
}

func (s ShoppingProxy) Shopping() {
	if s.Check() {
		s.shopper.Shopping()
	}
}

func (s ShoppingProxy) Check() bool{
	fmt.Println("鉴别真伪")
	return true
}

// 业务

func main() {
	shopper := &KrShopping{}
	proxy := NewShoppingProxy(shopper)
	proxy.Shopping()

	amShopper := &AmShopping{}
	amProxy := NewShoppingProxy(amShopper)
	amProxy.Shopping()
}

// 代理模式：是通过一个新的代理对象访问原来的对象。新的代理对象能够提供更多的功能。