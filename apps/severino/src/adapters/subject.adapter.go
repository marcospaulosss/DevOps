package adapters

import (
	"backend/apps/severino/src/structs"
	log "backend/libs/logger"
	pb "backend/proto"
)

func CreateSubjectRequest(item structs.Subject) *pb.SubjectRequest {
	log.Info("Vou criar o request para subject com com ID:", item.ID)
	req := &pb.SubjectRequest{
		Subject: ToProtoSubject(item),
		Id:      log.GetRequestID(),
	}
	log.Info("Criei. Vou enviar o request:", req.String())
	return req
}

func ToProtoSubjects(subjects []structs.Subject) []*pb.Subject {
	result := []*pb.Subject{}
	for _, item := range subjects {
		subject := ToProtoSubject(item)
		result = append(result, subject)
	}
	return result
}

func ToProtoSubject(subject structs.Subject) *pb.Subject {
	return &pb.Subject{
		Id:    subject.ID,
		Title: subject.Title,
	}
}

func ToDomainSubjects(t []*pb.Subject) []structs.Subject {
	subjects := []structs.Subject{}
	if t == nil {
		return subjects
	}
	for _, subject := range t {
		subjects = append(subjects, ToDomainSubject(subject))
	}
	return subjects
}

func ToDomainSubject(t *pb.Subject) structs.Subject {
	subject := structs.Subject{}
	if t != nil {
		subject.ID = t.GetId()
		subject.Title = t.GetTitle()
		subject.CreatedAt = t.GetCreatedAt()
		subject.UpdatedAt = t.GetUpdatedAt()
	}
	return subject
}
