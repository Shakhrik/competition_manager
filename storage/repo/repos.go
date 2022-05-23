package repo

import "small_projects/socker_table/models"

type TableRepoI interface {
}

type CompetitionRepoI interface {
	Create(req models.Competition) (res models.ResponseWithID, err error)
}
