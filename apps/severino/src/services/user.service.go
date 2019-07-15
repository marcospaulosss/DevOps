package services

import (
	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/src/structs"
)

type UserService struct {
	userClient interfaces.Client
}

func NewUserService(container rpc.Container) *UserService {
	return &UserService{
		userClient: container.Accounts.UserClient,
	}
}

func (this *UserService) Create(item interface{}) (interface{}, error) {
	return this.userClient.Create(item)
}

func (this *UserService) Update(item interface{}) (interface{}, error) {
	return this.userClient.Update(item)
}

func (this *UserService) Delete(item interface{}) (interface{}, error) {
	return this.userClient.Delete(item)
}

func (this *UserService) ReadOne(item interface{}) (interface{}, error) {
	return this.userClient.ReadOne(item)
}

func (this *UserService) ReadAll(search structs.Search) (structs.Result, error) {
	return this.userClient.ReadAll(search)
}
