package a13visitor

type EntryIntf interface {
	GetName() (Name string)
	GetSize() (Size int)
	ToString() (Str string)
	Accept(Visitor VisitorIntf)
}
