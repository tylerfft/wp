package a21proxy

type PrinterProxyStu struct {
	Name     string
	PrinterP *PrinterStu
}

func (r *PrinterProxyStu) SetName(name string) {

	if r.PrinterP != nil {
		r.PrinterP.SetName(name)
	}
	r.Name = name
}

func (r *PrinterProxyStu) GetName() (name string) {
	return r.Name
}

func (r *PrinterProxyStu) Print(str string) {
	if r.PrinterP == nil {
		r.Realize()
	}
	r.PrinterP.Print(str)
}

func (r *PrinterProxyStu) Realize() {
	var Printer PrinterStu
	Printer.Init(r.Name)
	r.PrinterP = &Printer
}
