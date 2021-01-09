package a21proxy

import (
	"fmt"
	"time"
)

type PrinterStu struct {
	Name string
}

func (r *PrinterStu) Init(name string) {
	r.Name = name
	r.HeaveJob("start constructor heavry job ...")
}

func (r *PrinterStu) SetName(name string) {
	r.Name = name
}

func (r *PrinterStu) GetName() (name string) {
	name = r.Name
	return
}

func (r *PrinterStu) Print(str string) {
	fmt.Println("=== " + r.Name + " ===")
	fmt.Println(str)
}

func (r *PrinterStu) HeaveJob(msg string) {

	fmt.Println(msg)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		fmt.Print(".")
	}
	fmt.Println("end")
}
