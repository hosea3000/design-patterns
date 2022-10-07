package single

import "sync"

var lock = sync.Mutex{}

type single struct {}

var singles *single

func NewSingle() *single {
	lock.Lock()
	defer lock.Unlock()

	if singles == nil {
		singles = &single{}
	}
	return singles
}






