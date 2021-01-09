package a14chain

func TestFunc() {
	alice := &SupportStu{Now: &NoSupportStu{Name: "alice"}}
	bob := &SupportStu{Now: &LimitSupportStu{Name: "bob", Limit: 10}}
	charlie := &SupportStu{Now: &SpecialSupportStu{Name: "charlie", Secial: 23}}
	diana := &SupportStu{Now: &LimitSupportStu{Name: "diana", Limit: 88}}
	elmo := &SupportStu{Now: &OddSupportStu{Name: "elmo"}}
	fred := &SupportStu{Now: &LimitSupportStu{Name: "fred", Limit: 1000}}

	alice.SetNext(bob).SetNext(charlie).SetNext(diana).SetNext(elmo).SetNext(fred)
	for i := 0; i < 100; i += 9 {
		alice.Support(&TroubleStu{Num: i})
	}
}
