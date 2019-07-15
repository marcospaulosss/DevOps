package rpc

import (
	"context"

	"github.com/labstack/gommon/log"

	"backend/apps/severino/src/adapters"
	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/structs"
	"backend/libs/errors"
	pb "backend/proto"
)

type HomeClientImpl struct {
	service pb.HomeServiceClient
	ctx     context.Context
}

func NewHomeClient(service pb.HomeServiceClient, ctx context.Context) interfaces.Client {
	return &HomeClientImpl{service, ctx}
}

func (this *HomeClientImpl) Create(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this *HomeClientImpl) Update(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this *HomeClientImpl) ReadOne(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this *HomeClientImpl) ReadAll(search structs.Search) (structs.Result, error) {
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

func (this *HomeClientImpl) Delete(item interface{}) (interface{}, error) {
	return nil, nil
}
