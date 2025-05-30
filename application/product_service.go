package application

type ProductService struct {
	ProductPersister ProductPersisterInterface
}

func NewProductService(persister ProductPersisterInterface) *ProductService {
	return &ProductService{
		ProductPersister: persister,
	}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.ProductPersister.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	if valid, err := product.IsValid(); !valid {
		return nil, err
	}

	savedProduct, err := s.ProductPersister.Save(product)
	if err != nil {
		return nil, err
	}
	return savedProduct, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	if err := product.Enable(); err != nil {
		return nil, err
	}

	savedProduct, err := s.ProductPersister.Save(product)
	if err != nil {
		return nil, err
	}
	return savedProduct, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	if err := product.Disable(); err != nil {
		return nil, err
	}

	savedProduct, err := s.ProductPersister.Save(product)
	if err != nil {
		return nil, err
	}
	return savedProduct, nil
}
