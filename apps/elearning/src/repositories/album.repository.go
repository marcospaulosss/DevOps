package repositories

import (
	"backend/apps/elearning/src/structs"
	"backend/libs/databases"
	"backend/libs/errors"
	log "backend/libs/logger"
	"fmt"
)

type AlbumRepository struct {
	db           databases.Database
	queryBuilder AlbumQueryBuilder
}

func NewAlbumRepository(db databases.Database) Repository {
	return AlbumRepository{
		db:           db,
		queryBuilder: NewAlbumQueryBuilder(),
	}
}

func (this AlbumRepository) Create(item interface{}) (interface{}, error) {
	album := item.(structs.Album)
	log.Info("Vou inserir o album:", album.Title)
	query := this.queryBuilder.Create()
	return this.upsert(query, album)
}

func (this AlbumRepository) Update(item interface{}) (interface{}, error) {
	album := item.(structs.Album)
	log.Info("Vou atualizar o album:", album.Title)
	query := this.queryBuilder.Update()
	return this.upsert(query, album)
}

func (this AlbumRepository) CreateSectionsAndAssociateItWithTracks(album structs.Album) error {
	log.Info("Vou inserir", len(album.Sections), "sections para o album com ID:", album.ID)
	conn := this.db.GetConnection()
	query := this.queryBuilder.CreateSection()
	for index, section := range album.Sections {
		var sectionID uint64
		if err := conn.Get(&sectionID, query, album.ID, section.Title, section.Description, index); err != nil {
			log.Error("Nao consegui criar a section com title:", section.Title, err)
			continue
		}

		if len(section.Tracks) == 0 {
			log.Info("Nao vou associar a section", section.Title, "pois ela nao possui tracks.")
			continue
		}

		section.ID = sectionID
		qry := this.queryBuilder.AssociateTracksToSection(section)
		if err := this.changeAssociation(qry, map[string]interface{}{"section_id": sectionID}); err != nil {
			log.Error("Nao consegui associar as tracks com a section ID", sectionID, err)
			return errors.NewGrpcError(errors.Invalid, err.Error())
		}
	}
	return nil
}

func (this AlbumRepository) DeleteSectionsByAlbumID(album structs.Album) error {
	query := this.queryBuilder.DeleteSectionsByAlbumID()
	return this.changeAssociation(query, map[string]interface{}{
		"album_id": album.ID,
	})
}

func (this AlbumRepository) changeAssociation(query string, data map[string]interface{}) error {
	conn := this.db.GetConnection()
	_, err := conn.NamedExec(query, data)
	if err != nil {
		log.Error("Falhei.", err)
		return errors.NewGrpcError(errors.Invalid, err.Error())
	}
	return nil
}

func (this AlbumRepository) upsert(query string, item structs.Album) (interface{}, error) {
	conn := this.db.GetConnection()
	nstmt, err := conn.PrepareNamed(query)
	if err != nil {
		log.Error("Falhei no prepare statement.", err)
		return nil, errors.NewGrpcError(errors.Invalid, err.Error())
	}

	var saved structs.Album
	if err = nstmt.Get(&saved, item); err != nil {
		log.Error("Falhei.", item.Title, err)
		return nil, errors.NewGrpcError(errors.Invalid, err.Error())
	}

	log.Info("Ok. Salvei o album:", saved.Title, saved.ID)
	return saved, nil
}

func (this AlbumRepository) ReadAll(search structs.Search) (structs.Result, error) {
	log.Info("Buscando itens usando search:", search.String())
	query := this.queryBuilder.ReadAll(search)
	fmt.Println("query", query)

	var items []structs.Album
	err := this.db.GetConnection().Select(&items, query)
	if err != nil {
		log.Error("Falhei ao buscar itens.", err)
		return structs.Result{}, errors.NewGrpcError(errors.Internal, err.Error())
	}
	log.Info("Ok. Encontrei", len(items), "albums,")

	var total int32
	if err := this.db.GetConnection().Get(&total, this.queryBuilder.Total(search)); err != nil {
		log.Error("Falhei ao obter o total de registros.", err)
		return structs.Result{}, errors.NewGrpcError(errors.Internal, err.Error())
	}

	search.Pagination.Total = total

	result := structs.Result{
		Items:      items,
		Pagination: search.Pagination,
	}
	log.Info("Sucesso. Retornando", len(items), "/", total, "albums.")
	return result, nil
}

