package models

type Tables struct {
	Table      map[string]Table
	TableCount int
	TeamsCount int
}
type Table struct {
	// Title      string
	Teams      []string
	TeamsCount int
}
