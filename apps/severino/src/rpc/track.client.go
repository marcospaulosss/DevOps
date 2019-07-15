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

type TrackClientImpl struct {
	service pb.TrackServiceClient
	ctx     context.Context
}

func NewTrackClient(service pb.TrackServiceClient, ctx context.Context) interfaces.Client {
	return &TrackClientImpl{service, ctx}

}

func (this *TrackClientImpl) Create(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.Create(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this *TrackClientImpl) Update(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.Update(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this *TrackClientImpl) Delete(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.Delete(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this *TrackClientImpl) ReadOne(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.ReadOne(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return nil, errors.NewErrorFrom(err)
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this *TrackClientImpl) ReadAll(search structs.Search) (structs.Result, error) {
	log.Info("Vou criar um search request contendo:", search.String())
	req := adapters.CreateSearchRequest(search)
	log.Info("Vou enviar o request:", req.String())
	res, err := this.service.ReadAll(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return structs.Result{}, errors.NewErrorFrom(err)
	}

	items := adapters.ToDomainTracks(res.GetTracks())
	pagination := adapters.ToDomainPagination(res.GetPagination())
	result := structs.Result{items, pagination}
	return result, nil
}

func (this *TrackClientImpl) createRequest(item interface{}) *pb.TrackRequest {
	track := item.(structs.Track)
	return adapters.CreateTrackRequest(track)
}

func (this *TrackClientImpl) toDomain(res *pb.TrackResponse) structs.Track {
	log.Info("Sucesso. Recebi o response:", res.String())
	saved := adapters.ToDomainTrack(res.GetTrack())
	log.Info("Vou retornar o item com ID:", saved.ID)
	return saved
}
