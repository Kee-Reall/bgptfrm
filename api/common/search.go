package common

type Direction string

func (d Direction) String() string {
	if d == "desc" {
		return "desc"
	}
	return "asc"
}

func (d Direction) IsValid() bool {
	return d == "desc" || d == "asc"
}

func (d Direction) IsAsc() bool {
	return d == "asc"
}

func (d Direction) IsDesc() bool {
	return d == "desc"
}

type PaginationParams struct {
	Page      int       `json:"page"`
	PageSize  int       `json:"page_size"`
	Direction Direction `json:"sortDirection" default:"asc"`
}
