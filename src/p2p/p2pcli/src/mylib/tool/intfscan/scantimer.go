package intfscan

import (
	"sync"
	"time"
)

var (
	DAY int64 = 86400
	HOR int64 = 3600
	MIN int64 = 60
)

func NewScanTimerStu(Hor, Min, Period int64) ScanIntf {
	var ScanTimer ScanTimerStu
	ScanTimer.Init(Hor, Min, Period)
	return &ScanTimer
}

type ScanTimerStu struct {
	Aggregate AggregateIntf
	Time      *time.Timer
	Period    int64
	sync.RWMutex
}

func (r *ScanTimerStu) SetAggregate(Aggregate AggregateIntf) {
	r.Aggregate = Aggregate
}

func (r *ScanTimerStu) Init(Hor, Min, Period int64) {

	tobj := Hor*HOR + Min*MIN

	tnow := time.Now().Unix() + HOR*8
	tpass := tnow % DAY

	tleft := tpass
	if tleft >= tobj {
		tleft = DAY - tpass
	}
	r.Period = Period

	r.Time = time.NewTimer(time.Second * time.Duration(tleft))

}

func (r *ScanTimerStu) Run() {
	for {
		<-r.Time.C
		r.RunOnce()
		r.Time = time.NewTimer(time.Second * time.Duration(r.Period))
	}
}
func (r *ScanTimerStu) RunOnce() {
	Iter := r.Aggregate.Iterator()
	for Iter.HasNext() {
		Entry := Iter.Next()
		Entry.Update()
	}
}
