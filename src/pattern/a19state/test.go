package a19state

import (
	"time"
)

func TestFunc() {
	var Bank SafeBankStu
	Bank.Init(&StateDayGlobal)
	for i := 0; i < 24; i += 2 {
		Bank.SetClock(i)
		time.Sleep(time.Second * 1)
		Bank.UseFacade()
	}
}

func DisFunc() {

}
