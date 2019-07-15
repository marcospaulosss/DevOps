package adapters

import (
	"time"

	"backend/apps/elearning/src/structs"
	pb "backend/proto"
)

func ToDomainShelf(in *pb.Shelf) structs.Shelf {
	return structs.Shelf{
		ID:     uint64(in.Id),
		Title:  in.Title,
		Albums: ToDomainAlbums(in.GetAlbums()),
	}
}

func ToProtoShelf(in structs.Shelf) *pb.Shelf {
	updatedAt := ""
	if in.UpdatedAt != nil {
		updatedAt = in.UpdatedAt.Format(time.RFC3339)
	}
	return &pb.Shelf{
		Id:        uint64(in.ID),
		Title:     in.Title,
		CreatedAt: in.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updatedAt,
		Albums:    ToProtoAlbums(in.Albums),
	}
}

func ToProtoShelves(list []structs.Shelf) []*pb.Shelf {
	var items []*pb.Shelf
	for _, item := range list {
		items = append(items, ToProtoShelf(item))
	}
	return items
}
