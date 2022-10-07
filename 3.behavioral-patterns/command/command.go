package main

import "fmt"

// 抽象层

type AbstractCooker interface {
	RoastMutton()
	RoastBeef()
}

type AbstractCommand interface {
	StartCook()
}



// 实现层

type Cooker struct {}

func (c Cooker) RoastMutton()  {
	fmt.Println("厨师烤羊肉")
}


func (c Cooker) RoastBeef()  {
	fmt.Println("厨师烤牛肉")
}

type Waiter struct {
	commandList []AbstractCommand
}

func (w *Waiter) PlaceOrder()  {
	for _, command := range w.commandList {
		command.StartCook()
	}
}

type MuttonCommand struct {
	c Cooker
}

func (mutton *MuttonCommand) StartCook()  {
	mutton.c.RoastMutton()
}

type BeefCommand struct {
	c Cooker
}

func (mutton *BeefCommand) StartCook()  {
	mutton.c.RoastBeef()
}



func main() {
	// 首先饭店得有一个厨师
	cooker := Cooker{}
	waiter := &Waiter{}

	// 开始点单
	muttonCommand := &MuttonCommand{cooker}
	beefCommand := &BeefCommand{cooker}

	// 把订单交给服务员
	waiter.commandList = append(waiter.commandList, muttonCommand, beefCommand)

	// 服务员下单。下单之后厨师就会开始做菜
	waiter.PlaceOrder()

}

// 命令模式： 调用者和接受者解耦。通过命令交互
// 使用场景： 1. 调用者和接受者解耦； 2. 需要把多个命令组合使用
