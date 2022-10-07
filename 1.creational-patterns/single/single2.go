package single

import "sync"

type single2 struct {}

var single2Instance *single2

func NewSingle2() *single2 {
	new(sync.Once).Do(func() {
		single2Instance = &single2{}
	})

	return single2Instance
}


