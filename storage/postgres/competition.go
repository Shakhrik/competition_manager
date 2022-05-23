package postgres

import (
	"small_projects/socker_table/models"
	"small_projects/socker_table/storage/repo"

	"github.com/jmoiron/sqlx"
)

type competitionRepo struct {
	db *sqlx.DB
}

func NewCompetitionRepo(db *sqlx.DB) repo.CompetitionRepoI {
	return &competitionRepo{
		db: db,
	}
}

func (c competitionRepo) Create(req models.Competition) (res models.ResponseWithID, err error) {
	err = c.db.Get(res.ID, "INSERT INTO competition(name, tables_count,teams_count) VALUES($1, $2,$3) RETURNING id", req.Title, req.TablesCount, req.TeamsCount)
	return
}
