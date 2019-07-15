package adapters

import (
	"time"

	"backend/apps/elearning/src/structs"
	pb "backend/proto"
)

func ToDomainSubject(in *pb.Subject) structs.Subject {
	return structs.Subject{
		ID:    uint64(in.Id),
		Title: in.Title,
	}
}

func ToProtoSubject(in structs.Subject) *pb.Subject {
	createdAt := ""
	if in.CreatedAt != nil {
		createdAt = in.CreatedAt.Format(time.RFC3339)
	}

	updatedAt := ""
	if in.UpdatedAt != nil {
		updatedAt = in.UpdatedAt.Format(time.RFC3339)
	}

	return &pb.Subject{
		Id:        uint64(in.ID),
		Title:     in.Title,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func ToProtoSubjects(list []structs.Subject) []*pb.Subject {
	var items []*pb.Subject
	for _, item := range list {
		items = append(items, ToProtoSubject(item))
	}
	return items
}
