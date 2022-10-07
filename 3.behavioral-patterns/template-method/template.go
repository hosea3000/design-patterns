package main

type Steps interface {
	Step1()
	Step2()
	Step3()
	Step4()
}


type template struct {
	s Steps
}

func (t *template) templateStep()  {
	if t.s == nil {
		return
	}

	t.s.Step1()
	t.s.Step2()
	t.s.Step3()
	t.s.Step4()
}

type ProductA struct {
	template
}

func NewProductA() *ProductA {
	productA := &ProductA{
		template: template{},
	}

	productA.template.s = productA
	return productA
}

func (p ProductA) Step1() {
	println("ProductA Step1")
}

func (p ProductA) Step2() {
	println("ProductA Step2")
}

func (p ProductA) Step3() {
	println("ProductA Step3")
}

func (p ProductA) Step4() {
	println("ProductA Step4")
}

func main() {
	productA := NewProductA()
	productA.templateStep()


	// t := template{
	// 	s: &ProductA{},
	// }
	//
	// t.templateStep()
}


// 模板方法模式：父类定义好流程。 子类只需要实现具体的步骤






