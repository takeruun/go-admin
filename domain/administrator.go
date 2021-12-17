package domain

type Administrator struct {
	Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	FontSize int    `json:"font_size" form:"fontSize"`
}
