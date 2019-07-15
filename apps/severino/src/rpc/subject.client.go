package rpc

import (
	"context"

	"backend/apps/severino/src/adapters"
	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/structs"
	log "backend/libs/logger"
	pb "backend/proto"
)

type SubjectClientImpl struct {
	service pb.SubjectServiceClient
	ctx     context.Context
}

func NewSubjectClient(service pb.SubjectServiceClient, ctx context.Context) interfaces.Client {
	return &SubjectClientImpl{service, ctx}
}

const LOGFAILED = "Falhei."

func (this *SubjectClientImpl) Create(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.Create(this.ctx, req)
	if err != nil {
		log.Error(LOGFAILED, err)
		return nil, err
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this *SubjectClientImpl) ReadAll(search structs.Search) (structs.Result, error) {
	log.Info("Vou criar um search request contendo:", search.String())
	req := adapters.CreateSearchRequest(search)
	log.Info("Vou enviar o request:", req.String())
	res, err := this.service.ReadAll(this.ctx, req)
	if err != nil {
		log.Error("Falhei.", err)
		return structs.Result{}, err
	}

	items := adapters.ToDomainSubjects(res.GetSubjects())
	pagination := adapters.ToDomainPagination(res.GetPagination())
	result := structs.Result{items, pagination}
	return result, nil
}

func (this *SubjectClientImpl) ReadOne(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.ReadOne(this.ctx, req)
	if err != nil {
		log.Error(LOGFAILED, err)
		return nil, err
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this *SubjectClientImpl) Update(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.Update(this.ctx, req)
	if err != nil {
		log.Error(LOGFAILED, err)
		return nil, err
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this *SubjectClientImpl) Delete(item interface{}) (interface{}, error) {
	req := this.createRequest(item)
	res, err := this.service.Delete(this.ctx, req)
	if err != nil {
		log.Error(LOGFAILED, err)
		return nil, err
	}
	saved := this.toDomain(res)
	return saved, nil
}

func (this *SubjectClientImpl) createRequest(item interface{}) *pb.SubjectRequest {
	subject := item.(structs.Subject)
	return adapters.CreateSubjectRequest(subject)
}

func (this *SubjectClientImpl) toDomain(res *pb.SubjectResponse) structs.Subject {
	log.Info("Sucesso. Recebi o response:", res.String())
	saved := adapters.ToDomainSubject(res.GetSubject())
	log.Info("Vou retornar o item com ID:", saved.ID)
	return saved
}
