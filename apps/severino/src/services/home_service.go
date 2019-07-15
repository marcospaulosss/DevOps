package services

import (
	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/src/structs"
)

type HomeService struct {
	homeClient interfaces.Client
}

func NewHomeService(container rpc.Container) *HomeService {
	return &HomeService{
		homeClient: container.Elearning.HomeClient,
	}
}

func (this *HomeService) Create(item interface{}) (interface{}, error) {
	return structs.Result{}, nil
}

func (this *HomeService) Update(item interface{}) (interface{}, error) {
	return structs.Result{}, nil
}

func (this *HomeService) Delete(item interface{}) (interface{}, error) {
	return structs.Result{}, nil
}

func (this *HomeService) ReadOne(item interface{}) (interface{}, error) {
	return structs.Result{}, nil
}

func (this *HomeService) ReadAll(search structs.Search) (structs.Result, error) {
	return this.homeClient.ReadAll(search)
}
