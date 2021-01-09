package a09bridge

func TestFunc() {

	var DisplayMult DisplayMultStu

	var StringDisplay StringDisplayStu
	StringDisplay.Init("shi")

	DisplayMult.Init(&StringDisplay)
	DisFunc(&DisplayMult)

}

func DisFunc(Intf DisplayIntfStu) {
	Intf.Display()
}
