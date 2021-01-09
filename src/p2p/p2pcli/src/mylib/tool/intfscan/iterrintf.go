package intfscan

type IterIntf interface {
	HasNext() bool
	Next() EntryIntf
}
