package domain

type Category struct {
	Model
	Name     string `json:"name" form:"name"`
	Products []Product
}
