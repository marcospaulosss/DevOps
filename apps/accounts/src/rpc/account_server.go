package rpc

import (
	"context"
	"math/rand"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"

	"backend/apps/accounts/libs/notification"
	"backend/apps/accounts/src/adapters"
	"backend/apps/accounts/src/repositories"
	"backend/apps/accounts/src/structs"
	log "backend/libs/logger"
	pb "backend/proto"
)

type AccountServer struct {
	accountRepository repositories.AccountRepository
	userRepository    repositories.UserRepository
	notification      notification.NotificationInterface
}

var PermissionDeniedMsg = "Usuario ja cadastrado na base."

func NewAccountServer(s *grpc.Server, repos repositories.Container) *AccountServer {
	accountRepository := repos.AccountRepository.(repositories.AccountRepository)
	userRepository := repos.UserRepository.(repositories.UserRepository)
	notification := repos.Notifications
	server := &AccountServer{accountRepository, userRepository, notification}
	if s != nil {
		pb.RegisterAccountServiceServer(s, server)
	}
	return server
}

func (this *AccountServer) Create(ctx context.Context, req *pb.AccountRequest) (*pb.AccountResponse, error) {
	log.SetRequestID(req.GetId())

	log.Info("Vou converter a request para account", req)
	account := adapters.ToDomainAccount(req.GetAccount())

	code := this.getCode()

	if account.Type == "email" {
		account.EmailCode = code
	} else {
		account.PhoneCode = code
	}

	if account.Exists {
		return this.login(&account, code)
	}

	return this.register(&account)
}

func (this *AccountServer) ReadOne(ctx context.Context, req *pb.AccountRequest) (*pb.AccountResponse, error) {
	log.SetRequestID(req.GetId())

	log.Info("Vou converter a request para account", req)
	account := adapters.ToDomainAccount(req.GetAccount())

	found, err := this.accountRepository.ReadOne(&account)
	if err != nil {
		log.Error("Falhei na validacao: ", err.Error())
		return nil, err
	}

	accountResult := found.(*structs.Account)

	if account.Exists && !accountResult.Exists {
		log.Info("Usuario ja cadastrado na base de dados: ", account)
		return nil, status.Errorf(codes.AlreadyExists, PermissionDeniedMsg)
	}

	if account.Exists {
		_, err = this.accountRepository.Delete(&account)
		if err != nil {
			log.Error("Falhei ao deletar account: ", err.Error())
			return nil, err
		}
	}

	log.Info("Vou Criar o response de account id = ", accountResult.ID)
	accountPb := adapters.ToProtoAccount(*accountResult)

	resp := &pb.AccountResponse{Account: accountPb}

	return resp, nil
}

func (this *AccountServer) getCode() string {
	log.Info("Vou gerar codigo para envio de email e sms...")

	code := (rand.Intn(999999-100000) + 100000)

	log.Info("Criei codigo com sucesso.", code)

	return strconv.Itoa(code)
}

func (this *AccountServer) login(account *structs.Account, code string) (*pb.AccountResponse, error) {
	searchTerm := "(email[eq]:'" + account.Email + "')"
	if account.Type == "phone" {
		searchTerm = "(phone[eq]:'" + account.Phone + "')"
	}
	search := structs.Search{
		Pagination: structs.Pagination{PerPage: 1, Page: 0},
		Raw:        searchTerm,
	}

	log.Info("Vou verificar a existencia do usuario...")
	createdUser, err := this.userRepository.ReadAll(search)
	if err != nil {
		log.Error("Falhei na busca do usuario: ", err.Error())
		return nil, err
	}
	user := createdUser.Items.([]structs.User)

	if len(user) == 0 {
		log.Info("Nao foi possivel encontrar o usuario.")
		return nil, status.Error(codes.NotFound, "User not found")
	}

	var saved interface{}
	saved, err = this.accountRepository.FetchCodeExisting(account)
	if err != nil {
		saved, err = this.accountRepository.Create(account)
		if err != nil {
			log.Error("Falhei na criacao do codigo: ", err.Error())
			return nil, err
		}
	}

	if account.Type == "email" {
		this.notification.SendEmailNotification(user[0].Email, account.EmailCode)
		this.notification.SendSmsNotification(user[0].Phone, account.EmailCode)
	} else {
		this.notification.SendEmailNotification(user[0].Email, account.PhoneCode)
		this.notification.SendSmsNotification(user[0].Phone, account.PhoneCode)
	}

	accountResult := adapters.ToProtoAccount(*saved.(*structs.Account))

	resp := &pb.AccountResponse{
		Account: accountResult,
	}

	return resp, nil
}

func (this *AccountServer) register(account *structs.Account) (*pb.AccountResponse, error) {
	log.Info("Vou verificar a existencia de um codigo para esse usuario.", account.Email, account.Phone)
	saved, err := this.accountRepository.FetchCodeExisting(account)
	if err == nil {
		if saved.Exists && !account.Exists {
			return nil, status.Errorf(codes.PermissionDenied, PermissionDeniedMsg)
		}
	} else {
		var created interface{}
		created, err = this.accountRepository.Create(account)
		if err != nil {
			log.Error("Falhei na criacao: ", err.Error())
			return nil, status.Errorf(codes.PermissionDenied, PermissionDeniedMsg)
		}

		sub := created.(*structs.Account)
		saved = sub
	}

	//account := created.(*structs.account)
	if account.Type == "email" {
		this.notification.SendEmailNotification(account.Email, account.EmailCode)
	} else {
		this.notification.SendSmsNotification(account.Phone, account.PhoneCode)
	}

	log.Info("Vou Criar o response de account id = ", saved.ID)
	accountPb := adapters.ToProtoAccount(*saved)

	resp := &pb.AccountResponse{Account: accountPb}

	return resp, nil
}
