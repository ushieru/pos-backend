package dto

type CreateTableDTO struct {
	Name string `json:"name" validate:"required,min=1"`
	PosX int    `json:"pos_x"  validate:"required,gte=0"`
	PosY int    `json:"pos_y"  validate:"required,gte=0"`
}
