package a19state

type ContextIntf interface {
	SetClock(hour int)
	ChangeState(State StateIntf)
	CallSecuCenter(msg string)
	RecordLog(msg string)
}
