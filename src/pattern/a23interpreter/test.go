package a23interpreter

import (
	"fmt"
)

func TestFunc() {

	var Context ContextStu
	Context.Init("program  repeat 3 go repeat left end end")
	fmt.Println(Context.GetData())
	var NodeProgram NodeProgramStu

	NodeProgram.Parse(&Context)
	fmt.Println(NodeProgram.ToString())

}
