package a12decorator

type DisplayStringStu struct {
	Str string
}

func (r *DisplayStringStu) Init(Str string) {
	r.Str = Str
}

func (r *DisplayStringStu) GetColumns() (ret int) {
	return len(r.Str)
}

func (r *DisplayStringStu) GetRows() (ret int) {
	return 1
}

func (r *DisplayStringStu) GetRowText(row int) (str string) {
	if 0 == row {
		str = r.Str
	} else {
		str = ""
	}
	return
}
