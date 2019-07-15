package repositories

import (
	"backend/apps/elearning/src/structs"
	"backend/libs/databases"
	"backend/libs/errors"
	log "backend/libs/logger"
)

type PreferenceRepository struct {
	db           databases.Database
	queryBuilder PreferenceQueryBuilder
}

func NewPreferenceRepository(db databases.Database) Repository {
	return PreferenceRepository{
		db:           db,
		queryBuilder: NewPreferenceQueryBuilder(),
	}
}

func (this PreferenceRepository) Create(item interface{}) (interface{}, error) {
	return nil, nil
}

func (this PreferenceRepository) Update(item interface{}) (interface{}, error) {
	preference := item.(structs.Preference)
	log.Info("Vou atualizar as preferencias da", preference.Type, "contendo:", preference.Content)
	query := this.queryBuilder.Update()
	conn := this.db.GetConnection()
	nstmt, err := conn.PrepareNamed(query)
	if err != nil {
		log.Error("Falhei no prepare statement.", err)
		return nil, errors.NewGrpcError(errors.Invalid, err.Error())
	}

	var saved structs.Preference
	if err = nstmt.Get(&saved, preference); err != nil {
		log.Error("Falhei.", preference.Type, preference.Content, err)
		return nil, errors.NewGrpcError(errors.Invalid, err.Error())
	}

	log.Info("Ok. As preferencias de", saved.Type, "foram salvas com conteudo", saved.Content)
	return saved, nil
}

func (this PreferenceRepository) ReadOne(item interface{}) (interface{}, error) {
	p := item.(structs.Preference)
	log.Info("Vou buscar configuracao de", p.Type)
	conn := this.db.GetConnection()
	query := this.queryBuilder.ReadOne()
	var preference structs.Preference
	err := conn.Get(&preference, query, p.Type)
	if err != nil {
		log.Error("Falhei ao buscar configuracao de", p.Type, err)
		return nil, errors.NewGrpcError(errors.NotFound, err.Error())
	}
	log.Info("Sucesso. Encontrei configuracao existente", preference)
	return preference, nil
}

func (this PreferenceRepository) ReadAll(search structs.Search) (structs.Result, error) {
	return structs.Result{}, nil
}

func (this PreferenceRepository) Delete(item interface{}) (interface{}, error) {
	return nil, nil
}
