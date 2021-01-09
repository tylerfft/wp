package a06prototype

func TestFunc() {
	var Manager ManagerStu
	Manager.Init()
	mbox := &MessageBoxStu{Deco: "*"}
	sbox := &MessageBoxStu{Deco: "/"}
	upen := &UnderLineStu{Unch: "~"}

	Manager.Register("mbox", mbox)
	Manager.Register("sbox", sbox)
	Manager.Register("upen", upen)

	Prod01, _ := Manager.Creat("mbox")
	Prod02, _ := Manager.Creat("sbox")
	Prod03, _ := Manager.Creat("upen")

	Prod01.Use("hello")
	Prod02.Use("hello")
	Prod03.Use("hello")

}
