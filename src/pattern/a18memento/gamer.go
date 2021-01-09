package a18memento

import (
	"math/rand"
	"time"
)

var FruitList []string

func init() {
	FruitList = make([]string, 4)
	FruitList[0] = "apple"
	FruitList[1] = "banana"
	FruitList[2] = "orange"
	FruitList[3] = "pear"
}

type GamerStu struct {
	Money  int
	Fruits []string
}

func (r *GamerStu) GetMoney() int {
	return r.Money
}
func (r *GamerStu) SetMoney(Money int) {
	r.Money = Money
}

func (r *GamerStu) GetFruits() []string {
	return r.Fruits
}

func (r *GamerStu) GetFruit() {
	rand.Seed(time.Now().UnixNano())
	r.Fruits = append(r.Fruits, FruitList[rand.Intn(4)])
}

func (r *GamerStu) CreatMemento() (Memento MementoStu) {
	Memento.Money = r.Money
	Memento.Fruits = r.Fruits
	return
}

func (r *GamerStu) ReatoreMemento(Memento MementoStu) {
	r.Money = Memento.Money
	r.Fruits = Memento.Fruits
	return
}
