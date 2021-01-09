package a23interpreter

type NodeProgramStu struct {
	Cmd NodeCmdListStu
}

func (r *NodeProgramStu) Parse(context *ContextStu) (ret int) {
	context.SkipToken("program")

	var CmdList NodeCmdListStu
	r.Cmd = CmdList
	r.Cmd.Parse(context)

	return
}

func (r *NodeProgramStu) ToString() (str string) {

	return "[ program " + r.Cmd.ToString() + " end ]"
}
