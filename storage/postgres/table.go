package postgres

import (
	"small_projects/socker_table/storage/repo"

	"github.com/jmoiron/sqlx"
)

type tableRepo struct {
	db *sqlx.DB
}

func NewTableRepo(db *sqlx.DB) repo.TableRepoI {
	return &tableRepo{
		db: db,
	}
}

// func (t tableRepo)
