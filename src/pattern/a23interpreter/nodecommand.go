package a23interpreter

type NodeCmdStu struct {
	Node NodeIntf
}

func (r *NodeCmdStu) Parse(context *ContextStu) (ret int) {

	if context.CurrentToken() == "repeat" {
		var CmdRepeat NodeCmdRepeatStu
		ret = CmdRepeat.Parse(context)
		r.Node = &CmdRepeat
	} else {
		var CmdPrimitive NodeCmdPrimitiveStu
		ret = CmdPrimitive.Parse(context)
		r.Node = &CmdPrimitive
	}

	return
}
func (r *NodeCmdStu) ToString() (str string) {

	return r.Node.ToString()
}
