package a14chain

type NoSupportStu struct {
	Name string
	Next *SupportStu
}

func (r *NoSupportStu) Resolve(Trouble *TroubleStu) bool {
	return false
}
func (r *NoSupportStu) SetNext(Next *SupportStu) *SupportStu {
	r.Next = Next
	return Next
}

func (r *NoSupportStu) GetNext() *SupportStu {
	return r.Next
}
func (r *NoSupportStu) ToString() string {
	return "[ NoNoSupportStu ]"
}

type LimitSupportStu struct {
	Name  string
	Limit int
	Next  *SupportStu
}

func (r *LimitSupportStu) Resolve(Trouble *TroubleStu) bool {
	if r.Limit > Trouble.Get() {
		return true
	}
	return false
}
func (r *LimitSupportStu) SetNext(Next *SupportStu) *SupportStu {
	r.Next = Next
	return Next
}

func (r *LimitSupportStu) GetNext() *SupportStu {

	return r.Next
}
func (r *LimitSupportStu) ToString() string {
	return "[ LimitSupportStu ]"
}

type OddSupportStu struct {
	Name string
	Next *SupportStu
}

func (r *OddSupportStu) Resolve(Trouble *TroubleStu) bool {
	if Trouble.Get()%2 == 1 {
		return true
	}
	return false
}
func (r *OddSupportStu) SetNext(Next *SupportStu) *SupportStu {
	r.Next = Next
	return Next
}

func (r *OddSupportStu) GetNext() *SupportStu {

	return r.Next
}
func (r *OddSupportStu) ToString() string {
	return "[ OddSupportStu ]"
}

type SpecialSupportStu struct {
	Name   string
	Secial int
	Next   *SupportStu
}

func (r *SpecialSupportStu) Resolve(Trouble *TroubleStu) bool {
	if Trouble.Get() == r.Secial {
		return true
	}
	return false
}
func (r *SpecialSupportStu) SetNext(Next *SupportStu) *SupportStu {
	r.Next = Next
	return Next
}

func (r *SpecialSupportStu) GetNext() *SupportStu {

	return r.Next
}
func (r *SpecialSupportStu) ToString() string {
	return "[ SpecialSupportStu ]"
}
