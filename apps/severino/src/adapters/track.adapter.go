package adapters

import (
	"backend/apps/severino/src/structs"
	log "backend/libs/logger"
	pb "backend/proto"
)

func CreateTrackRequest(item structs.Track) *pb.TrackRequest {
	log.Info("Vou criar o request para track com com ID:", item.ID)
	req := &pb.TrackRequest{
		Track: ToProtoTrack(item),
		Id:    log.GetRequestID(),
	}
	log.Info("Criei. Vou enviar o request:", req.String())
	return req
}

func ToProtoTracks(tracks []structs.Track) []*pb.Track {
	result := []*pb.Track{}
	for _, item := range tracks {
		track := ToProtoTrack(item)
		result = append(result, track)
	}
	return result
}

func ToProtoTrack(track structs.Track) *pb.Track {
	return &pb.Track{
		Id:          track.ID,
		CreatedAt:   track.CreatedAt,
		UpdatedAt:   track.UpdatedAt,
		Title:       track.Title,
		Description: track.Description,
		Teachers:    track.Teachers,
		Duration:    track.Duration,
		Media:       track.Media,
		Subject:     &pb.Subject{Id: track.Subject},
	}
}

func ToDomainTracks(t []*pb.Track) []structs.Track {
	tracks := []structs.Track{}
	if t == nil {
		return tracks
	}
	for _, track := range t {
		tracks = append(tracks, ToDomainTrack(track))
	}
	return tracks
}

func ToDomainTrack(t *pb.Track) structs.Track {
	track := structs.Track{}
	if t != nil {
		track.ID = t.GetId()
		track.Title = t.GetTitle()
		track.CreatedAt = t.GetCreatedAt()
		track.UpdatedAt = t.GetUpdatedAt()
		track.Description = t.GetDescription()
		track.Teachers = t.GetTeachers()
		track.Duration = t.GetDuration()
		track.Media = addMediaUrlPrefixTo(t.GetMedia())
		track.Albums = ToDomainAlbums(t.GetAlbums())
	}
	return track
}
