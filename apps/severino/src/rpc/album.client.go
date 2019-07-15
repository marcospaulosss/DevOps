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

type AlbumClientImpl struct {
	service pb.AlbumServiceClient
	ctx     context.Context
}

func NewAlbumClient(service pb.AlbumServiceClient, ctx context.Context) interfaces.AlbumClient {
	return AlbumClientImpl{service, ctx}
}

func (this AlbumClientImpl) Create(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.Create(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this AlbumClientImpl) Update(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.Update(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this AlbumClientImpl) Delete(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.Delete(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this AlbumClientImpl) ReadOne(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.ReadOne(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this AlbumClientImpl) ReadAll(search structs.Search) (structs.Result, error) {
	log.Info("Vou criar um search request contendo:", search.String())
	req := adapters.CreateSearchRequest(search)
	log.Info("Vou enviar o request:", req.String())
	res, err := this.service.ReadAll(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return structs.Result{}, errors.NewErrorFrom(err)
	}

	items := adapters.ToDomainAlbums(res.GetAlbums())
	pagination := adapters.ToDomainPagination(res.GetPagination())
	result := structs.Result{items, pagination}
	return result, nil
}

func (this AlbumClientImpl) Publish(album structs.Album) (structs.Album, error) {
	req := adapters.CreateAlbumRequest(album)
	_, err := this.service.Publish(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
	}
	log.Info("Sucesso. Publiquei o album.")
	return album, err
}

func (this AlbumClientImpl) Unpublish(album structs.Album) (structs.Album, error) {
	req := adapters.CreateAlbumRequest(album)
	_, err := this.service.Unpublish(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
	}
	log.Info("Sucesso. Despubliquei o album.")
	return album, err
}
func (this AlbumClientImpl) createRequest(item interface{}) *pb.AlbumRequest {
	album := item.(structs.Album)
	return adapters.CreateAlbumRequest(album)
}

func (this AlbumClientImpl) toDomain(res *pb.AlbumResponse) structs.Album {
	log.Info("Sucesso. Recebi o response:", res.String())
	saved := adapters.ToDomainAlbum(res.GetAlbum())
	log.Info("Vou retornar o item com ID:", saved.ID)
	return saved
}
