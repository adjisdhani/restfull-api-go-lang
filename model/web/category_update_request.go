package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required,numeric" json:"id"`
	Name string `validate:"required" json:"name"`
}
