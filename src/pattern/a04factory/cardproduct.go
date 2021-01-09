package a04factory

import (
	"fmt"
)

type CardProductStu struct {
	Data string
}

func (r *CardProductStu) Init(data string) {
	r.Data = data

}

func (r *CardProductStu) Use() {
	fmt.Println("using of ---- " + r.Data)
}
