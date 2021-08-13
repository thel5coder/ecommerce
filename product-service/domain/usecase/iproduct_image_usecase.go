package usecase

type IProductImageUseCase interface {
	Add(productID, imageKey string) (err error)
	Delete(productID string) (err error)
	Store(productID string, imageKeys []string) (err error)
}
