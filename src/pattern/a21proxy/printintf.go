package a21proxy

type PrintIntf interface {
	SetName(str string)
	GetName() (str string)
	Print(str string)
}
