package adapters

import (
	"backend/apps/elearning/src/structs"
	pb "backend/proto"
)

func ToDomainSearch(in *pb.Search) structs.Search {
	return structs.Search{
		Pagination: ToDomainPagination(in.Pagination),
		Raw:        in.GetRaw(),
		Extra:      in.GetExtra(),
	}
}

func ToProtoSearch(in structs.Search) *pb.Search {
	return &pb.Search{
		Pagination: ToProtoPagination(in.Pagination),
		Raw:        in.Raw,
	}
}

func ToDomainPagination(in *pb.Pagination) structs.Pagination {
	return structs.Pagination{
		Order:   in.GetOrder(),
		Page:    in.GetPage(),
		PerPage: in.GetPerPage(),
		SortBy:  in.GetSortBy(),
	}
}

func ToProtoPagination(in structs.Pagination) *pb.Pagination {
	return &pb.Pagination{
		Order:   in.Order,
		Page:    in.Page,
		PerPage: in.PerPage,
		SortBy:  in.SortBy,
		Total:   in.Total,
	}
}
