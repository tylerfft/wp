package a07builder

type DirectorStu struct {
	Builder BuilderIntf
}

func (r *DirectorStu) Init(Builder BuilderIntf) {
	r.Builder = Builder
}
func (r *DirectorStu) Construct() {
	r.Builder.MakeTitle("Greatting")
	r.Builder.MakeString("type time")
	r.Builder.MakeItem([]string{"mornig", "afternoon"})
	r.Builder.MakeString("type kind")
	r.Builder.MakeItem([]string{"ege", "bread"})
	r.Builder.Close()
}
