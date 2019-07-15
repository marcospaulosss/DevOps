package repositories

import (
	"fmt"

	"backend/apps/accounts/src/structs"
	"backend/libs/databases"
	"backend/libs/errors"
	log "backend/libs/logger"
)

type UserRepository struct {
	db           databases.Database
	queryBuilder UserQueryBuilder
}

func NewUserRepository(db databases.Database) Repository {
	return UserRepository{
		db:           db,
		queryBuilder: NewUserQueryBuilder(),
	}
}

func (this UserRepository) Create(data interface{}) (interface{}, error) {
	log.Info("Irei criar um usuario", fmt.Sprintf("%+v", data))

	query := this.queryBuilder.CreateUser()
	conn := this.db.GetConnection()
	nstmt, err := conn.PrepareNamed(query)
	if err != nil {
		log.Error("Falhei no PrepareNamed", query, err.Error())
		return nil, err
	}

	var id string
	if err = nstmt.Get(&id, data); err != nil {
		log.Error("Falhei no Get", query, err.Error())
		return nil, err
	}

	saved := &structs.User{
		ID: id,
	}

	log.Info("Criei novo usuario com ID", id)
	return saved, nil
}

func (this UserRepository) ReadAll(search structs.Search) (structs.Result, error) {
	log.Info("Buscando itens usando search:", search.String())
	query := this.queryBuilder.ReadAllUsers(search)
	var items []structs.User
	err := this.db.GetConnection().Select(&items, query)
	if err != nil {
		log.Error("Falhei ao buscar itens.", err)
		return structs.Result{}, errors.NewGrpcError(errors.Internal, err.Error())
	}

	type Users struct {
		Total int32 `db:"total"`
	}
	var users Users
	err = this.db.GetConnection().Get(&users, this.queryBuilder.CountUsers(search))
	if err != nil {
		log.Error("Falhei ao buscar itens.", err)
		return structs.Result{}, errors.NewGrpcError(errors.Internal, err.Error())
	}
	log.Info("Total de itens retornados do banco de dados", users.Total)
	search.Pagination.Total = users.Total

	result := structs.Result{
		Items:      items,
		Pagination: search.Pagination,
	}

	return result, nil

}

func (this UserRepository) ReadOne(data interface{}) (interface{}, error) {

	return nil, nil
}

func (this UserRepository) Update(data interface{}) (interface{}, error) {

	return nil, nil
}

func (this UserRepository) Delete(data interface{}) (interface{}, error) {

	return nil, nil
}
