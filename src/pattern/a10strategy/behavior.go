package a10strategy

import (
	"fmt"
)

type FlyIntf interface {
	Fly()
}

type FlyWithWingsStu struct {
}

func (r *FlyWithWingsStu) Fly() {
	fmt.Println("FlyWithWingsStu")
}

type FlyNoWayStu struct {
}

func (r *FlyNoWayStu) Fly() {
	fmt.Println("FlyNoWayStu")
}

type FlyWithRocketStu struct {
}

func (r *FlyWithRocketStu) Fly() {
	fmt.Println("FlyWithRocketStu")
}
