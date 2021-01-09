package a04factory

func TestFunc() {
	var Fact CardFactoryStu
	var Card01 ProductIntf = Fact.CreateProduct("shi")
	var Card02 ProductIntf = Fact.CreateProduct("zxc")
	var Card03 ProductIntf = Fact.CreateProduct("asd")
	Card01.Use()
	Card02.Use()
	Card03.Use()

	DisFunc()

}

func DisFunc() {

}
