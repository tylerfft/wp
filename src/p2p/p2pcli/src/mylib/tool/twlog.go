package tool

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

func NewTwLogStu(Format string) *TwLogStu {
	var Log TwLogStu
	Log.LogInit(Format)
	return &Log
}

type TwLogStu struct {
	Format string
	Tw     *tabwriter.Writer
	Buf    *bytes.Buffer
}

func (r *TwLogStu) LogInit(Format string) (ret int) {
	r.Format = Format
	r.Buf = new(bytes.Buffer)
	r.Tw = new(tabwriter.Writer).Init(r.Buf, 0, 8, 2, ' ', 0)
	return
}

func (r *TwLogStu) Print(args ...interface{}) {
	fmt.Fprintf(r.Tw, r.Format, args...)
}
func (r *TwLogStu) Flush() string {
	r.Tw.Flush()
	return r.Buf.String()
}
