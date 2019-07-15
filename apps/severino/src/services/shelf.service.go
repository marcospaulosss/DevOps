package services

import (
	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/src/structs"
)

type ShelfService struct {
	shelfClient interfaces.Client
}

func NewShelfService(container rpc.Container) *ShelfService {
	return &ShelfService{
		shelfClient: container.Elearning.ShelfClient,
	}
}

func (this *ShelfService) Create(item interface{}) (interface{}, error) {
	return this.shelfClient.Create(item)
}

func (this *ShelfService) Update(item interface{}) (interface{}, error) {
	return this.shelfClient.Update(item)
}

func (this *ShelfService) Delete(item interface{}) (interface{}, error) {
	return this.shelfClient.Delete(item)
}

func (this *ShelfService) ReadOne(item interface{}) (interface{}, error) {
	return this.shelfClient.ReadOne(item)
}

func (this *ShelfService) ReadAll(search structs.Search) (structs.Result, error) {
	return this.shelfClient.ReadAll(search)
}
