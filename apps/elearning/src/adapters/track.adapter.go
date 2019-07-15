package adapters

import (
	"time"

	"backend/apps/elearning/src/structs"
	pb "backend/proto"
)

func ToProtoTracks(list []structs.Track) []*pb.Track {
	var items []*pb.Track
	for _, item := range list {
		items = append(items, ToProtoTrack(item))
	}
	return items
}

func ToProtoTrack(in structs.Track) *pb.Track {
	updatedAt := ""
	if in.UpdatedAt != nil {
		updatedAt = in.UpdatedAt.Format(time.RFC3339)
	}
	return &pb.Track{
		Id:          uint64(in.ID),
		Title:       in.Title,
		CreatedAt:   in.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   updatedAt,
		Description: in.Description,
		Teachers:    in.Teachers,
		Duration:    in.Duration,
		Media:       in.Media,
		Albums:      ToProtoAlbums(in.Albums),
	}
}

func ToDomainTracks(list []*pb.Track) []structs.Track {
	var items []structs.Track
	for _, item := range list {
		items = append(items, ToDomainTrack(item))
	}
	return items
}

func ToDomainTrack(in *pb.Track) structs.Track {
	return structs.Track{
		ID:          uint64(in.Id),
		Title:       in.Title,
		Description: in.Description,
		Teachers:    in.Teachers,
		Duration:    in.Duration,
		Media:       in.Media,
		Subject:     ToDomainSubject(in.Subject),
	}
}
