package a23interpreter

type NodeIntf interface {
	Parse(*ContextStu) int
	ToString() (str string)
}
