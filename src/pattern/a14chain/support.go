package a14chain

import (
	"fmt"
)

type SupportIntf interface {
	Resolve(Trouble *TroubleStu) bool
	SetNext(Next *SupportStu) *SupportStu
	GetNext() *SupportStu
	ToString() string
}

type SupportStu struct {
	Name string
	Now  SupportIntf
}

func (r *SupportStu) Init(Name string) {
	r.Name = Name
}
func (r *SupportStu) SetNext(Next *SupportStu) *SupportStu {
	r.Now.SetNext(Next)
	return Next
}

func (r *SupportStu) Support(Trouble *TroubleStu) {

	if r.Now.Resolve(Trouble) {
		r.Done(Trouble)
	} else if r.Now.GetNext() != nil {
		r.Now.GetNext().Support(Trouble)
	} else {
		r.Fail(Trouble)
	}
}

func (r *SupportStu) GetName() (Name string) {
	return r.Name
}
func (r *SupportStu) ToString() (Str string) {
	return "[ Support " + r.Name + "]"
}
func (r *SupportStu) Done(Trouble *TroubleStu) {
	fmt.Println(Trouble.ToString() + " can  be solved by " + r.Now.ToString())
}
func (r *SupportStu) Fail(Trouble *TroubleStu) {
	fmt.Println(Trouble.ToString() + " can not be solved ")
}
