package intfscan

import (
	"sync"
	"time"
)

func NewScanTickStu(Period int) ScanIntf {
	var ScanTick ScanTickStu
	ScanTick.Init(Period)
	return &ScanTick
}

type ScanTickStu struct {
	Aggregate AggregateIntf
	sync.RWMutex
	Tick *time.Ticker
}

func (r *ScanTickStu) Init(Period int) {
	r.Tick = time.NewTicker(time.Second * time.Duration(Period))
}
func (r *ScanTickStu) SetAggregate(Aggregate AggregateIntf) {
	r.Aggregate = Aggregate
}

func (r *ScanTickStu) Run() {
	for range r.Tick.C {
		r.RunOnce()
	}
}
func (r *ScanTickStu) RunOnce() {
	Iter := r.Aggregate.Iterator()
	for Iter.HasNext() {
		Entry := Iter.Next()
		if Entry != nil {
		Entry.Update()
	}
}
}
