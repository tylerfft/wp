package a10strategy

func TestFunc() {
	var Duck DuckStu
	Duck.Init()
	Duck.Fly()
	Duck.SetFly(&FlyWithWingsStu{})
	Duck.Fly()
	Duck.SetFly(&FlyWithRocketStu{})
	Duck.Fly()
}
