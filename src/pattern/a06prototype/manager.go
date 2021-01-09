package a06prototype

type ManagerStu struct {
	Products map[string]ProductIntf
}

func (r *ManagerStu) Init() {
	r.Products = make(map[string]ProductIntf)
}

func (r *ManagerStu) Register(Name string, Product ProductIntf) {
	r.Products[Name] = Product
}
func (r *ManagerStu) Creat(Str string) (out ProductIntf, exist bool) {
	out, exist = r.Products[Str]
	return
}
