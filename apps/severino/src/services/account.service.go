package services

import (
	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/src/structs"
)

type AccountService struct {
	accountClient interfaces.Client
}

type AccountResult struct {
	Data interface{}
	Err  error
}

func NewAccountService(container rpc.Container) *AccountService {
	return &AccountService{
		accountClient: container.Accounts.AccountClient,
	}
}

func (this *AccountService) Create(item interface{}) (interface{}, error) {
	return this.accountClient.Create(item)
}

func (this *AccountService) ReadOne(item interface{}) (interface{}, error) {
	return this.accountClient.ReadOne(item)
}

func (this *AccountService) Update(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this *AccountService) Delete(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this *AccountService) ReadAll(search structs.Search) (structs.Result, error) {
	return structs.Result{}, nil
}

//func (this *AccountService) validateParamsEmpty(params *structs.account) error {
//	if params.Type == "email" && (strings.Contains(params.Email, " ") || params.Email == "") {
//		return status.Errorf(codes.InvalidArgument, "email incorrect")
//	}
//
//	if params.Type == "phone" && params.Phone == "" {
//		return status.Errorf(codes.InvalidArgument, "phone incorrect")
//	}
//
//	return nil
//}
