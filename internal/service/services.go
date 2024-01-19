package service

type Services struct {
	Products ProductService
	Reviews  ReviewService
}

func NewServices() Services {
	return Services{
		Products: ProductService{},
		Reviews:  ReviewService{},
	}
}
