package repositories

import (
	"backend/apps/elearning/src/structs"
	"backend/libs/databases"
	"backend/libs/errors"
	log "backend/libs/logger"
)

type TrackRepository struct {
	db           databases.Database
	queryBuilder TrackQueryBuilder
}

func NewTrackRepository(db databases.Database) TrackRepository {
	return TrackRepository{
		db:           db,
		queryBuilder: NewTrackQueryBuilder(),
	}
}

func (this TrackRepository) Create(item interface{}) (interface{}, error) {
	track := item.(structs.Track)
	log.Info("Vou inserir a track:", track.Title)
	query := this.queryBuilder.Create()
	return this.upsert(query, track)
}

func (this TrackRepository) Update(item interface{}) (interface{}, error) {
	track := item.(structs.Track)
	log.Info("Vou alterar a track com ID:", track.ID)
	query := this.queryBuilder.Update()
	return this.upsert(query, track)
}

func (this TrackRepository) ReadOne(item interface{}) (interface{}, error) {
	track := item.(structs.Track)
	log.Info("Vou buscar o track com ID:", track.ID)
	conn := this.db.GetConnection()
	query := this.queryBuilder.ReadOne()
	var result structs.Track
	err := conn.Get(&result, query, track.ID)
	if err != nil {
		log.Error("Falhei. ID:", track.ID, err)
		return nil, errors.NewGrpcError(errors.NotFound, err.Error())
	}
	log.Info("Sucesso. Achei a track", result.Title)
	return result, nil
}

func (this TrackRepository) ReadAll(search structs.Search) (structs.Result, error) {
	log.Info("Buscando itens usando search:", search.String())
	query := this.queryBuilder.ReadAll(search)

	conn := this.db.GetConnection()
	var items []structs.Track
	err := conn.Select(&items, query)
	if err != nil {
		log.Error("Falhei ao buscar itens.", err, query)
		return structs.Result{}, errors.NewGrpcError(errors.Invalid, err.Error())
	}

	var total int32
	err = conn.Get(&total, this.queryBuilder.Total(search))
	if err != nil {
		log.Error("Falhei ao obter o total de registros.", err)
		return structs.Result{}, errors.NewGrpcError(errors.Internal, err.Error())
	}

	search.Pagination.Total = total
	result := structs.Result{
		Items:      items,
		Pagination: search.Pagination,
	}
	log.Info("Sucesso. Retornando", len(items), "/", total, "tracks.")
	return result, nil
}

func (this TrackRepository) Delete(item interface{}) (interface{}, error) {
	track := item.(structs.Track)
	log.Info("Vou deletar a track com ID:", track.ID)
	conn := this.db.GetConnection()
	query := this.queryBuilder.Delete()
	var id uint64
	err := conn.Get(&id, query, track.ID)
	if err != nil {
		log.Error("Falhei.", err)
		return 0, errors.NewGrpcError(errors.NotFound, err.Error())
	}
	log.Info("Sucesso. Track deletada.")
	return id, nil
}

func (this TrackRepository) upsert(query string, item structs.Track) (interface{}, error) {
	conn := this.db.GetConnection()
	nstmt, err := conn.PrepareNamed(query)
	if err != nil {
		log.Error("Falhei no prepare statement.", err)
		return nil, err
	}

	var saved structs.Track
	if err = nstmt.Get(&saved, item); err != nil {
		log.Error("Falhei.", item.Title, err)
		return nil, errors.NewGrpcError(errors.Invalid, err.Error())
	}

	log.Info("Ok. Salvei a track:", saved.Title, saved.ID)
	return saved, nil
}

func (this TrackRepository) CreateSubjectsAndAssociateItWithTracks(item interface{}) error {
	track := item.(structs.Track)
	log.Info("Vou associar a materia id:", track.Subject.ID, "a uma faixa:", track.ID)
	query := this.queryBuilder.AssociateSubjectsAndTracks()
	conn := this.db.GetConnection()
	if _, err := conn.Exec(query, track.Subject.ID, track.ID); err != nil {
		log.Error("Nao consegui criar a associacao entre materia e faixa.:")
		return errors.NewGrpcError(errors.Invalid, err.Error())
	}
	return nil
}
