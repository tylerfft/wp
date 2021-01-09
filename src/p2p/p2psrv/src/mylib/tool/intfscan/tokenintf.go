package intfscan

type EntryIntf interface {
	Update() int
	GetData() interface{}
	GetId() string
}
