package rpc

import (
	"context"

	"backend/apps/severino/src/adapters"

	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/structs"
	log "backend/libs/logger"
	pb "backend/proto"
)

type AccountClientImpl struct {
	service pb.AccountServiceClient
	ctx     context.Context
}

func NewAccountClient(service pb.AccountServiceClient, ctx context.Context) interfaces.Client {
	return &AccountClientImpl{service, ctx}
}

func (this *AccountClientImpl) Create(item interface{}) (interface{}, error) {
	req := this.createRequest(item)

	log.Info("Criei. Vou enviar o request:", req.String())
	created, err := this.service.Create(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, err
	}

	return this.toDomain(created), err
}

func (this *AccountClientImpl) ReadOne(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	found, err := this.service.ReadOne(this.ctx, req)
	if err != nil {
		log.Error("Falhei. Erro na resposta do servico.", err.Error())
		return nil, err
	}

	return this.toDomain(found), err
}

func (this *AccountClientImpl) Update(album interface{}) (interface{}, error) {
	return nil, nil
}

func (this *AccountClientImpl) ReadAll(search structs.Search) (structs.Result, error) {
	return structs.Result{}, nil
}

func (this *AccountClientImpl) Delete(album interface{}) (interface{}, error) {
	return nil, nil
}

func (this *AccountClientImpl) createRequest(item interface{}) *pb.AccountRequest {
	account := item.(*structs.Account)
	return adapters.CreateAccountRequest(*account)
}

func (this *AccountClientImpl) toDomain(res *pb.AccountResponse) structs.Account {
	log.Info("Sucesso. Recebi o response:", res.String())
	saved := adapters.ToDomainAccount(res.GetAccount())
	log.Info("Vou retornar o item com ID:", saved.ID)
	return saved
}
