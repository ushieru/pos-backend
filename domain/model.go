package domain

import (
	"time"

	"github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Model struct {
	ID        string         `json:"id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time      `json:"create_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at" gorm:"index" swaggertype:"string"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := gonanoid.New()
	if err != nil {
		panic(err)
	}
	m.ID = id
	return
}
