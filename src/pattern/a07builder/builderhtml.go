package a07builder

type BuilderHtmlStu struct {
	Buf *bytes.Buffer
}

func (r *BuilderHtmlStu) Init() {
	r.Buf = new(bytes.Buffer)
}
func (r *BuilderHtmlStu) MakeTitle(title string) {

}
func (r *BuilderHtmlStu) MakeString(str string) {

}
func (r *BuilderHtmlStu) MakeItem(items []string) {

}
func (r *BuilderHtmlStu) Close() {

}
