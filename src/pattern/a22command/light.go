package a22command

import (
	"fmt"
)

// LightStu
type LightStu struct {
}

func (r *LightStu) On() {
	fmt.Println("loght on")
}
func (r *LightStu) Off() {
	fmt.Println("loght Off")
}

// LightOnCommandStu
type LightOnCommandStu struct {
	Light *LightStu
}

func (r *LightOnCommandStu) Init(Light *LightStu) {
	r.Light = Light
}

func (r *LightOnCommandStu) Execute() {
	r.Light.On()
}

func (r *LightOnCommandStu) Undo() {
	r.Light.Off()
}

// LightOffCommandStu
type LightOffCommandStu struct {
	Light *LightStu
}

func (r *LightOffCommandStu) Init(Light *LightStu) {
	r.Light = Light
}

func (r *LightOffCommandStu) Execute() {
	r.Light.Off()
}

func (r *LightOffCommandStu) Undo() {
	r.Light.On()
}
