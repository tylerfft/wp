package a16mediator

import (
	"fmt"
)

type MediatorIntf interface {
	ColleagueChange()
}

type FrameStu struct {
	CheckGuest   ColleagueCheckStu
	CheckLogin   ColleagueCheckStu
	TextUser     ColleagueTextStu
	TextPass     ColleagueTextStu
	ButtonOk     ColleagueButtonStu
	ButtonCancle ColleagueButtonStu
}

func (r *FrameStu) Init() {

	r.CheckGuest.SetMidiator(r)
	r.CheckLogin.SetMidiator(r)
	r.TextPass.SetMidiator(r)
	r.TextUser.SetMidiator(r)
	r.ButtonCancle.SetMidiator(r)
	r.ButtonCancle.SetMidiator(r)

	r.CheckGuest.SetColleagueEnabled(true)
	r.CheckLogin.SetColleagueEnabled(false)
}

func (r *FrameStu) ColleagueChange() {
	if r.CheckGuest.GetState() { //guest model
		r.TextUser.SetColleagueEnabled(false)
		r.TextUser.SetColleagueEnabled(false)
		r.ButtonOk.SetColleagueEnabled(true)
	} else {
		r.TextUser.SetColleagueEnabled(true)
		r.UserPassChanged()
	}
}
func (r *FrameStu) UserPassChanged() {
	if len(r.TextUser.GetText()) > 0 {
		fmt.Println("r.TextPass.SetColleagueEnabled(true)")
		r.TextPass.SetColleagueEnabled(true)
		if len(r.TextPass.GetText()) > 0 {
			fmt.Println("r.ButtonOk.SetColleagueEnabled(true)")
			r.ButtonOk.SetColleagueEnabled(true)
		} else {
			fmt.Println("r.ButtonOk.SetColleagueEnabled(false)")
			r.ButtonOk.SetColleagueEnabled(false)
		}
	} else {
		fmt.Println("r.TextPass.SetColleagueEnabled(false)")
		r.TextPass.SetColleagueEnabled(false)
		fmt.Println("r.ButtonOk.SetColleagueEnabled(false)")
		r.ButtonOk.SetColleagueEnabled(false)
	}
}

func (r *FrameStu) Exec() {
	r.CheckGuest.SetState(false)
	fmt.Println("r.TextUser.SetText( username ) ------------------------")
	r.TextUser.SetText("username")
	fmt.Println("r.TextPass.SetText( password ) ------------------------")
	r.TextPass.SetText("password")

}
