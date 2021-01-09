package a19state

type StateIntf interface {
	DoClock(context ContextIntf, hour int)
	DoUse(context ContextIntf)
	DoAlarm(context ContextIntf)
	DoPhone(context ContextIntf)
	ToString() (str string)
}
