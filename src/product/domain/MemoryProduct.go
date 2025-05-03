package domain

type MemoryProduct struct {
	products []Product
	index    int
}

// NewMemoryRepository creates a new in-memory product repository
func NewMemoryRepository() *MemoryProduct {
	return &MemoryProduct{
		products: make([]Product, 0),
		index:    0,
	}
}

func (repo *MemoryProduct) FindAll() []Product {
	return repo.products
}

func (repo *MemoryProduct) FindById(id int) Product {
	for _, product := range repo.products {
		if product.Id == id {
			return product
		}
	}
	return Product{}
}

func (repo *MemoryProduct) Save(product Product) Product {
	product.Id = repo.index + 1
	repo.index++
	repo.products = append(repo.products, product)
	return product
}

func (repo *MemoryProduct) Delete(id int) {
	for i, product := range repo.products {
		if product.Id == id {
			repo.products = append(repo.products[:i], repo.products[i+1:]...)
			break
		}
	}
}

func (repo *MemoryProduct) Update(product Product) Product {
	for i, p := range repo.products {
		if p.Id == product.Id {
			repo.products[i] = product
			break
		}
	}
	return product
}
