package repositories

import (
	"backend/apps/elearning/src/structs"
	"backend/libs/databases"
	"backend/libs/errors"
	log "backend/libs/logger"
)

type ShelfRepository struct {
	db           databases.Database
	queryBuilder ShelfQueryBuilder
}

func NewShelfRepository(db databases.Database) Repository {
	return ShelfRepository{
		db:           db,
		queryBuilder: NewShelfQueryBuilder(),
	}
}

func (this ShelfRepository) Create(item interface{}) (interface{}, error) {
	shelf := item.(structs.Shelf)
	log.Info("Vou inserir a shelf:", shelf.Title)
	query := this.queryBuilder.Create()
	return this.upsert(query, shelf)
}

func (this ShelfRepository) Update(item interface{}) (interface{}, error) {
	shelf := item.(structs.Shelf)
	log.Info("Vou alterar a shelf com ID:", shelf.ID)
	query := this.queryBuilder.Update()
	return this.upsert(query, shelf)
}

func (this ShelfRepository) ReadOne(item interface{}) (interface{}, error) {
	s := item.(structs.Shelf)
	log.Info("Vou buscar shelf com ID:", s.ID)
	conn := this.db.GetConnection()
	query := this.queryBuilder.ReadOne()
	var shelf structs.Shelf
	err := conn.Get(&shelf, query, s.ID)
	if err != nil {
		log.Error("Falhei. ID:", s.ID, err)
		return nil, errors.NewGrpcError(errors.NotFound, err.Error())
	}
	log.Info("Sucesso. Achei a shelf.")
	return shelf, nil
}

func (this ShelfRepository) ReadAll(search structs.Search) (structs.Result, error) {
	log.Info("Buscando itens usando search:", search.String())
	query := this.queryBuilder.ReadAll(search)

	var items []structs.Shelf
	conn := this.db.GetConnection()
	err := conn.Select(&items, query)
	if err != nil {
		log.Error("Falhei ao buscar itens.", err)
		return structs.Result{}, errors.NewGrpcError(errors.Invalid, err.Error())
	}
	log.Info("Ok. Encontrei", len(items), "shelves. Vou obter o total de shelves cadastradas...")

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
	log.Info("Sucesso. Retornando", len(items), "/", total, "shelves.")
	return result, nil
}

func (this ShelfRepository) Delete(item interface{}) (interface{}, error) {
	shelf := item.(structs.Shelf)
	log.Info("Vou deletar a shelf com ID:", shelf.ID)
	conn := this.db.GetConnection()
	query := this.queryBuilder.Delete()
	var id uint64
	err := conn.Get(&id, query, shelf.ID)
	if err != nil {
		log.Error("Falhei.", err)
		return 0, errors.NewGrpcError(errors.NotFound, err.Error())
	}
	log.Info("Sucesso. Shelf deletada.")
	return id, nil
}

func (this ShelfRepository) AssociateShelfAndAlbums(shelf structs.Shelf) error {
	log.Info("Vou associar os albums na shelf com ID:", shelf.ID)
	query := this.queryBuilder.AssociateShelfAndAlbums(shelf)
	_, err := this.db.GetConnection().Exec(query, shelf.ID)
	if err != nil {
		log.Error("Falhei.", err)
		return err
	}
	return nil
}

func (this ShelfRepository) DisassociateAlbumsFromShelf(shelf structs.Shelf) error {
	log.Info("Vou desassociar os albums da shelf com ID:", shelf.ID)
	query := this.queryBuilder.DisassociateAlbumsFromShelf()
	_, err := this.db.GetConnection().Exec(query, shelf.ID)
	if err != nil {
		log.Error("Falhei.", err)
		return err
	}
	return nil
}

func (this ShelfRepository) FindShelvesByAlbumID(albumID uint64) []structs.Shelf {
	var shelves []structs.Shelf
	query := this.queryBuilder.FindShelvesByAlbumID()
	err := this.db.GetConnection().Select(&shelves, query, albumID)
	if err != nil {
		log.Error("Falhei ao buscar shelves pelo album com ID", albumID)
	}
	return shelves

}

func (this ShelfRepository) FindHomeShelves(search structs.Search) (structs.Result, error) {
	log.Info("Buscando itens usando search:", search.String())
	query := this.queryBuilder.FindHomeShelves(search)

	var items []structs.Shelf
	conn := this.db.GetConnection()
	err := conn.Select(&items, query)
	if err != nil {
		log.Error("Falhei ao buscar itens.", err)
		return structs.Result{}, errors.NewGrpcError(errors.Invalid, err.Error())
	}
	log.Info("Ok. Encontrei", len(items), "shelves. Vou obter o total de shelves que possuem albuns atribuidos")

	search.Pagination.Total = items[0].Total
	result := structs.Result{
		Items:      items,
		Pagination: search.Pagination,
	}
	log.Info("Sucesso. Retornando", len(items), "/", search.Pagination.Total, "shelves.")
	return result, nil
}

func (this ShelfRepository) upsert(query string, item structs.Shelf) (interface{}, error) {
	conn := this.db.GetConnection()
	nstmt, err := conn.PrepareNamed(query)
	if err != nil {
		log.Error("Falhei no prepare statement.", err)
		return nil, err
	}

	var saved structs.Shelf
	if err = nstmt.Get(&saved, item); err != nil {
		log.Error("Falhei.", item.Title, err)
		return nil, errors.NewGrpcError(errors.Invalid, err.Error())
	}

	log.Info("Ok. Salvei a shelf:", saved.Title, saved.ID)
	return saved, nil
}
