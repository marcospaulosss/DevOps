package rpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"backend/apps/accounts/src/adapters"
	"backend/apps/accounts/src/repositories"
	"backend/apps/accounts/src/structs"
	log "backend/libs/logger"
	pb "backend/proto"
)

type UserServer struct {
	userRepository    repositories.Repository
	accountRepository repositories.Repository
}

func NewUserServer(s *grpc.Server, repos repositories.Container) *UserServer {
	accountRepository := repos.AccountRepository.(repositories.AccountRepository)
	userRepository := repos.UserRepository.(repositories.UserRepository)
	server := &UserServer{userRepository, accountRepository}
	if s != nil {
		pb.RegisterUserServiceServer(s, server)
	}
	return server
}

func (this *UserServer) Create(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	log.Info("Irei validar o telefone/email do usuario", fmt.Sprintf("%+v", in))

	var err error
	foundPhoneAccount, errPhone := this.accountRepository.ReadOne(&structs.Account{
		Phone:     in.GetUser().GetPhone(),
		PhoneCode: in.GetUser().GetPhoneCode(),
		Type:      "phone",
	})
	if errPhone != nil {
		log.Error(fmt.Sprintf("Erro na validacao do telefone: %v", errPhone))
		return nil, errPhone
	}

	foundEmailAccount, errEmail := this.accountRepository.ReadOne(&structs.Account{
		Email:     in.GetUser().GetEmail(),
		EmailCode: in.GetUser().GetEmailCode(),
		Type:      "email",
	})

	if errEmail != nil {
		log.Error(fmt.Sprintf("Erro na validacao do e-mail: %v", errEmail))
		return nil, errEmail
	}

	emailValidation := foundPhoneAccount.(*structs.Account)
	phoneValidation := foundEmailAccount.(*structs.Account)

	if phoneValidation.Exists || emailValidation.Exists {
		log.Info("Usuario ja cadastrado na base de dados",
			fmt.Sprintf("Telefone existe? %v Email existe? %v",
				phoneValidation.Exists, emailValidation.Exists))

		return nil, status.Errorf(codes.AlreadyExists, "O usuario ja existe")
	}

	log.Info("Preparando dados recebidos para inserir usuario", fmt.Sprintf("%+v", in))

	created, err := this.userRepository.Create(in.GetUser())
	if err != nil {
		log.Error("Falhei na criacao:", err.Error())
		return nil, err
	}

	_, err = this.accountRepository.Delete(&structs.Account{
		Phone: in.GetUser().GetPhone(),
		Email: in.GetUser().GetEmail(),
	})
	if err != nil {
		log.Error("Falhei ao deletar codigos de validacao da base de dados:", err.Error())
		return nil, err
	}

	user := created.(*structs.User)
	resp := &pb.UserResponse{
		User: &pb.User{
			Id: user.ID,
		},
	}

	log.Info("Retornando usuario recem criado para o severino", fmt.Sprintf("%+v", resp))
	return resp, nil
}

func (this *UserServer) ReadAll(ctx context.Context, in *pb.SearchRequest) (*pb.UsersResponse, error) {
	log.Info("Preparando dados recebidos para consultar usuarios", fmt.Sprintf("%+v", in))

	users, err := this.userRepository.ReadAll(adapters.ToDomainSearch(in.GetSearch()))
	if err != nil {
		log.Error("Falhei no ReadAll do repository: ", err.Error())
		return nil, err
	}

	usersPb := []*pb.User{}
	usersList := users.Items.([]structs.User)
	for key, _ := range usersList {

		updated_at := ""
		if usersList[key].UpdatedAt != nil {
			updated_at = *usersList[key].UpdatedAt
		}

		usersPb = append(usersPb, &pb.User{
			Id:        usersList[key].ID,
			Name:      usersList[key].Name,
			Phone:     usersList[key].Phone,
			Email:     usersList[key].Email,
			Active:    usersList[key].Active,
			CreatedAt: usersList[key].CreatedAt,
			UpdatedAt: updated_at,
		})
	}

	log.Info("Retornei", users.Pagination.Total, "usuarios para o severino:", fmt.Sprintf("%+v", usersPb))

	return &pb.UsersResponse{
		Users: usersPb,
		Total: users.Pagination.Total,
	}, nil
}
