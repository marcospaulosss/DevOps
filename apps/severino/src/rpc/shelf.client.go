package rpc

import (
	"context"

	"backend/apps/severino/src/adapters"
	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/structs"
	"backend/libs/errors"
	log "backend/libs/logger"
	pb "backend/proto"
)

type ShelfClientImpl struct {
	service pb.ShelfServiceClient
	ctx     context.Context
}

func NewShelfClient(service pb.ShelfServiceClient, ctx context.Context) interfaces.Client {
	return &ShelfClientImpl{service, ctx}
}

func (this *ShelfClientImpl) Create(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.Create(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this *ShelfClientImpl) Update(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.Update(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this *ShelfClientImpl) Delete(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.Delete(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this *ShelfClientImpl) ReadOne(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.ReadOne(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this *ShelfClientImpl) ReadAll(search structs.Search) (structs.Result, error) {
	log.Info("Vou criar um search request contendo:", search.String())
	req := adapters.CreateSearchRequest(search)
	log.Info("Vou enviar o request:", req.String())
	res, err := this.service.ReadAll(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return structs.Result{}, errors.NewErrorFrom(err)
	}

	items := adapters.ToDomainShelves(res.GetShelves())
	pagination := adapters.ToDomainPagination(res.GetPagination())
	result := structs.Result{items, pagination}
	return result, nil
}

func (this *ShelfClientImpl) createRequest(item interface{}) *pb.ShelfRequest {
	shelf := item.(structs.Shelf)
	return adapters.CreateShelfRequest(shelf)
}

func (this *ShelfClientImpl) toDomain(res *pb.ShelfResponse) structs.Shelf {
	log.Info("Sucesso. Recebi o response:", res.String())
	saved := adapters.ToDomainShelf(res.GetShelf())
	log.Info("Vou retornar o item com ID:", saved.ID)
	return saved
}
