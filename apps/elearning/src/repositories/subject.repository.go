package repositories

import (
	"github.com/labstack/gommon/log"

	"backend/apps/elearning/src/structs"
	"backend/libs/databases"
	"backend/libs/errors"
)

type SubjectRepository struct {
	db           databases.Database
	queryBuilder SubjectQueryBuilder
}

func NewSubjectRepository(db databases.Database) Repository {
	return SubjectRepository{
		db:           db,
		queryBuilder: NewSubjectQueryBuilder(),
	}
}

func (this SubjectRepository) Create(item interface{}) (interface{}, error) {
	subject := item.(structs.Subject)
	log.Info("Vou inserir a subject:", subject.Title)
	query := this.queryBuilder.Create()
	return this.upsert(query, subject)
}

func (this SubjectRepository) ReadOne(item interface{}) (interface{}, error) {
	s := item.(structs.Subject)
	log.Info("Vou buscar subject com ID:", s.ID)
	conn := this.db.GetConnection()
	query := this.queryBuilder.ReadOne()
	var subject structs.Subject
	err := conn.Get(&subject, query, s.ID)
	if err != nil {
		log.Error("Falhei. ID:", s.ID, err)
		return nil, errors.NewGrpcError(errors.NotFound, err.Error())
	}
	log.Info("Sucesso. Achei a subject.")
	return subject, nil
}

func (this SubjectRepository) Update(item interface{}) (interface{}, error) {
	subject := item.(structs.Subject)

	log.Info("Vou verificar se a subject existe:", subject.ID)
	query := this.queryBuilder.ReadOne()
	var read structs.Subject
	if err := this.db.GetConnection().Get(&read, query, subject.ID); err != nil {
		log.Error("Falhei. ID nao encontrado:", subject.ID, err)
		return nil, errors.NewGrpcError(errors.NotFound, err.Error())
	}

	log.Info("Vou atualizar a subject:", subject.Title)
	query = this.queryBuilder.Update()
	return this.upsert(query, subject)
}

func (this SubjectRepository) upsert(query string, item structs.Subject) (interface{}, error) {
	conn := this.db.GetConnection()
	nstmt, err := conn.PrepareNamed(query)
	if err != nil {
		log.Error("Falhei no prepare statement.", err)
		return nil, err
	}

	var saved structs.Subject
	if err = nstmt.Get(&saved, item); err != nil {
		log.Error("Falhei.", item.Title, err)
		return nil, errors.NewGrpcError(errors.Unknown, err.Error())
	}

	log.Info("Ok. Salvei o subject:", saved.Title, saved.ID)
	return saved, nil
}

func (this SubjectRepository) ReadAll(search structs.Search) (structs.Result, error) {
	log.Info("Buscando itens usando search:", search.String())
	query := this.queryBuilder.ReadAll(search)

	var items []structs.Subject
	conn := this.db.GetConnection()
	if err := conn.Select(&items, query); err != nil {
		log.Error("Falhei ao buscar itens.", err)
		return structs.Result{}, errors.NewGrpcError(errors.Invalid, err.Error())
	}
	log.Info("Ok. Encontrei", len(items), "subjects. Vou obter o total de subjects cadastradas...")

	var total int32
	if err := conn.Get(&total, this.queryBuilder.Total(search)); err != nil {
		log.Error("Falhei ao obter o total de registros.", err)
		return structs.Result{}, errors.NewGrpcError(errors.Internal, err.Error())
	}

	log.Info("Sucesso. Retornando", len(items), "/", total, "subjects.")

	search.Pagination.Total = total
	return structs.Result{
		Items:      items,
		Pagination: search.Pagination,
	}, nil
}

func (this SubjectRepository) Delete(item interface{}) (interface{}, error) {
	subject := item.(structs.Subject)
	log.Info("Vou deletar a subject com ID:", subject.ID)
	conn := this.db.GetConnection()
	query := this.queryBuilder.Delete()
	var id uint64
	err := conn.Get(&id, query, subject.ID)
	if err != nil {
		log.Error("Falhei.", err)
		return 0, errors.NewGrpcError(errors.NotFound, err.Error())
	}
	log.Info("Sucesso. subject deletada.")
	return id, nil
}
