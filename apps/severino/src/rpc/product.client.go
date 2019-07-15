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

type ProductClientImpl struct {
	service pb.ProductServiceClient
	ctx     context.Context
}

func NewProductClient(service pb.ProductServiceClient, ctx context.Context) interfaces.Client {
	return &ProductClientImpl{service, ctx}
}

func (this *ProductClientImpl) Create(item interface{}) (interface{}, error) {
	product := item.(structs.Product)
	log.Info("Vou criar o request para inserir o item com name:", product.Name)
	req := this.createRequest(product)
	log.Info("Criei. Vou enviar o request:", req.String())
	resp, err := this.service.Create(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	log.Info("Sucesso. Recebi o response:", resp.String())
	saved := adapters.ToDomainProduct(resp.GetProduct())
	log.Info("Foi criado com ID:", saved.ID)
	return saved, nil
}

func (this *ProductClientImpl) Update(item interface{}) (interface{}, error) {
	product := item.(structs.Product)
	log.Info("Vou criar o request para atualizar o item com name:", product.Name)
	req := this.createRequest(product)
	log.Info("Criei. Vou enviar o request:", req.String())
	resp, err := this.service.Update(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	log.Info("Sucesso. Recebi o response:", resp.String())
	saved := adapters.ToDomainProduct(resp.GetProduct())
	log.Info("Foi atualizado com ID:", saved.ID)
	return saved, nil
}

func (this *ProductClientImpl) Delete(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this *ProductClientImpl) ReadOne(item interface{}) (interface{}, error) {
	product := item.(structs.Product)
	log.Info("Vou criar o request para buscar o item com ID:", product.ID)
	req := this.createRequest(product)
	log.Info("Criei. Vou enviar o request:", req.String())
	resp, err := this.service.ReadOne(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	log.Info("Sucesso. Recebi o response:", resp.String())
	saved := adapters.ToDomainProduct(resp.GetProduct())
	log.Info("Foi encontrado com ID:", saved.ID)
	return saved, nil
}

func (this *ProductClientImpl) ReadAll(search structs.Search) (structs.Result, error) {
	log.Info("Vou criar um search request contendo:", search.String())
	req := adapters.CreateSearchRequest(search)
	log.Info("Vou enviar o request:", req.String())
	resp, err := this.service.ReadAll(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return structs.Result{}, errors.NewErrorFrom(err)
	}

	items := adapters.ToDomainProducts(resp.GetProducts())
	log.Info("Sucesso. Total de itens:", len(items))
	pagination := adapters.ToDomainPagination(resp.GetPagination())
	result := structs.Result{Items: items, Pagination: pagination}
	return result, nil
}

func (this *ProductClientImpl) createRequest(item structs.Product) *pb.ProductRequest {
	return &pb.ProductRequest{
		Product: adapters.ToProtoProduct(item),
	}
}
