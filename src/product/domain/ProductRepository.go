package domain

type ProductRepository interface {
	FindAll() []Product
	FindById(id int) Product
	Save(product Product) Product
	Delete(id int)
	Update(product Product) Product
}
