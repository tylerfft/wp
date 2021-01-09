package intfscan

func NewAggregateScanStu(Scan ScanIntf) *AggregateScanStu {
	var AggregateScan AggregateScanStu
	AggregateScan.Init(Scan)
	return &AggregateScan
}

type AggregateScanStu struct {
	Aggregate AggregateIntf
	Scan      ScanIntf
}

func (r *AggregateScanStu) Init(Scan ScanIntf) {
	r.Aggregate = NewAggregateStu()
	r.Scan = Scan
	r.Scan.SetAggregate(r.Aggregate)
	go r.Scan.Run()
}

func (r *AggregateScanStu) Add(Id string, Entry EntryIntf) (Idx int, ret int) {
	return r.Aggregate.Add(Id, Entry)
}
func (r *AggregateScanStu) Find(Id string) (Entry EntryIntf, exist bool) {
	return r.Aggregate.Find(Id)
}
