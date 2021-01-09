package a10strategy

type DuckStu struct {
	FlyTool FlyIntf
}

func (r *DuckStu) Init() {
	r.FlyTool = &FlyNoWayStu{}
}

func (r *DuckStu) SetFly(FlyTool FlyIntf) {
	r.FlyTool = FlyTool
}

func (r *DuckStu) Fly() {
	r.FlyTool.Fly()
}
