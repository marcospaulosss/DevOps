package adapters

import (
	"fmt"

	"github.com/labstack/gommon/log"

	"backend/apps/severino/src/structs"
	"backend/libs/json"
	pb "backend/proto"
)

func ToDomainPreference(in *pb.Preference) structs.Preference {
	var preference structs.Preference
	if err := json.Unmarshal([]byte(in.GetContent()), &preference); err != nil {
		log.Error("Falhou no marshal de content", err)
	}
	return preference
}

func ToProtoPreference(in structs.Preference) *pb.Preference {
	b, err := json.Marshal(in.Shelves)
	if err == nil {

		return &pb.Preference{
			Type:    in.Type,
			Content: fmt.Sprintf(`{"shelves": %s}`, string(b)),
		}
	}
	log.Error("Falhou no marshal de preference.Content", err)
	return &pb.Preference{}
}

func CreatePreferenceRequest(in structs.Preference) *pb.PreferenceRequest {
	return &pb.PreferenceRequest{
		Preference: ToProtoPreference(in),
	}
}
