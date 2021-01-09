package aggr

type IterIntf interface {
	HasNext() bool
	Next() EntryIntf
}
