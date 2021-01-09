package a18memento

import (
	"fmt"
)

func TestFunc() {
	fmt.Println("a18memento")

	var Gamer GamerStu
	Gamer.SetMoney(199)
	Gamer.GetFruit()
	Gamer.GetFruit()

	Memento := Gamer.CreatMemento()
	fmt.Println(Gamer)
	Gamer.GetFruit()
	Gamer.SetMoney(188)
	fmt.Println(Gamer)
	Gamer.ReatoreMemento(Memento)
	fmt.Println(Gamer)

}
