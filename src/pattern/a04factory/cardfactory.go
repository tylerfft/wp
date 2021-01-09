package a04factory

type CardFactoryStu struct {
}

func (r *CardFactoryStu) CreateProduct(data string) (prod ProductIntf) {
	var Prod CardProductStu
	Prod.Init(data)
	prod = &Prod
	return
}
