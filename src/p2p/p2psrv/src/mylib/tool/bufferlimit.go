package tool

import (
	"sync"
	"sync/atomic"
)

var LimitSize int64 = 20

func NewLimitStu() *LimitStu {
	var Limit LimitStu
	Limit.Init(LimitSize)
	return &Limit
}

type LimitStu struct {
	Able   int64
	Size   int64
	Enable bool
	sync.RWMutex
}

func (r *LimitStu) Init(Size int64) {
	r.Lock()
	defer r.Unlock()
	r.Size = Size
}
func (r *LimitStu) FeedToken() {
	r.Lock()
	defer r.Unlock()

	atomic.AddInt64(&r.Able, r.Size)
	Able := atomic.LoadInt64(&r.Able)
	if Able >= 3*r.Size {
		atomic.SwapInt64(&r.Able, 3*r.Size)
	}
}

func (r *LimitStu) GetToken() (ret int) {
	r.Lock()
	defer r.Unlock()
	if !r.Enable {
		return 0
	}

	Able := atomic.LoadInt64(&r.Able)
	if Able <= 0 {
		ret = -1
	} else {
		atomic.AddInt64(&r.Able, -1)
		ret = 0
	}
	return
}

func (r *LimitStu) SetEnable(Enable bool) {
	r.Lock()
	defer r.Unlock()
	r.Enable = Enable
}
func (r *LimitStu) SetSize(Size int64) {
	r.Lock()
	defer r.Unlock()
	r.Size = Size
}
