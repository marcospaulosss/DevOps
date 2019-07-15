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

type AlbumServer struct {
	shelfRepository repositories.ShelfRepository
	albumRepository repositories.AlbumRepository
}

func NewAlbumServer(s *grpc.Server, repos repositories.Container) *AlbumServer {
	shelfRepository := repos.ShelfRepository.(repositories.ShelfRepository)
	albumRepository := repos.AlbumRepository.(repositories.AlbumRepository)
	server := &AlbumServer{shelfRepository, albumRepository}
	if s != nil {
		pb.RegisterAlbumServiceServer(s, server)
	}
	return server
}

func (this *AlbumServer) Create(ctx context.Context, req *pb.AlbumRequest) (*pb.AlbumResponse, error) {
	log.SetRequestID(req.GetId())
	album := adapters.ToDomainAlbum(req.GetAlbum())
	log.Info("Recebi o request e vou converter pb.Album para structs.Album.")
	result, err := this.albumRepository.Create(album)
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com erro.", err)
		return &pb.AlbumResponse{}, err
	}
	saved := result.(structs.Album)
	album.ID = saved.ID
	if err := this.albumRepository.CreateSectionsAndAssociateItWithTracks(album); err != nil {
		return &pb.AlbumResponse{}, err
	}
	log.Info("Ok. Vou converter para pb.Album e criar o response.")
	resp := &pb.AlbumResponse{
		Album: adapters.ToProtoAlbum(saved),
	}
	log.Info("Tudo certo. Vou retornar o response contendo o item com ID:", album.ID)
	return resp, nil
}

func (this *AlbumServer) Update(ctx context.Context, req *pb.AlbumRequest) (*pb.AlbumResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetAlbum()
	log.Info("Recebi o request e vou converter pb.Album para structs.Album.")
	album := adapters.ToDomainAlbum(item)
	result, err := this.albumRepository.Update(album)
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com erro.", err)
		return &pb.AlbumResponse{}, err
	}

	saved := result.(structs.Album)
	log.Info("Vou remover as sections do album:", album.Title)
	if err := this.albumRepository.DeleteSectionsByAlbumID(album); err != nil {
		return &pb.AlbumResponse{}, err
	}
	log.Info("Ok. Vou recriar as sections e associar as tracks novamente.")
	if err := this.albumRepository.CreateSectionsAndAssociateItWithTracks(album); err != nil {
		return &pb.AlbumResponse{}, err
	}
	log.Info("Vou converter para pb.Album e criar o response.")
	resp := &pb.AlbumResponse{
		Album: adapters.ToProtoAlbum(saved),
	}
	log.Info("Tudo certo. Vou retornar o response contendo o item com ID:", saved.ID)
	return resp, nil
}

func (this *AlbumServer) Delete(ctx context.Context, req *pb.AlbumRequest) (*pb.AlbumResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetAlbum()
	log.Info("Recebi o request e vou converter pb.Album para structs.Album.")
	album := adapters.ToDomainAlbum(item)
	_, err := this.albumRepository.Delete(album)
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com erro.", err)
		return &pb.AlbumResponse{}, err
	}

	resp := &pb.AlbumResponse{
		Album: item,
	}
	log.Info("Tudo certo. Vou retornar o response contendo o item com ID:", item.Id)
	return resp, nil
}

func (this *AlbumServer) ReadOne(ctx context.Context, req *pb.AlbumRequest) (*pb.AlbumResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetAlbum()
	log.Info("Vou buscar os albums da album com ID:", item.Id)
	found, err := this.albumRepository.ReadOne(adapters.ToDomainAlbum(item))
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com o erro.", err)
		return &pb.AlbumResponse{}, err
	}
	log.Info("Achei. Vou buscar as sections desse album.")
	album := found.(structs.Album)
	sections := this.albumRepository.FetchSectionsByAlbumID(album.ID)
	log.Info("Encontrei", len(sections), "sections. Vou obter os teachers...")
	album.Sections = sections
	album.Teachers = this.albumRepository.GetTeachersByAlbumID(album.ID)
	log.Info("Ok. Vou obter as shelves...")
	album.Shelves = this.shelfRepository.FindShelvesByAlbumID(album.ID)
	log.Info("Ok. Achei", len(album.Shelves), "shelves.")

	resp := &pb.AlbumResponse{
		Album: adapters.ToProtoAlbum(album),
	}
	log.Info("Tudo certo. Vou retornar o response contendo o item com ID:", item.Id)
	return resp, nil
}

func (this *AlbumServer) ReadAll(ctx context.Context, req *pb.SearchRequest) (*pb.AlbumsResponse, error) {
	log.SetRequestID(req.GetId())
	search := adapters.ToDomainSearch(req.GetSearch())
	log.Info("Recebi o request e vou converter pb.Search para structs.Search:", search.String())
	result, err := this.albumRepository.ReadAll(search)
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com erro.", err)
		return &pb.AlbumsResponse{}, err
	}
	albums := result.Items.([]structs.Album)
	for index, _ := range albums {
		album := &albums[index]
		sections := this.albumRepository.FetchSectionsByAlbumID(album.ID)
		album.Sections = sections
		album.Teachers = this.albumRepository.GetTeachersByAlbumID(album.ID)
		album.Shelves = this.shelfRepository.FindShelvesByAlbumID(album.ID)
	}

	resp := &pb.AlbumsResponse{
		Albums:     adapters.ToProtoAlbums(albums),
		Pagination: adapters.ToProtoPagination(result.Pagination),
	}
	log.Info("Tudo certo. Vou retornar o response contendo", len(albums), "albums.")
	return resp, nil
}

func (this *AlbumServer) Publish(ctx context.Context, req *pb.AlbumRequest) (*pb.AlbumResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetAlbum()
	log.Info("Vou publicar o album com ID:", item.Id)
	saved, err := this.albumRepository.Publish(adapters.ToDomainAlbum(item))
	if err != nil {
		log.Error("Nao publicou. Vou retornar o response com o erro.", err)
		return &pb.AlbumResponse{}, err
	}
	log.Info("Ok. Publicou.")
	resp := &pb.AlbumResponse{
		Album: adapters.ToProtoAlbum(saved),
	}
	log.Info("retornando o response do elearning.")
	return resp, nil
}

func (this *AlbumServer) Unpublish(ctx context.Context, req *pb.AlbumRequest) (*pb.AlbumResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetAlbum()
	log.Info("Vou despublicar o album com ID:", item.Id)
	saved, err := this.albumRepository.Unpublish(adapters.ToDomainAlbum(item))
	if err != nil {
		log.Error("Nao despublicou. Vou retornar o response com o erro.", err)
		return &pb.AlbumResponse{}, err
	}
	log.Info("Ok. Despublicou.")
	resp := &pb.AlbumResponse{
		Album: adapters.ToProtoAlbum(saved),
	}
	return resp, nil
}
