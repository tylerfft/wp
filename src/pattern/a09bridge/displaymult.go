package a09bridge

type DisplayMultStu struct {
	Impl DisplayImplIf
}

func (r *DisplayMultStu) Init(Impl DisplayImplIf) {
	r.Impl = Impl
}

func (r *DisplayMultStu) Display() {
	r.Impl.Open()
	for i := 0; i < 5; i++ {
		r.Impl.Print()
	}
	r.Impl.Close()
}
