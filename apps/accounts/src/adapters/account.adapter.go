package adapters

import (
	"backend/apps/accounts/src/structs"
	pb "backend/proto"
)

func ToProtoAccount(in structs.Account) *pb.Account {
	return &pb.Account{
		Id:        in.ID,
		Email:     in.Email,
		EmailCode: in.EmailCode,
		Phone:     in.Phone,
		PhoneCode: in.PhoneCode,
		CreatedAt: in.CreatedAt,
		Exists:    in.Exists,
		UserId:    in.UserID,
	}
}

func ToProtoAccounts(list []structs.Account) []*pb.Account {
	var items []*pb.Account
	for _, item := range list {
		items = append(items, ToProtoAccount(item))
	}
	return items
}

func ToDomainAccounts(list []*pb.Account) []structs.Account {
	var items []structs.Account
	for _, item := range list {
		items = append(items, ToDomainAccount(item))
	}
	return items
}

func ToDomainAccount(in *pb.Account) structs.Account {
	return structs.Account{
		ID:        in.GetId(),
		Email:     in.GetEmail(),
		EmailCode: in.GetEmailCode(),
		Phone:     in.GetPhone(),
		PhoneCode: in.GetPhoneCode(),
		CreatedAt: in.GetCreatedAt(),
		Type:      in.GetType(),
		Exists:    in.Exists,
	}
}
