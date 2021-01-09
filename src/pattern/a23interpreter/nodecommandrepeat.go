package a23interpreter

type NodeCmdRepeatStu struct {
	Num     string
	CmdList NodeIntf
}

func (r *NodeCmdRepeatStu) Parse(context *ContextStu) (ret int) {
	context.SkipToken("repeat")
	r.Num = context.CurrentToken()
	context.NextToken()
	var CmdList NodeCmdListStu
	CmdList.Parse(context)
	r.CmdList = &CmdList

	return
}
func (r *NodeCmdRepeatStu) ToString() (str string) {
	return "[ repeate " + r.Num + " " + r.CmdList.ToString() + " end ]"
}
