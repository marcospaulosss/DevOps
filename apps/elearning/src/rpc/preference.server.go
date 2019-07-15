package rpc

import (
	"context"

	"google.golang.org/grpc"

	"backend/apps/elearning/src/adapters"
	"backend/apps/elearning/src/repositories"
	log "backend/libs/logger"
	pb "backend/proto"
)

type PreferenceServer struct {
	preferenceRepository repositories.PreferenceRepository
}

func NewPreferenceServer(s *grpc.Server, repos repositories.Container) *PreferenceServer {
	preferenceRepository := repos.PreferenceRepository.(repositories.PreferenceRepository)
	server := &PreferenceServer{preferenceRepository}
	if s != nil {
		pb.RegisterPreferenceServiceServer(s, server)
	}
	return server
}

func (this *PreferenceServer) Update(ctx context.Context, req *pb.PreferenceRequest) (*pb.PreferenceResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetPreference()
	log.Info("Vou atualizar a preference do tipo:", item.Type)
	_, err := this.preferenceRepository.Update(adapters.ToDomainPreference(item))
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com o erro.", err)
		return &pb.PreferenceResponse{}, err
	}
	log.Info("Ok. Salvei a preference com conteudo:", item.Content)
	resp := &pb.PreferenceResponse{
		Preference: item,
	}
	log.Info("Tudo certo. Vou retornar o response contendo a preference salva.")
	return resp, nil
}
