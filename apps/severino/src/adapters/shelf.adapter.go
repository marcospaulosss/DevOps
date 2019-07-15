package adapters

import (
	"backend/apps/severino/src/structs"
	log "backend/libs/logger"
	pb "backend/proto"
)

func CreateShelfRequest(item structs.Shelf) *pb.ShelfRequest {
	log.Info("Vou criar o request para shelf com com ID:", item.ID)
	req := &pb.ShelfRequest{
		Shelf: ToProtoShelf(item),
		Id:    log.GetRequestID(),
	}
	log.Info("Criei. Vou enviar o request:", req.String())
	return req
}

func ToProtoShelves(shelves []structs.Shelf) []*pb.Shelf {
	result := []*pb.Shelf{}
	for _, item := range shelves {
		shelf := ToProtoShelf(item)
		result = append(result, shelf)
	}
	return result
}

func ToProtoShelf(shelf structs.Shelf) *pb.Shelf {
	return &pb.Shelf{
		Id:     shelf.ID,
		Title:  shelf.Title,
		Albums: ToProtoAlbums(shelf.Albums),
	}
}

func ToDomainShelves(t []*pb.Shelf) []structs.Shelf {
	shelves := []structs.Shelf{}
	if t == nil {
		return shelves
	}
	for _, shelf := range t {
		shelves = append(shelves, ToDomainShelf(shelf))
	}
	return shelves
}

func ToDomainShelf(t *pb.Shelf) structs.Shelf {
	shelf := structs.Shelf{}
	if t != nil {
		shelf.ID = t.GetId()
		shelf.Title = t.GetTitle()
		shelf.CreatedAt = t.GetCreatedAt()
		shelf.UpdatedAt = t.GetUpdatedAt()
		shelf.Albums = ToDomainAlbums(t.GetAlbums())
	}
	return shelf
}
