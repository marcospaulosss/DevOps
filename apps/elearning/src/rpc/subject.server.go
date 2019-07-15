package rpc

import (
	"context"

	"google.golang.org/grpc"

	"backend/apps/elearning/src/adapters"
	"backend/apps/elearning/src/repositories"
	"backend/apps/elearning/src/structs"
	log "backend/libs/logger"
	pb "backend/proto"
)

type SubjectServer struct {
	subjectRepository repositories.SubjectRepository
}

func NewSubjectServer(s *grpc.Server, repos repositories.Container) *SubjectServer {
	subjectRepository := repos.SubjectRepository.(repositories.SubjectRepository)
	server := &SubjectServer{subjectRepository: subjectRepository}
	if s != nil {
		pb.RegisterSubjectServiceServer(s, server)
	}
	return server
}

const LOGCONVERT = "Recebi o request e vou converter pb.Subject para structs.Subject."
const LOGERROR = "Tudo errado. Vou retornar o response com erro."
const LOGSUCCESS = "Tudo certo. Vou retornar o response contendo o item com ID:"

func (this SubjectServer) Create(ctx context.Context, req *pb.SubjectRequest) (*pb.SubjectResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetSubject()
	log.Info(LOGCONVERT)
	subject := adapters.ToDomainSubject(item)
	result, err := this.subjectRepository.Create(subject)
	if err != nil {
		log.Error(LOGERROR, err)
		return &pb.SubjectResponse{}, err
	}

	saved := result.(structs.Subject)
	log.Info("Vou converter para pb.Subject e criar o response.")
	resp := &pb.SubjectResponse{
		Subject: adapters.ToProtoSubject(saved),
	}
	log.Info(LOGSUCCESS, saved.ID)
	return resp, nil
}

func (this SubjectServer) Update(ctx context.Context, req *pb.SubjectRequest) (*pb.SubjectResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetSubject()
	log.Info(LOGCONVERT)
	subject := adapters.ToDomainSubject(item)
	result, err := this.subjectRepository.Update(subject)
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com erro.", err)
		return &pb.SubjectResponse{}, err
	}

	saved := result.(structs.Subject)
	log.Info("Vou converter para pb.Subject e criar o response.")
	resp := &pb.SubjectResponse{
		Subject: adapters.ToProtoSubject(saved),
	}
	log.Info(LOGSUCCESS, saved.ID)
	return resp, nil
}

func (this *SubjectServer) ReadOne(ctx context.Context, req *pb.SubjectRequest) (*pb.SubjectResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetSubject()
	log.Info(LOGCONVERT)
	found, err := this.subjectRepository.ReadOne(adapters.ToDomainSubject(item))
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com o erro.", err)
		return &pb.SubjectResponse{}, err
	}

	subject := found.(structs.Subject)
	resp := &pb.SubjectResponse{
		Subject: adapters.ToProtoSubject(subject),
	}
	log.Info("Tudo certo. Vou retornar o response contendo o subject:", item)
	return resp, nil
}

func (this *SubjectServer) ReadAll(ctx context.Context, req *pb.SearchRequest) (*pb.SubjectsResponse, error) {
	log.SetRequestID(req.GetId())
	search := adapters.ToDomainSearch(req.GetSearch())
	log.Info(LOGCONVERT, search.String())
	result, err := this.subjectRepository.ReadAll(search)
	if err != nil {
		log.Error(LOGERROR, err)
		return &pb.SubjectsResponse{}, err
	}
	subjects := result.Items.([]structs.Subject)

	resp := &pb.SubjectsResponse{
		Subjects:   adapters.ToProtoSubjects(subjects),
		Pagination: adapters.ToProtoPagination(result.Pagination),
	}
	log.Info("Tudo certo. Vou retornar o response contendo o subject:", len(subjects), "subjects.")
	return resp, nil
}

func (this *SubjectServer) Delete(ctx context.Context, req *pb.SubjectRequest) (*pb.SubjectResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetSubject()
	log.Info(LOGCONVERT)
	subject := adapters.ToDomainSubject(item)
	_, err := this.subjectRepository.Delete(subject)
	if err != nil {
		log.Error(LOGERROR, err)
		return &pb.SubjectResponse{}, err
	}

	resp := &pb.SubjectResponse{
		Subject: item,
	}
	log.Info(LOGSUCCESS, item.Id)
	return resp, nil
}
