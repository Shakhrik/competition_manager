package storage

import (
	"small_projects/socker_table/storage/postgres"
	"small_projects/socker_table/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Competition() repo.CompetitionRepoI
}

type storagePG struct {
	db              *sqlx.DB
	competitionRepo repo.CompetitionRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return storagePG{
		db:              db,
		competitionRepo: postgres.NewCompetitionRepo(db),
	}
}

func (s storagePG) Competition() repo.CompetitionRepoI {
	return s.competitionRepo
}
