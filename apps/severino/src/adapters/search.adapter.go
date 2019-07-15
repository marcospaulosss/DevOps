package adapters

import (
	"backend/apps/severino/src/structs"
	log "backend/libs/logger"
	pb "backend/proto"
)

func CreateSearchRequest(search structs.Search) *pb.SearchRequest {
	return &pb.SearchRequest{
		Id: log.GetRequestID(),
		Search: &pb.Search{
			Pagination: ToProtoPagination(search.Pagination),
			Raw:        search.Raw,
			Extra:      search.Extra,
		},
	}
}

func ToDomainPagination(pagination *pb.Pagination) structs.Pagination {
	return structs.Pagination{
		Page:    int(pagination.GetPage()),
		PerPage: int(pagination.GetPerPage()),
		Total:   int(pagination.GetTotal()),
	}
}

func ToProtoPagination(in structs.Pagination) *pb.Pagination {
	return &pb.Pagination{
		PerPage: int32(in.PerPage),
		Page:    int32(in.Page),
		Order:   in.Order,
		SortBy:  in.SortBy,
	}
}
