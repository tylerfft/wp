package intfscan

import (
	"fmt"
	libf "libfunc"
	"testing"
	"time"
)

func Test_Scan(t *testing.T) {
	Pool := NewAggregateStu()
	Scan := NewScanTickStu(1)
	Scan.SetAggregate(Pool)
	go Scan.Run()
	Entry001 := NewEntryTestStu("001")
	Entry002 := NewEntryTestStu("002")
	Entry003 := NewEntryTestStu("003")
	Entry004 := NewEntryTestStu("004")
	Entry005 := NewEntryTestStu("005")

	Pool.Add(Entry001.GetId(), Entry001)
	time.Sleep(time.Millisecond * 1100)
	fmt.Println("\nadd 001")

	Pool.Add(Entry002.GetId(), Entry002)
	time.Sleep(time.Millisecond * 1100)
	fmt.Println("\nadd 002")

	Pool.Add(Entry003.GetId(), Entry003)
	time.Sleep(time.Millisecond * 1100)
	fmt.Println("\nadd 003")

	Pool.Add(Entry004.GetId(), Entry004)
	time.Sleep(time.Millisecond * 1100)
	fmt.Println("\nadd 004")

	Pool.Delete("001")
	time.Sleep(time.Millisecond * 1020)
	fmt.Println("\ndel 004")

	Pool.Delete("002")
	time.Sleep(time.Millisecond * 1100)
	fmt.Println("\ndel 002")

	Pool.Delete("003")
	time.Sleep(time.Millisecond * 1100)
	fmt.Println("\ndel 003")

	Pool.Add(Entry005.GetId(), Entry005)
	time.Sleep(time.Millisecond * 1100)
	fmt.Println("\nadd 005")

	Pool.Add(Entry002.GetId(), Entry002)
	time.Sleep(time.Millisecond * 1100)
	fmt.Println("\nadd 002")

	Pool.Add(Entry003.GetId(), Entry003)
	time.Sleep(time.Millisecond * 1100)
	fmt.Println("\nadd 003")

	Pool.Delete("001")
	time.Sleep(time.Millisecond * 1020)
	fmt.Println("\ndel 001")

	Pool.Add(Entry001.GetId(), Entry001)
	time.Sleep(time.Millisecond * 1020)
	fmt.Println("\nadd 001")

	Pool.Delete("003")
	time.Sleep(time.Millisecond * 1020)
	fmt.Println("\ndel 003")

	Pool.Delete("004")
	time.Sleep(time.Millisecond * 1020)
	fmt.Println("\ndel 004")

	Pool.Delete("001")
	time.Sleep(time.Millisecond * 1020)
	fmt.Println("\ndel 001")

	Pool.Delete("005")
	time.Sleep(time.Millisecond * 1020)
	fmt.Println("\ndel 005")

}

func NewEntryTestStu(Id string) *EntryTestStu {
	var Entry EntryTestStu
	Entry.Init(Id)
	return &Entry
}

type EntryTestStu struct {
	Id string
}

func (r *EntryTestStu) Init(Id string) {
	r.Id = Id
}
func (r *EntryTestStu) ToJson() string {
	return libf.SructToJsonStringOne(r)
}
func (r *EntryTestStu) Update() (ret int) {
	fmt.Print(r.Id, ",")
	return
}
func (r *EntryTestStu) GetId() (ID string) {
	return r.Id
}
func (r *EntryTestStu) GetData() interface{} {
	return r
}
