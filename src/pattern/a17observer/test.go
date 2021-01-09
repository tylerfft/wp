package a17observer

func TestFunc() {
	Ob01 := &DigitObserverStu{}
	Ob02 := &GarphObserverStu{}
	Generator := &Numgenerator{}
	Generator.Register(Ob01)
	Generator.Register(Ob02)
	Generator.Exec()
}
