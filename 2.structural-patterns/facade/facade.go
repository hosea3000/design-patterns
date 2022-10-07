package main

import "fmt"

type System1 struct {}

func (s *System1) Show()  {
	fmt.Println("System1")
}

type System2 struct {}

func (s *System2) Show()  {
	fmt.Println("System2")
}

type System3 struct {}

func (s *System3) Show()  {
	fmt.Println("System3")
}

type System4 struct {}

func (s *System4) Show()  {
	fmt.Println("System4")
}

type Facade struct {
	system1 *System1
	system2 *System2
	system3 *System3
	system4 *System4
}

func (f *Facade) Show() {
	f.system1.Show()
	f.system2.Show()
	f.system3.Show()
	f.system4.Show()
}

func main() {
	new(Facade).Show()
}

// 外观模式： 将子系统组合起来统一对外表达