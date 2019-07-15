package services

import (
	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/src/structs"
)

type ProductService struct {
	productClient interfaces.Client
}

func NewProductService(container rpc.Container) *ProductService {
	return &ProductService{
		productClient: container.Ecommerce.ProductClient,
	}
}

func (this *ProductService) Create(item interface{}) (interface{}, error) {
	return this.productClient.Create(item)
}

func (this *ProductService) Update(item interface{}) (interface{}, error) {
	return this.productClient.Update(item)
}

func (this *ProductService) Delete(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this *ProductService) ReadOne(item interface{}) (interface{}, error) {
	return this.productClient.ReadOne(item)
}

func (this *ProductService) ReadAll(search structs.Search) (structs.Result, error) {
	return this.productClient.ReadAll(search)
}
