package a17observer

import (
	"math/rand"
	"time"
)

type GeneratorIntf interface {
	GetNum() int
}

type Numgenerator struct {
	Num       int
	Observers []ObserverIntf
}

func (r *Numgenerator) Register(Observer ObserverIntf) {
	r.Observers = append(r.Observers, Observer)
}
func (r *Numgenerator) Notify() {
	for _, v := range r.Observers {
		v.Update(r)
	}
}
func (r *Numgenerator) GetNum() int {
	return r.Num
}
func (r *Numgenerator) Exec() {
	for i := 0; i < 20; i++ {
		rand.Seed(time.Now().UnixNano())
		r.Num = rand.Intn(50)
		r.Notify()
		time.Sleep(time.Millisecond * 500)
	}
}
