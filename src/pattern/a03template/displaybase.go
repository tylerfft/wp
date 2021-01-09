package a03template

type DisplayBaseStu struct {
	dis DisplayIf
}

func (r *DisplayBaseStu) Init(dis DisplayIf) {
	r.dis = dis
}

func (r *DisplayBaseStu) Display() {
	r.dis.Open()
	for i := 0; i < 5; i++ {
		r.dis.Print()
	}
	r.dis.Close()
}
