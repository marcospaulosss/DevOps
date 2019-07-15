package adapters

import (
	"strings"
	"fmt"

	"backend/apps/severino/src/structs"
	"backend/libs/configuration"
	log "backend/libs/logger"
	pb "backend/proto"
)

func CreateAlbumRequest(item structs.Album) *pb.AlbumRequest {
	log.Info("Vou criar o request para album com com ID:", item.ID)
	req := &pb.AlbumRequest{
		Album: ToProtoAlbum(item),
		Id:    log.GetRequestID(),
	}
	log.Info("Criei. Vou enviar o request:", req.String())
	return req
}

func ToProtoAlbums(a []structs.Album) []*pb.Album {
	var albums []*pb.Album
	for _, album := range a {
		albums = append(albums, ToProtoAlbum(album))
	}
	return albums
}

func ToProtoAlbum(album structs.Album) *pb.Album {
	return &pb.Album{
		Id:          album.ID,
		Title:       album.Title,
		Description: album.Description,
		Image:       album.Image,
		IsPublished: album.IsPublished,
		Shelves:     ToProtoShelves(album.Shelves),
		Sections:    ToProtoSections(album.Sections),
	}
}

func ToDomainAlbums(a []*pb.Album) []structs.Album {
	albums := []structs.Album{}
	if a == nil {
		return albums
	}
	for _, album := range a {
		albums = append(albums, ToDomainAlbum(album))
	}
	return albums
}

func ToDomainAlbum(a *pb.Album) structs.Album {
	album := structs.Album{}
	if a != nil {
		album.ID = a.GetId()
		album.CreatedAt = a.GetCreatedAt()
		album.UpdatedAt = a.GetUpdatedAt()
		album.Title = a.GetTitle()
		album.Description = a.GetDescription()
		album.Image = addMediaUrlPrefixTo(a.GetImage())
		album.IsPublished = a.GetIsPublished()
		album.PublishedAt = a.GetPublishedAt()
		album.Shelves = ToDomainShelves(a.GetShelves())
		album.Sections = ToDomainSections(a.GetSections())
		album.Teachers = strings.Split(a.GetTeachers(), ",")
	}
	return album
}

func addMediaUrlPrefixTo(value string) string {
	splited := strings.Split(value, ".")
	if len(splited) > 1 {
		filename := splited[0]
		ext := splited[1]
		media := configuration.Get().GetEnvConfString("media_url")
		return fmt.Sprintf(`%s/%s/%s-raw.%s`, media, filename, filename, ext)
	}
	return value
}
