package a23interpreter

type NodeCmdPrimitiveStu struct {
	Name string
}

func (r *NodeCmdPrimitiveStu) Parse(context *ContextStu) (ret int) {
	r.Name = context.CurrentToken()
	context.SkipToken(r.Name)

	return
}
func (r *NodeCmdPrimitiveStu) ToString() (str string) {

	return r.Name
}
