package main

import "fmt"

// 抽象层

type AbstractListener interface {
	// AttachListener 监听者

	AttachListener(...AbstractStudent)
	DetachListener(AbstractStudent)
	Notify()
}


type AbstractStudent interface {
	// AbstractStudent 观察者

	OnTeacherComeIn()
	DoBadThings()
}


// 实现层

type Monitor struct {
	listener []AbstractStudent
}

var _ AbstractListener = (*Monitor)(nil)

func (m *Monitor) Notify() {
	for _, stu := range m.listener {
		stu.OnTeacherComeIn()
	}
}

func (m *Monitor) AttachListener(stu ...AbstractStudent) {
	m.listener = append(m.listener, stu...)
}

func (m *Monitor) DetachListener(stu AbstractStudent) {
	for i := 0; i < len(m.listener); i++ {
		if m.listener[i] == stu {
			m.listener = append(m.listener[:i], m.listener[i+1:]...)
			return
		}
	}
}

// 业务

type Student struct {
	name  string
	things string
}

var _ AbstractStudent = (*Student)(nil)

func (s *Student) OnTeacherComeIn() {
	fmt.Printf("%s 停止 %s\n", s.name, s.things)
}

func (s *Student) DoBadThings() {
	fmt.Printf("%s 正在 %s\n", s.name, s.things)
}

func main() {
	stu1 := &Student{"张三", "打游戏"}
	stu2 := &Student{"李四", "抄作业"}
	stu3 := &Student{"王五", "听歌"}


	m := new(Monitor)
	m.AttachListener(stu1, stu2, stu3)

	fmt.Println("刚开始的时候。。。")
	stu1.DoBadThings()
	stu2.DoBadThings()
	stu3.DoBadThings()

	fmt.Println("当班长通知说老师来了。。。")
	m.Notify()
}

// 观察者模式: 一种订阅和通知的机制。 可以通知所有订阅者，订阅者收到消息可以做不同的动作
// 注意：不能在订阅者收到消息后发通知，这样很有可能造成死循环
