package models

type ProductImage struct {
	id        string
	productId string
	imageKey  string
}

func NewProductImageModel() *ProductImage{
	return &ProductImage{}
}

func (model *ProductImage) Id() string {
	return model.id
}

func (model *ProductImage) SetId(id string) *ProductImage {
	model.id = id

	return model
}

func (model *ProductImage) ProductId() string {
	return model.productId
}

func (model *ProductImage) SetProductId(productId string) *ProductImage {
	model.productId = productId

	return model
}

func (model *ProductImage) ImageKey() string {
	return model.imageKey
}

func (model *ProductImage) SetImageKey(imageKey string) *ProductImage {
	model.imageKey = imageKey

	return model
}
