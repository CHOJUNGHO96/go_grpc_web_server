package repository

import (
	"context"
	"sequence_game_server/core/db"
	"sequence_game_server/pkg/test2/model"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *db.PostgresDB
}

func NewRepository(db *db.PostgresDB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Test2Select(ctx context.Context, p map[string]interface{}) ([]model.SelectCompanyDepartment, error) {
	var results []model.SelectCompanyDepartment

	query, args, err := sqlx.In("SELECT * FROM company_department WHERE id IN (?)", p["id"])
	if err != nil {
		return nil, err
	}

	query = r.db.DB.Rebind(query)

	err = r.db.DB.SelectContext(ctx, &results, query, args...)
	if err != nil {
		return nil, err
	}

	return results, nil
}
