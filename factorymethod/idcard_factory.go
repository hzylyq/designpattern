package factorymethod

type IDCardFactory struct {
	owners []string
}

func (f *IDCardFactory) createProduct(owner string) IProduct {
	return New(owner)
}

func (f *IDCardFactory) registerProduct(product IProduct) {
	f.owners = append(f.owners, (product).(*IdCard).Owner)
}

func (f *IDCardFactory) getOwners() []string {
	return f.owners
}
