package a12decorator

func TestFunc() {
	var b1 DisplayStringStu
	b1.Init("shi")

	var b2 SideBorderStu
	b2.Init(&b1, "#")

	var b3 FullBorderStu
	b3.Init(&b2)
	var Dis DisplayBaseStu
	Dis.SetIntf(&b1)
	Dis.Show()
	Dis.SetIntf(&b2)
	Dis.Show()
	Dis.SetIntf(&b3)
	Dis.Show()
	b4 := &SideBorderStu{
		Ch: "$",
		Display: &FullBorderStu{
			Display: &SideBorderStu{
				Ch: "=",
				Display: &DisplayStringStu{
					Str: "hello"}}}}

	Dis.SetIntf(b4)
	Dis.Show()
}
