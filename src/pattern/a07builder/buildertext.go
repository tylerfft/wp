package a07builder

import (
	"bytes"
	"fmt"
)

type BuilderTextStu struct {
	Buf *bytes.Buffer
}

func (r *BuilderTextStu) Init() {
	r.Buf = new(bytes.Buffer)
}
func (r *BuilderTextStu) MakeTitle(title string) {

	fmt.Fprint(r.Buf, "==============================\n")
	fmt.Fprint(r.Buf, "["+title+"]\n")
	fmt.Fprint(r.Buf, "\n")
}
func (r *BuilderTextStu) MakeString(str string) {
	fmt.Fprint(r.Buf, "\t[]"+str+"\n")
	fmt.Fprint(r.Buf, "\n")
}
func (r *BuilderTextStu) MakeItem(items []string) {
	for _, v := range items {
		fmt.Fprint(r.Buf, "\t\t>"+v+"\n")
		fmt.Fprint(r.Buf, "\n")
	}
}
func (r *BuilderTextStu) Close() {
	fmt.Fprint(r.Buf, "==============================\n")
}

func (r *BuilderTextStu) GetResult() (str string) {
	return r.Buf.String()
}
