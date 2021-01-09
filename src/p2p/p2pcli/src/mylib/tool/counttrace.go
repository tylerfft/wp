package tool

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type CyclicSliceStu struct {
	Init     bool
	size     int
	idx      int
	datasum  int64
	dataList []int64
}

func (r *CyclicSliceStu) Initialize(sliceSize int) {
	if sliceSize <= 0 {
		r.Init = false
		return
	}

	r.Init = true
	r.datasum = 0
	r.size = sliceSize

	r.dataList = make([]int64, sliceSize)
	for i := 0; i < sliceSize; i++ {
		r.dataList[i] = 0
	}

}

func (r *CyclicSliceStu) Add(data int64) {

	if !r.Init {
		return
	}

	if r.idx >= r.size {
		r.idx = 0
	}
	r.datasum = r.datasum - r.dataList[r.idx]
	r.dataList[r.idx] = data
	r.datasum = r.datasum + data

	r.idx++
}

func (r *CyclicSliceStu) GetAvr() (ret int64) {
	if !r.Init {
		return
	}

	ret = r.datasum / int64(r.size)
	return
}

func NewTrace(name string) *TraceStu {
	var Trace TraceStu
	Trace.Init(name)
	return &Trace
}

type TraceStu struct {
	cnt              int64
	cntpermeasuresum int64
	totalsum         int64
	cntPerSec        int64
	tsNano           int64
	min1Tmp          CyclicSliceStu
	min5Tmp          CyclicSliceStu
	min10Tmp         CyclicSliceStu
	name             string
}

func (r *TraceStu) Init(name string) {
	r.name = name
	r.cnt = 0
	r.cntpermeasuresum = 0
	r.totalsum = 0
	r.cntPerSec = 0

	r.tsNano = time.Now().UnixNano()
	r.min1Tmp.Initialize(6)
	r.min5Tmp.Initialize(30)
	r.min10Tmp.Initialize(60)
}

func (r *TraceStu) toString() (name string) {
	return r.name
}

func (r *TraceStu) Incr() {
	atomic.AddInt64(&r.cnt, 1)
	atomic.AddInt64(&r.cntpermeasuresum, 1)
	atomic.AddInt64(&r.totalsum, 1)
}
func (r *TraceStu) GetTotal() int64 {
	return atomic.LoadInt64(&r.totalsum)
}
func (r *TraceStu) Decr() {
	atomic.AddInt64(&r.cnt, -1)
}

func (r *TraceStu) GetCnt() int64 {
	return atomic.LoadInt64(&r.cnt)
}
func (r *TraceStu) GetCntPerSec() int64 {
	rtsum := atomic.LoadInt64(&r.cntpermeasuresum)
	rtTs := time.Now().UnixNano() - r.tsNano
	if rtTs == 0 {
		rtTs = 1
	}
	secAvr := int64(float64(rtsum) / (float64(rtTs) / 1.0e9))
	return secAvr
}

func (r *TraceStu) cntsumReset() {
	r.tsNano = time.Now().UnixNano()
	atomic.SwapInt64(&r.cntpermeasuresum, 0)
}

func (r *TraceStu) Measure() {
	cur := r.GetCnt()

	r.min1Tmp.Add(cur)
	r.min5Tmp.Add(cur)
	r.min10Tmp.Add(cur)

	r.cntPerSec = r.GetCntPerSec()
	r.cntsumReset()
}

func (r *TraceStu) Print() {

	t := time.Now()
	fmt.Println(r.toString(),
		"===== \t1m ->", r.min1Tmp.GetAvr(),
		"\t5m ->", r.min5Tmp.GetAvr(),
		"\t10m ->", r.min10Tmp.GetAvr(),
		"\tcntPerSec ->", r.cntPerSec,
		"\ttime at ->", t.Format("2006-01-02 15:04:05"))

	return
}

var SingleTraceMan TraceManStu
var TraceManOnce sync.Once

func GetTraceMan() *TraceManStu {
	return &SingleTraceMan
}
func NewTraceMan() *TraceManStu {
	var TraceMan TraceManStu
	TraceMan.Init()
	return &TraceMan
}

type TraceManStu struct {
	TarceList []*TraceStu
}

func (r *TraceManStu) Init() {

}
func (r *TraceManStu) Rgister(Trace *TraceStu) {

	r.TarceList = append(r.TarceList, Trace)
}

func (r *TraceManStu) MeasurePrint() {

	for _, v := range r.TarceList {
		(*v).Measure()
	}
	for _, v := range r.TarceList {
		(*v).Print()
	}
}

func (r *TraceManStu) TickerFunc() {

	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for range ticker.C {
			r.MeasurePrint()
		}
	}()
}
