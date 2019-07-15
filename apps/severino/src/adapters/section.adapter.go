package adapters

import (
	"backend/apps/severino/src/structs"
	pb "backend/proto"
)

func ToProtoSection(section structs.Section) *pb.Section {
	return &pb.Section{
		Id:          section.ID,
		Title:       section.Title,
		Description: section.Description,
		Tracks:      ToProtoTracks(section.Tracks),
	}
}

func ToProtoSections(list []structs.Section) []*pb.Section {
	var items []*pb.Section
	for _, item := range list {
		items = append(items, ToProtoSection(item))
	}
	return items
}

func ToDomainSections(a []*pb.Section) []structs.Section {
	sections := []structs.Section{}
	if a == nil {
		return sections
	}
	for _, section := range a {
		sections = append(sections, ToDomainSection(section))
	}
	return sections
}

func ToDomainSection(a *pb.Section) structs.Section {
	section := structs.Section{}
	if a != nil {
		section.ID = a.GetId()
		section.Title = a.GetTitle()
		section.Description = a.GetDescription()
		section.Tracks = ToDomainTracks(a.GetTracks())
	}
	return section
}
