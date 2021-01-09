package a19state

import (
	"fmt"
)

type SafeBankStu struct {
	State StateIntf
}

func (r *SafeBankStu) Init(State StateIntf) {
	r.State = State
}
func (r *SafeBankStu) SetClock(hour int) {
	r.State.DoClock(r, hour)
}
func (r *SafeBankStu) ChangeState(State StateIntf) {
	fmt.Println("Change from " + r.State.ToString() + " to " + State.ToString() + "------------------------------------------")
	r.State = State
}
func (r *SafeBankStu) CallSecuCenter(msg string) {
	fmt.Println("call !  ... " + msg)
}
func (r *SafeBankStu) RecordLog(msg string) {
	fmt.Println("recording ... " + msg)
}

func (r *SafeBankStu) UseFacade() {
	r.State.DoAlarm(r)
	r.State.DoPhone(r)
	r.State.DoUse(r)
}
