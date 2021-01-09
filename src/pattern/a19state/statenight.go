package a19state

type StateNightStu struct {
}

var StateNightGlobal StateNightStu

func (r *StateNightStu) DoClock(context ContextIntf, hour int) {
	if hour > 7 && 17 >= hour {
		context.ChangeState(&StateDayGlobal)
	}
}

func (r *StateNightStu) DoUse(context ContextIntf) {
	context.CallSecuCenter("use bank (night)")
}
func (r *StateNightStu) DoAlarm(context ContextIntf) {
	context.CallSecuCenter("press alarm (night)")
}
func (r *StateNightStu) DoPhone(context ContextIntf) {
	context.RecordLog("phone (night)")
}
func (r *StateNightStu) ToString() (str string) {
	return "(noight)"
}
