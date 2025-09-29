package shared

import "errors"

type Direction string

const (
	desc Direction = "DESC"
	asc  Direction = "ASC"
)

var emptyInput = errors.New("empty input")

type PaginationParams struct {
	SortBy        string
	SortDirection Direction
	PageSize      int
	PageNumber    int
}
