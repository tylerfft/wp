package a17observer

import (
	"fmt"
	"strconv"
)

type ObserverIntf interface {
	Update(GeneratorIntf)
}
type DigitObserverStu struct {
}

func (r *DigitObserverStu) Update(Gene GeneratorIntf) {
	fmt.Println("DigitObserverStu:" + strconv.Itoa(Gene.GetNum()))
}

type GarphObserverStu struct {
}

func (r *GarphObserverStu) Update(Gene GeneratorIntf) {
	fmt.Print("GarphObserverStu:")
	for i := 0; i < Gene.GetNum(); i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
