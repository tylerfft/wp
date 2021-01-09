package a22command

type RemoteUndoStu struct {
	CmdOnS  []CommandIntf
	CmdOffS []CommandIntf
	CmdUndo CommandIntf
}

func (r *RemoteUndoStu) Init() {

	r.CmdOnS = make([]CommandIntf, 7)
	r.CmdOffS = make([]CommandIntf, 7)

}
func (r *RemoteUndoStu) SetCmd(slot int, CmdOn CommandIntf, CmdOff CommandIntf) {
	r.CmdOnS[slot] = CmdOn
	r.CmdOffS[slot] = CmdOff
}
func (r *RemoteUndoStu) OnPress(slot int) {
	r.CmdOnS[slot].Execute()
	r.CmdUndo = r.CmdOnS[slot]
}
func (r *RemoteUndoStu) OffPress(slot int) {
	r.CmdOffS[slot].Execute()
	r.CmdUndo = r.CmdOffS[slot]
}

func (r *RemoteUndoStu) Undo() {
	r.CmdUndo.Undo()
}
