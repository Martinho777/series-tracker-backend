package models

type SeriesFilters struct {
	Query string
	Sort  string
	Order string
	Page  int
	Limit int
}