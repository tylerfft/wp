package a12decorator

import (
	"fmt"
)

type DisplayIntf interface {
	GetColumns() (ret int)
	GetRows() (ret int)
	GetRowText(row int) (str string)
}

type DisplayBaseStu struct {
	Display DisplayIntf
}

func (r *DisplayBaseStu) SetIntf(Display DisplayIntf) {
	r.Display = Display
}

func (r *DisplayBaseStu) Show() {
	for i := 0; i < r.Display.GetRows(); i++ {
		fmt.Println(r.Display.GetRowText(i))
	}
}
