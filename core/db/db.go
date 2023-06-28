package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type PostgresDB struct {
	DB *sqlx.DB
}

func (pg *PostgresDB) NewPgConnection(dns string) (*sqlx.DB, error) {
	var err error
	pg.DB, err = sqlx.Connect("postgres", dns)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to Postgresql")
	}

	return pg.DB, nil
}

func (pg *PostgresDB) Close() {
	if err := pg.DB.Close(); err != nil {
		panic(err)
	}
}
