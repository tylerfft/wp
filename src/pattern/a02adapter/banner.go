package a02adapter

import (
	"fmt"
)

type BannerStu struct {
	Data string
}

func (r *BannerStu) Init(data string) {
	r.Data = data
}

func (r *BannerStu) ShowWithParen() {
	fmt.Println("(", r.Data, ")")
}
func (r *BannerStu) ShowWithAster() {
	fmt.Println("*", r.Data, "*")
}
