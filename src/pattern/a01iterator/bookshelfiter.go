package a01iterator

type BookShelfIterStu struct {
	Obj *BookShelfStu
	Idx int
	Cnt int
}

func (r *BookShelfIterStu) Init(Obj *BookShelfStu) {
	r.Obj = Obj
	r.Cnt = Obj.GetLength()
	r.Idx = 0
}

func (r *BookShelfIterStu) HasNext() (exist bool) {
	if r.Idx < r.Cnt {
		return true
	}
	return false
}

func (r *BookShelfIterStu) Next() (out interface{}) {
	if r.Idx < r.Cnt {
		out = r.Obj.GetByIdx(r.Idx)
		r.Idx++
	}
	return
}
