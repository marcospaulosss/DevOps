package adapters

import (
	"time"

	"backend/apps/elearning/src/structs"
	pb "backend/proto"
)

func ToProtoAlbum(in structs.Album) *pb.Album {
	var updatedAt, publishedAt string
	if in.UpdatedAt != nil {
		updatedAt = in.UpdatedAt.Format(time.RFC3339)
	}

	if in.PublishedAt != nil {
		publishedAt = in.PublishedAt.Format(time.RFC3339)
	}
	return &pb.Album{
		Id:          uint64(in.ID),
		CreatedAt:   in.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   updatedAt,
		Title:       in.Title,
		Description: in.Description,
		Image:       in.Image,
		IsPublished: in.IsPublished,
		PublishedAt: publishedAt,
		Teachers:    in.Teachers,
		Sections:    ToProtoSections(in.Sections),
		Shelves:     ToProtoShelves(in.Shelves),
	}
}

func ToProtoAlbums(list []structs.Album) []*pb.Album {
	var items []*pb.Album
	for _, item := range list {
		items = append(items, ToProtoAlbum(item))
	}
	return items
}

func ToDomainAlbums(list []*pb.Album) []structs.Album {
	var items []structs.Album
	for _, item := range list {
		items = append(items, ToDomainAlbum(item))
	}
	return items
}

func ToDomainAlbum(in *pb.Album) structs.Album {
	return structs.Album{
		ID:          uint64(in.Id),
		Title:       in.Title,
		Description: in.Description,
		Image:       in.Image,
		IsPublished: in.IsPublished,
		Sections:    ToDomainSections(in.Sections),
		Teachers:    in.Teachers,
	}
}
