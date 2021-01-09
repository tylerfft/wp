package a05singleton

import (
	"fmt"
	"sync"
)

var SingleP *SingleStu
var once sync.Once

func GetInstance() (out *SingleStu) {
	once.Do(func() {
		SingleP = &SingleStu{}
	})
	return SingleP
}

type SingleStu struct {
	sync.RWMutex
}

func (r *SingleStu) Print() {
	r.Lock()
	defer r.Unlock()

	fmt.Println("I am a singleton")
}
