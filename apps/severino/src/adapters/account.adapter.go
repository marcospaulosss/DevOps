package adapters

import (
	"fmt"

	"backend/apps/severino/src/structs"
	log "backend/libs/logger"
	pb "backend/proto"
)

func CreateAccountRequest(item structs.Account) *pb.AccountRequest {
	log.Info("Vou criar o request para account:", item)

	req := &pb.AccountRequest{
		Account: ToProtoAccount(item),
		Id:      log.GetRequestID(),
	}

	return req
}

func ToProtoAccount(account structs.Account) *pb.Account {
	var req = pb.Account{
		Id:        account.ID,
		Type:      account.Type,
		CreatedAt: account.CreatedAt,
		Exists:    account.Exists,
	}

	if account.Type == "email" {
		req.Email = account.Email
		req.EmailCode = account.Code
	} else {
		req.Phone = fmt.Sprintf("55%s", account.Phone)
		req.PhoneCode = account.Code
	}

	return &req
}

func ToDomainAccounts(s []*pb.Account) []structs.Account {
	accounts := []structs.Account{}
	if s == nil {
		return accounts
	}
	for _, account := range s {
		accounts = append(accounts, ToDomainAccount(account))
	}
	return accounts
}

func ToDomainAccount(s *pb.Account) structs.Account {
	account := structs.Account{}
	if s != nil {
		account.ID = s.GetId()
		account.Type = s.GetType()

		if account.Type == "phone" {
			account.Phone = s.GetPhone()
			account.Code = s.GetPhoneCode()
		} else {
			account.Code = s.GetEmailCode()
			account.Email = s.GetEmail()
		}

		account.Exists = s.Exists
		account.CreatedAt = s.GetCreatedAt()
		account.UserID = s.GetUserId()
	}
	return account
}
