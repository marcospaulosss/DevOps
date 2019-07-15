package adapters

import (
	"backend/apps/elearning/src/structs"
	pb "backend/proto"
)

func ToDomainPreference(in *pb.Preference) structs.Preference {
	return structs.Preference{
		Type:    in.GetType(),
		Content: in.GetContent(),
	}
}
