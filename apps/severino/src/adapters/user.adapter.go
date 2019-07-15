package adapters

import (
	"backend/apps/severino/src/structs"
	log "backend/libs/logger"
	pb "backend/proto"
)

func CreateUserRequest(in *structs.User) *pb.UserRequest {
	log.Info("Vou criar o request o usuario com email", in.Email, "e telefone", in.Phone)
	req := &pb.UserRequest{
		User: ToProtoUser(*in),
		Id:   log.GetRequestID(),
	}
	log.Info("Criei. Vou enviar o request:", req.String())
	return req
}

func ToProtoUser(in structs.User) *pb.User {
	return &pb.User{
		Name:      in.Name,
		Phone:     "55" + in.Phone,
		PhoneCode: in.PhoneCode,
		Email:     in.Email,
		EmailCode: in.EmailCode,
		Active:    in.Active,
	}
}

func ToDomainUser(in *pb.User) structs.User {
	user := structs.User{}
	if in != nil {
		user.ID = in.GetId()
		user.Name = in.GetName()
		user.Phone = in.GetPhone()
		user.PhoneCode = in.GetPhoneCode()
		user.Email = in.GetEmail()
		user.EmailCode = in.GetEmailCode()
		user.Active = in.GetActive()
		user.CreatedAt = in.GetCreatedAt()
		user.UpdatedAt = in.GetUpdatedAt()
	}
	return user
}

func ToDomainUsers(u []*pb.User) []structs.User {
	users := []structs.User{}
	if u == nil {
		return users
	}
	for _, user := range u {
		users = append(users, ToDomainUser(user))
	}
	return users
}
