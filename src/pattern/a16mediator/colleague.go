package a16mediator

type ColleagueIntf interface {
	SerMidiator(MediatorIntf)
	SetColleagueEnabled()
}

type ColleagueButtonStu struct {
	Mediator MediatorIntf
	Name     string
	Enabled  bool
}

func (r *ColleagueButtonStu) SetMidiator(Mediator MediatorIntf) {
	r.Mediator = Mediator
}

func (r *ColleagueButtonStu) SetColleagueEnabled(Enabled bool) {
	r.Enabled = Enabled
}

type ColleagueTextStu struct {
	Mediator MediatorIntf
	Text     string
	Enabled  bool
}

func (r *ColleagueTextStu) SetMidiator(Mediator MediatorIntf) {
	r.Mediator = Mediator
}

func (r *ColleagueTextStu) SetColleagueEnabled(Enabled bool) {
	r.Enabled = Enabled
}

func (r *ColleagueTextStu) SetText(Text string) {
	r.Text = Text
	r.Mediator.ColleagueChange()
}
func (r *ColleagueTextStu) GetText() (Text string) {
	return r.Text
}

type ColleagueCheckStu struct {
	Mediator MediatorIntf
	Enabled  bool
}

func (r *ColleagueCheckStu) SetMidiator(Mediator MediatorIntf) {
	r.Mediator = Mediator
}

func (r *ColleagueCheckStu) SetColleagueEnabled(Enabled bool) {
	r.Enabled = Enabled
	r.Mediator.ColleagueChange()
}

func (r *ColleagueCheckStu) GetState() bool {
	return r.Enabled
}
func (r *ColleagueCheckStu) SetState(Enabled bool) {
	r.Enabled = Enabled
	r.Mediator.ColleagueChange()
}
