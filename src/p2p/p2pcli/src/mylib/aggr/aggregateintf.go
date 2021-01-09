package aggr

type AggregateIntf interface {
	Iterator() IterIntf
	Add(Id string, Token EntryIntf) (Idx, ret int)
	Find(Id string) (Token EntryIntf, exist bool)
	GetByIdx(int) (Token EntryIntf, exist bool)
	Delete(Id string)
	Length() int
}
