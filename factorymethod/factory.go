package factorymethod

type IFactory interface {
	createProduct(owner string) IProduct
	registerProduct(product IProduct)
}

type Factory struct {
	AbstractFactory IFactory
}

func (f *Factory) create(owner string) IProduct {
	p := f.AbstractFactory.createProduct(owner)
	f.AbstractFactory.registerProduct(p)
	return p
}
