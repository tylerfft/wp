package a22command

import (
	"fmt"
)

func TestFunc() {
	fmt.Println("normal  ------------------------------------------------")
	var Remote RemoteStu
	var Light LightStu
	var LightOnCmd LightOnCommandStu
	LightOnCmd.Init(&Light)
	var LightOffCmd LightOffCommandStu
	LightOffCmd.Init(&Light)

	Remote.SetCmd(&LightOnCmd)
	Remote.Press()

	fmt.Println("undo   ------------------------------------------------")
	var RemoteUndo RemoteUndoStu
	RemoteUndo.Init()
	RemoteUndo.SetCmd(0, &LightOnCmd, &LightOffCmd)
	RemoteUndo.OnPress(0)
	RemoteUndo.Undo()

}
