package a04factory

type FactoryIntf interface {
	CreateProduct() ProductIntf
}
