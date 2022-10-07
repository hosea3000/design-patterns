package main

import "fmt"

// 抽象

type Phone interface {
	Show()
}

//
// type Decorator struct {
// 	phone Phone
// }
//
// func (d Decorator) Show() {}

// 实现

type HuaweiPhone struct {}

func (h HuaweiPhone) Show() {
	fmt.Println("秀出了华为手机")
}

type XiaomiPhone struct {}

func (h XiaomiPhone) Show() {
	fmt.Println("秀出了小米手机")
}


type KeDecorator struct {
	phone Phone
}

func (ke KeDecorator) Show() {
	fmt.Println("对手机进行了套壳")
	ke.phone.Show()
}

func NewKeDecorator(phone Phone) Phone {
	return KeDecorator{
		phone: phone,
	}
}


type MoDecorator struct {
	phone Phone
}

func NewMoDecorator(phone Phone) Phone {
	return MoDecorator{
		phone: phone,
	}
}

func (mo MoDecorator) Show() {
	fmt.Println("对手机进行了贴膜")
	mo.phone.Show()
}

// 业务

func main()  {
	var huaweiPhone Phone
	huaweiPhone = new(HuaweiPhone)
	huaweiPhone.Show()

	fmt.Println("----------------")
	var keMoPhone Phone
	keMoPhone = NewKeDecorator(NewMoDecorator(huaweiPhone))
	keMoPhone.Show()
}

// 装饰器模式：是通过新装饰对象访问原来的对象。新装饰对象能够提供更多的功能。和代理模式不同的是对象可以装饰多次，且顺序不定