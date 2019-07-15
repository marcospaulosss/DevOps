package rpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"backend/apps/severino/src/adapters"
	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/structs"
	log "backend/libs/logger"
	pb "backend/proto"
)

type UserClientImpl struct {
	service pb.UserServiceClient
	ctx     context.Context
}

func NewUserClient(conn *grpc.ClientConn, ctx context.Context) interfaces.Client {
	return &UserClientImpl{
		service: pb.NewUserServiceClient(conn),
		ctx:     ctx,
	}
}

func (this *UserClientImpl) Create(user interface{}) (interface{}, error) {
	log.Info("Criando request...")
	req := adapters.CreateUserRequest(user.(*structs.User))

	log.Info("Enviando o request: ", fmt.Sprintf("%+v", req))
	resp, err := this.service.Create(this.ctx, req)

	if err != nil {
		log.Error("Falhei. Erro na resposta do servico.", err.Error())
		return nil, err
	}

	log.Info("Preparando a resposta para retorno: ", resp)
	response := adapters.ToDomainUser(resp.User)
	return response, err
}

func (this *UserClientImpl) ReadOne(params interface{}) (interface{}, error) {
	return nil, nil
}

func (this *UserClientImpl) Update(album interface{}) (interface{}, error) {
	return nil, nil
}

func (this *UserClientImpl) ReadAll(search structs.Search) (structs.Result, error) {
	log.Info("Vou criar um search request contendo:", search.String())
	req := adapters.CreateSearchRequest(search)
	log.Info("Enviando o request para o serviço de accounts", req)
	resp, err := this.service.ReadAll(this.ctx, req)
	if err != nil {
		log.Error("Falhei. Erro na resposta do servidor: ", err.Error())
		return structs.Result{}, err
	}

	items := adapters.ToDomainUsers(resp.Users)
	pagination := search.Pagination
	pagination.Total = int(resp.Total)

	result := structs.Result{Items: items, Pagination: pagination}
	return result, nil
}

func (this *UserClientImpl) Delete(album interface{}) (interface{}, error) {
	return nil, nil
}
