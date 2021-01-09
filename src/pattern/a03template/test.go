package a03template

func TestFunc() {

	var CharDisplay CharDisplayStu
	CharDisplay.Init("shi")

	var DisplayBase DisplayBaseStu

	DisplayBase.Init(&CharDisplay)
	DisFunc(&DisplayBase)

	var StringDisplay StringDisplayStu
	StringDisplay.Init("shi")

	DisplayBase.Init(&StringDisplay)
	DisFunc(&DisplayBase)

}

func DisFunc(disbase *DisplayBaseStu) {
	disbase.Display()

}