func (this AlbumRepository) FetchSectionsByAlbumID(albumID uint64) []structs.Section {
	query := this.queryBuilder.GetSectionsByAlbumID()
	rows, err := this.db.GetConnection().Queryx(query, albumID)
	var sections []structs.Section
	if err != nil {
		log.Error("Falhei ao buscar sections.", err)
		return sections
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		log.Error("Falhei ao obter os dados das sections e converter para struct.", err)
		return sections
	}
	for rows.Next() {
		var section structs.Section
		err = rows.StructScan(&section)
		if err != nil {
			log.Error("Nao consegui converter a section para struct.", err)
			continue
		}
		section.Tracks = this.getTracksBySectionID(section.ID)
		sections = append(sections, section)
	}
	return sections
}

func (this AlbumRepository) getTracksBySectionID(sectionID uint64) []structs.Track {
	query := this.queryBuilder.GetTracksBySectionID()
	var tracks []structs.Track
	conn := this.db.GetConnection()
	err := conn.Select(&tracks, query, sectionID)
	if err != nil {
		log.Error("Falhei ao executar query de obter tracks.", err)
	}
	return tracks
}

func (this AlbumRepository) GetTeachersByAlbumID(albumID uint64) string {
	query := this.queryBuilder.GetTeachersByAlbumID()
	var teachers string
	err := this.db.GetConnection().Get(&teachers, query, albumID)
	if err != nil {
		log.Error("Falhei ao obter os professores do album:", albumID)
	}
	return teachers
}

func (this AlbumRepository) ReadOne(item interface{}) (interface{}, error) {
	album := item.(structs.Album)
	log.Info("Vou buscar o album com ID:", album.ID)
	query := this.queryBuilder.ReadOne()
	conn := this.db.GetConnection()
	var a, shelves string
	row := conn.QueryRow(query, album.ID)
	err := row.Scan(&a, &shelves)
	if err != nil {
		log.Error("Falhei. ID:", album.ID, err)
		return nil, errors.NewGrpcError(errors.NotFound, err.Error())
	}
	result := structs.ParseAlbum(a)
	result.Shelves = structs.ParseShelves(shelves)
	log.Info("Ok. Achei o album. Vou obter as sections...")
	return result, nil
}

func (this AlbumRepository) Delete(item interface{}) (interface{}, error) {
	album := item.(structs.Album)
	log.Info("Vou deletar o album com ID:", album.ID)
	conn := this.db.GetConnection()
	query := this.queryBuilder.Delete()
	var id uint64
	err := conn.Get(&id, query, album.ID)
	if err != nil {
		log.Error("Falhei.", err)
		return 0, errors.NewGrpcError(errors.NotFound, err.Error())
	}
	log.Info("Sucesso. Album deletado.")
	return id, nil
}

func (this AlbumRepository) FindAlbumsByShelfID(id uint64) ([]structs.Album, error) {
	query := this.queryBuilder.FindAlbumsByShelfID()
	conn := this.db.GetConnection()
	var albums []structs.Album
	err := conn.Select(&albums, query, id)
	return albums, err
}

func (this AlbumRepository) FindPublishedAlbumsByShelfID(id uint64) ([]structs.Album, error) {
	query := this.queryBuilder.FindPublishedAlbumsByShelfID()
	conn := this.db.GetConnection()
	var albums []structs.Album
	err := conn.Select(&albums, query, id)
	return albums, err
}

func (this AlbumRepository) FindAlbumsByTrackID(id uint64) ([]structs.Album, error) {
	query := this.queryBuilder.FindAlbumsByTrackID()
	conn := this.db.GetConnection()
	var albums []structs.Album
	err := conn.Select(&albums, query, id)
	return albums, err
}

func (this AlbumRepository) Publish(album structs.Album) (structs.Album, error) {
	query := this.queryBuilder.Publish()
	_, err := this.db.GetConnection().Exec(query, album.ID)
	return album, err
}

func (this AlbumRepository) Unpublish(album structs.Album) (structs.Album, error) {
	query := this.queryBuilder.Unpublish()
	_, err := this.db.GetConnection().Exec(query, album.ID)
	return album, err
}
