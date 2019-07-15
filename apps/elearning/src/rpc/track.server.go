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

type TrackServer struct {
	trackRepository repositories.TrackRepository
	albumRepository repositories.AlbumRepository
}

func NewTrackServer(s *grpc.Server, repos repositories.Container) *TrackServer {
	trackRepository := repos.TrackRepository.(repositories.TrackRepository)
	albumRepository := repos.AlbumRepository.(repositories.AlbumRepository)
	server := &TrackServer{trackRepository, albumRepository}
	if s != nil {
		pb.RegisterTrackServiceServer(s, server)
	}
	return server
}

func (this *TrackServer) Create(ctx context.Context, req *pb.TrackRequest) (*pb.TrackResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetTrack()
	log.Info("Recebi o request e vou converter pb.Track para structs.Track.")
	track := adapters.ToDomainTrack(item)
	result, err := this.trackRepository.Create(track)
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com erro.", err)
		return &pb.TrackResponse{}, err
	}

	saved := result.(structs.Track)
	track.ID = saved.ID
	if err = this.trackRepository.CreateSubjectsAndAssociateItWithTracks(track); err != nil {
		log.Error("Erro na associacao. Vou retornar o response com erro.", err)
		return &pb.TrackResponse{}, err
	}

	log.Info("Vou converter para pb.Track e criar o response.")
	resp := &pb.TrackResponse{
		Track: adapters.ToProtoTrack(saved),
	}
	log.Info("Tudo certo. Vou retornar o response contendo o item com ID:", saved.ID)
	return resp, nil
}

func (this *TrackServer) Update(ctx context.Context, req *pb.TrackRequest) (*pb.TrackResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetTrack()
	log.Info("Recebi o request e vou converter pb.Track para structs.Track.")
	track := adapters.ToDomainTrack(item)
	result, err := this.trackRepository.Update(track)
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com erro.", err)
		return &pb.TrackResponse{}, err
	}
	saved := result.(structs.Track)
	log.Info("Vou converter para pb.Track e criar o response.")
	resp := &pb.TrackResponse{
		Track: adapters.ToProtoTrack(saved),
	}
	log.Info("Tudo certo. Vou retornar o response contendo o item com ID:", saved.ID)
	return resp, nil
}

func (this *TrackServer) Delete(ctx context.Context, req *pb.TrackRequest) (*pb.TrackResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetTrack()
	log.Info("Recebi o request e vou converter pb.Track para structs.Track.")
	track := adapters.ToDomainTrack(item)
	_, err := this.trackRepository.Delete(track)
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com erro.", err)
		return &pb.TrackResponse{}, err
	}

	resp := &pb.TrackResponse{
		Track: item,
	}
	log.Info("Tudo certo. Vou retornar o response contendo o item com ID:", item.Id)
	return resp, nil
}

func (this *TrackServer) ReadOne(ctx context.Context, req *pb.TrackRequest) (*pb.TrackResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetTrack()
	log.Info("Recebi o request e vou converter pb.Track para structs.Track.")
	found, err := this.trackRepository.ReadOne(adapters.ToDomainTrack(item))
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com o erro.", err)
		return &pb.TrackResponse{}, err
	}
	track := found.(structs.Track)
	albums, err := this.albumRepository.FindAlbumsByTrackID(track.ID)
	if err != nil {
		log.Error("Falhei ao obter os albums da track", track.Title)
	}
	track.Albums = albums
	resp := &pb.TrackResponse{
		Track: adapters.ToProtoTrack(track),
	}
	log.Info("Tudo certo. Vou retornar o response com a track com ID", item.Id, "contendo", len(albums), "albums.")
	return resp, nil
}

func (this *TrackServer) ReadAll(ctx context.Context, req *pb.SearchRequest) (*pb.TracksResponse, error) {
	log.SetRequestID(req.GetId())
	search := adapters.ToDomainSearch(req.GetSearch())
	log.Info("Recebi o request e vou converter pb.Search para structs.Search.", search.String())
	result, err := this.trackRepository.ReadAll(search)
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com erro.", err)
		return &pb.TracksResponse{}, err
	}

	var tracks []structs.Track
	for _, track := range result.Items.([]structs.Track) {
		albums, err := this.albumRepository.FindAlbumsByTrackID(track.ID)
		if err != nil {
			log.Error("Nao consegui obter os albums da track", track.ID)
		}
		log.Info("A track", track.ID, "possui", len(albums), "albums.")
		track.Albums = albums
		tracks = append(tracks, track)
	}
	resp := &pb.TracksResponse{
		Tracks:     adapters.ToProtoTracks(tracks),
		Pagination: adapters.ToProtoPagination(result.Pagination),
	}
	log.Info("Tudo certo. Vou retornar o response contendo", len(tracks), "tracks.")
	return resp, nil
}
