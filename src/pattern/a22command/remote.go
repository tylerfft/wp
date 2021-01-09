package a22command

type RemoteStu struct {
	Cmd CommandIntf
}

func (r *RemoteStu) SetCmd(Cmd CommandIntf) {
	r.Cmd = Cmd
}

func (r *RemoteStu) Press() {
	r.Cmd.Execute()
}
