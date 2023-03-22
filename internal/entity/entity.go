package entity

type UsualRequest struct {
	ID          ID
	Name        string
	PageRequest PageStr
}

type ID struct {
	Uint uint
	Str  string
}

type PageStr struct {
	PageNumber string
	PerPage    string
}
