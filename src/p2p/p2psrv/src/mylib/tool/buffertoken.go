package tool

import (
	"sync/atomic"
)

var KByte int32 = 1000

func NewBufferTokenStu() *BufferTokenStu {
	var BufferToken BufferTokenStu
	BufferToken.Init(100*KByte, 10*KByte)
	return &BufferToken
}

type BufferTokenStu struct {
	Able int32
	Step int32
	Size int32
}

func (r *BufferTokenStu) Init(Size, Step int32) {
	r.Size = Size
	r.Step = Step
}

func (r *BufferTokenStu) FeedToken() {
	Able := atomic.LoadInt32(&r.Able)
	Size := atomic.LoadInt32(&r.Size)
	if Able < Size {
		atomic.AddInt32(&r.Able, r.Step)
	}
}

func (r *BufferTokenStu) GetToken(Need int32) (ret int) {
	if Need < 0 {
		return -1
	}

	Able := atomic.LoadInt32(&r.Able)
	if Able <= 0 {
		return -1
	}
	if Able < Need {
		ret = -1
	} else {
		atomic.AddInt32(&r.Able, -Need)
		ret = -1
	}
	return
}
