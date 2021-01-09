package a19state

type StateDayStu struct {
}

var StateDayGlobal StateDayStu

func (r *StateDayStu) DoClock(context ContextIntf, hour int) {
	if hour < 7 || 17 <= hour {
		context.ChangeState(&StateNightGlobal)
	}
}

func (r *StateDayStu) DoUse(context ContextIntf) {
	context.RecordLog("use bank (day)")
}
func (r *StateDayStu) DoAlarm(context ContextIntf) {
	context.CallSecuCenter("press alarm (day)")
}
func (r *StateDayStu) DoPhone(context ContextIntf) {
	context.CallSecuCenter("phone (day)")
}
func (r *StateDayStu) ToString() (str string) {
	return "(day)"
}
