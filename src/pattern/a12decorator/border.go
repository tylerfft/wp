package a12decorator

import (
	"bytes"
)

type SideBorderStu struct {
	Ch      string
	Display DisplayIntf
}

func (r *SideBorderStu) Init(Display DisplayIntf, Ch string) {
	r.Display = Display
	r.Ch = Ch
}

func (r *SideBorderStu) GetColumns() (ret int) {
	return 1 + r.Display.GetColumns() + 1
}
func (r *SideBorderStu) GetRows() (ret int) {
	return r.Display.GetRows()
}

func (r *SideBorderStu) GetRowText(row int) (str string) {
	return r.Ch + r.Display.GetRowText(row) + r.Ch
}

type FullBorderStu struct {
	Display DisplayIntf
}

func (r *FullBorderStu) Init(Display DisplayIntf) {
	r.Display = Display
}

func (r *FullBorderStu) GetColumns() (ret int) {
	return 1 + r.Display.GetColumns() + 1
}
func (r *FullBorderStu) GetRows() (ret int) {
	return 1 + r.Display.GetRows() + 1
}

func (r *FullBorderStu) GetRowText(row int) (str string) {
	if 0 == row {
		var tmp []byte
		tmp = append(tmp, '+')
		rst := bytes.Repeat(tmp, 2+r.Display.GetColumns())
		str = string(rst)
	} else if r.Display.GetRows()+1 == row {
		var tmp []byte
		tmp = append(tmp, '+')
		rst := bytes.Repeat(tmp, 2+r.Display.GetColumns())
		str = string(rst)
	} else {
		str = "|" + r.Display.GetRowText(row-1) + "|"
	}
	return
}
