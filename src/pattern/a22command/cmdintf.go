package a22command

type CommandIntf interface {
	Execute()
	Undo()
}
