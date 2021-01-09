package a23interpreter

type NodeCmdListStu struct {
	List []NodeIntf
}

func (r *NodeCmdListStu) Parse(context *ContextStu) (ret int) {

	for true {
		if context.CurrentToken() == "" {
			ret = -1
			return
		} else if context.CurrentToken() == "end" {
			context.SkipToken("end")
			break
		} else {
			var Cmd NodeCmdStu
			Cmd.Parse(context)
			r.List = append(r.List, &Cmd)
		}

	}
	return
}
func (r *NodeCmdListStu) ToString() (str string) {
	for _, v := range r.List {
		str += v.ToString()
	}
	return
}
