package a07builder

type BuilderIntf interface {
	MakeTitle(title string)
	MakeString(str string)
	MakeItem(items []string)
	Close()
}
