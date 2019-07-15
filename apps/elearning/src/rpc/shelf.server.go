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

type ShelfServer struct {
	shelfRepository repositories.ShelfRepository
	albumRepository repositories.AlbumRepository
}

func NewShelfServer(s *grpc.Server, repos repositories.Container) *ShelfServer {
	shelfRepository := repos.ShelfRepository.(repositories.ShelfRepository)
	albumRepository := repos.AlbumRepository.(repositories.AlbumRepository)
	server := &ShelfServer{shelfRepository, albumRepository}
	if s != nil {
		pb.RegisterShelfServiceServer(s, server)
	}
	return server
}

func (this *ShelfServer) Create(ctx context.Context, req *pb.ShelfRequest) (*pb.ShelfResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetShelf()
	log.Info("Recebi o request e vou converter pb.Shelf para structs.Shelf.")
	shelf := adapters.ToDomainShelf(item)
	result, err := this.shelfRepository.Create(shelf)
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com erro.", err)
		return &pb.ShelfResponse{}, err
	}

	saved := result.(structs.Shelf)
	if len(shelf.Albums) > 0 {
		log.Info("Ok. A shelf possui", len(shelf.Albums), "albums. Vou criar as associacoes...")
		shelf.ID = saved.ID
		err := this.shelfRepository.AssociateShelfAndAlbums(shelf)
		if err == nil {
			log.Info("Feito.")
		}
	}

	log.Info("Vou converter para pb.Shelf e criar o response.")
	resp := &pb.ShelfResponse{
		Shelf: adapters.ToProtoShelf(saved),
	}
	log.Info("Tudo certo. Vou retornar o response contendo o item com ID:", saved.ID)
	return resp, nil
}

func (this *ShelfServer) Update(ctx context.Context, req *pb.ShelfRequest) (*pb.ShelfResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetShelf()
	log.Info("Recebi o request e vou converter pb.Shelf para structs.Shelf.")
	shelf := adapters.ToDomainShelf(item)
	result, err := this.shelfRepository.Update(shelf)
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com erro.", err)
		return &pb.ShelfResponse{}, err
	}

	saved := result.(structs.Shelf)
	if len(shelf.Albums) > 0 {
		log.Info("Ok. A shelf possui", len(shelf.Albums), "albums. Vou remover as associacoes...")
		shelf.ID = saved.ID
		err := this.shelfRepository.DisassociateAlbumsFromShelf(shelf)
		if err == nil {
			log.Info("Desassociei. Vou associar novamente...")
			err = this.shelfRepository.AssociateShelfAndAlbums(shelf)
			if err == nil {
				log.Info("Feito.")
			}
		}
	}

	log.Info("Vou converter para pb.Shelf e criar o response.")
	resp := &pb.ShelfResponse{
		Shelf: adapters.ToProtoShelf(saved),
	}
	log.Info("Tudo certo. Vou retornar o response contendo o item com ID:", saved.ID)
	return resp, nil
}

func (this *ShelfServer) Delete(ctx context.Context, req *pb.ShelfRequest) (*pb.ShelfResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetShelf()
	log.Info("Recebi o request e vou converter pb.Shelf para structs.Shelf.")
	shelf := adapters.ToDomainShelf(item)
	_, err := this.shelfRepository.Delete(shelf)
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com erro.", err)
		return &pb.ShelfResponse{}, err
	}

	resp := &pb.ShelfResponse{
		Shelf: item,
	}
	log.Info("Tudo certo. Vou retornar o response contendo o item com ID:", item.Id)
	return resp, nil
}

func (this *ShelfServer) ReadOne(ctx context.Context, req *pb.ShelfRequest) (*pb.ShelfResponse, error) {
	log.SetRequestID(req.GetId())
	item := req.GetShelf()
	log.Info("Recebi o request e vou converter pb.Shelf para structs.Shelf.")
	found, err := this.shelfRepository.ReadOne(adapters.ToDomainShelf(item))
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com o erro.", err)
		return &pb.ShelfResponse{}, err
	}
	shelf := found.(structs.Shelf)
	log.Info("Vou buscar os albums da shelf com ID:", shelf.ID)
	albums, err := this.albumRepository.FindAlbumsByShelfID(shelf.ID)
	if err != nil {
		log.Error("Falhei.", err)
		return &pb.ShelfResponse{}, err
	}
	for i, _ := range albums {
		teachers := this.albumRepository.GetTeachersByAlbumID(albums[i].ID)
		albums[i].Teachers = teachers
	}
	shelf.Albums = albums

	resp := &pb.ShelfResponse{
		Shelf: adapters.ToProtoShelf(shelf),
	}
	log.Info("Tudo certo. Vou retornar o response contendo o item com ID:", item.Id, "contendo", len(shelf.Albums), "albums.")
	return resp, nil
}

func (this *ShelfServer) ReadAll(ctx context.Context, req *pb.SearchRequest) (*pb.ShelvesResponse, error) {
	log.SetRequestID(req.GetId())
	search := adapters.ToDomainSearch(req.GetSearch())
	log.Info("Recebi o request e vou converter pb.Search para structs.Search.", search.String())
	result, err := this.shelfRepository.ReadAll(search)
	if err != nil {
		log.Error("Tudo errado. Vou retornar o response com erro.", err)
		return &pb.ShelvesResponse{}, err
	}
	items := result.Items.([]structs.Shelf)
	var shelves []structs.Shelf
	var onlyPublishedAlbums bool
	if _, ok := search.GetExtra()["album__is_published"]; ok {
		onlyPublishedAlbums = true
	}
	log.Info("Vou buscar os albums das shelves...")
	for _, shelf := range items {
		var albums []structs.Album
		var err error
		if onlyPublishedAlbums {
			albums, err = this.albumRepository.FindPublishedAlbumsByShelfID(shelf.ID)
		} else {
			albums, err = this.albumRepository.FindAlbumsByShelfID(shelf.ID)
		}
		if err != nil {
			log.Error("Falhei ao obter os albums da shelf:", shelf.Title)
			continue
		}
		for i, _ := range albums {
			teachers := this.albumRepository.GetTeachersByAlbumID(albums[i].ID)
			albums[i].Teachers = teachers
		}
		shelf.Albums = albums
		shelves = append(shelves, shelf)
	}

	resp := &pb.ShelvesResponse{
		Shelves:    adapters.ToProtoShelves(shelves),
		Pagination: adapters.ToProtoPagination(result.Pagination),
	}
	log.Info("Tudo certo. Vou retornar o response contendo", len(shelves), "shelves.")
	return resp, nil
}
