package adapters

import (
	"backend/apps/elearning/src/structs"
	pb "backend/proto"
)

func ToProtoSections(list []structs.Section) []*pb.Section {
	var items []*pb.Section
	for _, item := range list {
		items = append(items, ToProtoSection(item))
	}
	return items
}

func ToProtoSection(in structs.Section) *pb.Section {
	return &pb.Section{
		Id:          in.ID,
		Title:       in.Title,
		Description: in.Description,
		Tracks:      ToProtoTracks(in.Tracks),
	}
}

func ToDomainSections(list []*pb.Section) []structs.Section {
	var items []structs.Section
	for _, item := range list {
		items = append(items, ToDomainSection(item))
	}
	return items
}

func ToDomainSection(in *pb.Section) structs.Section {
	return structs.Section{
		ID:          in.Id,
		Title:       in.Title,
		Description: in.Description,
		Tracks:      ToDomainTracks(in.Tracks),
	}
}
