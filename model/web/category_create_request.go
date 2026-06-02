package web

type CategoryCreateRequest struct {
	Name string `validate:"required,gt=0,lt=201" json:"name"`
}
