package a01iterator

type IterIntf interface {
	HasNext() bool
	Next() interface{}
}
