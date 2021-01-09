package a18memento

type MementoStu struct {
	Money  int
	Fruits []string
}

func (r *MementoStu) GetMoney() int {
	return r.Money
}

func (r *MementoStu) GetFruits() []string {
	return r.Fruits
}
func (r *MementoStu) AddFruits(Friut string) {
	r.Fruits = append(r.Fruits, Friut)
}
func (r *MementoStu) Init(Money int) {
	r.Money = Money
}
