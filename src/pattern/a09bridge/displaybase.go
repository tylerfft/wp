package a09bridge

type DisplayBaseStu struct {
	Impl DisplayImplIf
}

func (r *DisplayBaseStu) Init(Impl DisplayImplIf) {
	r.Impl = Impl
}

func (r *DisplayBaseStu) Display() {
	r.Impl.Open()
	r.Impl.Print()
	r.Impl.Close()
}
