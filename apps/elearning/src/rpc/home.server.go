package rpc

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc"

	"backend/apps/elearning/src/adapters"
	"backend/apps/elearning/src/repositories"
	"backend/apps/elearning/src/structs"
	"backend/libs/json"
	log "backend/libs/logger"
	pb "backend/proto"
)

type HomeServer struct {
	preferenceRepository repositories.PreferenceRepository
	shelfRepository      repositories.ShelfRepository
	albumRepository      repositories.AlbumRepository
}

func NewHomeServer(s *grpc.Server, repos repositories.Container) *HomeServer {
	preferenceRepository := repos.PreferenceRepository.(repositories.PreferenceRepository)
	shelfRepository := repos.ShelfRepository.(repositories.ShelfRepository)
	albumRepository := repos.AlbumRepository.(repositories.AlbumRepository)
	server := &HomeServer{preferenceRepository, shelfRepository, albumRepository}
	if s != nil {
		pb.RegisterHomeServiceServer(s, server)
	}
	return server
}

func (this *HomeServer) ReadAll(ctx context.Context, req *pb.SearchRequest) (*pb.ShelvesResponse, error) {
	log.SetRequestID(req.GetId())
	log.Info("Recebi o request e vou converter pb.Search para structs.Search.")
	result, err := this.preferenceRepository.ReadOne(structs.Preference{Type: "home"})
	if err != nil {
		log.Error("Houve um problema no preferenceRepository. Vou retornar o response com erro.", err)
		return &pb.ShelvesResponse{}, err
	}
	preference := result.(structs.Preference)
	var home structs.Home
	err = json.Unmarshal([]byte(preference.Content), &home)
	if err != nil {
		log.Error("Houve um problema ao fazer o unmarshal do preference de home. Vou retornar o response com erro.", err)
		return &pb.ShelvesResponse{}, err
	}

	page := req.GetSearch().GetPagination().GetPage()
	per_page := req.GetSearch().GetPagination().GetPerPage()

	shelvesFilter := arrayUint64ToString(home.Shelves, ",")
	resultShelves, err := this.shelfRepository.FindHomeShelves(structs.Search{
		Raw: shelvesFilter,
		Pagination: structs.Pagination{
			Page:    page,
			PerPage: per_page,
		},
	})
	shelves := resultShelves.Items.([]structs.Shelf)

	for k, _ := range shelves {
		resultAlbums, err := this.albumRepository.FindAlbumsByShelfID(shelves[k].ID)
		if err != nil {
			log.Error("Houve um problema ao buscar os albums da shelf", shelves[k].ID, ". Vou retornar o response com erro")
			return &pb.ShelvesResponse{}, err
		}

		this.addTeachersToAlbums(resultAlbums)

		shelves[k].Albums = resultAlbums
	}

	resp := &pb.ShelvesResponse{
		Shelves:    adapters.ToProtoShelves(shelves),
		Pagination: adapters.ToProtoPagination(resultShelves.Pagination),
	}
	return resp, nil
}

func (this *HomeServer) addTeachersToAlbums(albums []structs.Album) {
	for i, _ := range albums {
		teachers := this.albumRepository.GetTeachersByAlbumID(albums[i].ID)
		albums[i].Teachers = teachers
	}
}

func arrayUint64ToString(a []uint64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
