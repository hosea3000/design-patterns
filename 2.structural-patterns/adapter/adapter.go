package main

import "fmt"

// 抽象

type V5 interface {
	Use5V()
}

type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v: v}
}

func (phone *Phone) Charge()  {
	fmt.Println("手机正在充电。。。")
	phone.v.Use5V()
}


type V220 struct {}

func (receiver *V220) Show()  {
	fmt.Println("正在输出220V电压")

}


// 没有5V的电压, 所以我们需要实现一个适配器

type V220ToV5Adapter struct {
	v *V220
}

func NewV220ToV5Adapter(v *V220) *V220ToV5Adapter {
	return &V220ToV5Adapter{v: v}
}

func (adapter *V220ToV5Adapter) Use5V() {
	fmt.Println("使用了适配器转换电压")
	adapter.v.Show()
}



func main() {
	phone := NewPhone(NewV220ToV5Adapter(new(V220)))
	phone.Charge()
}

// 1. 我有一个手机需要使用5V的电压来充电
// 2. 我现在只有22V电压的电源
// 3. 必须实现一个V220转5V的适配器
// 4. 使用适配器对接手机和22V电压完成充电

// 适配器模式：使用适配器将一个类型转换为另一个类型给别的对象使用

