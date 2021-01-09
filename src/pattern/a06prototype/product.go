package a06prototype

type ProductIntf interface {
	Use(string)
	Clone() ProductIntf
}
