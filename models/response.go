package models

type ResponseWithID struct {
	ID int64 `json:"id" db:"id"`
}
