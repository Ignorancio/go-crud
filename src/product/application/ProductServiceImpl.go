package application

import "backend/src/product/domain"

type ProductServiceImpl struct {
	productRepository domain.ProductRepository
}

// NewProductService creates a new ProductServiceImpl with the given repository
func NewProductService(repo domain.ProductRepository) domain.ProductService {
	return &ProductServiceImpl{
		productRepository: repo,
	}
}

func (ps *ProductServiceImpl) FindAll() []domain.Product {
	return ps.productRepository.FindAll()
}

func (ps *ProductServiceImpl) FindById(id int) domain.Product {
	return ps.productRepository.FindById(id)
}

func (ps *ProductServiceImpl) Save(product domain.Product) domain.Product {
	return ps.productRepository.Save(product)
}

func (ps *ProductServiceImpl) Delete(id int) {
	ps.productRepository.Delete(id)
}

func (ps *ProductServiceImpl) Update(product domain.Product) domain.Product {
	return ps.productRepository.Update(product)
}
