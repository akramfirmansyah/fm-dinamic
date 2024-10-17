package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attribut struct {
	UUID      string `json:"uuid" gorn:"primaryKey"`
	CreateAt  time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (att *Attribut) BeforeCreate(ctx *gorm.DB) (err error) {
	att.UUID = uuid.NewString()

	if uuid.Validate(att.UUID) != nil {
		err = errors.New("can't save invalid data")
	}

	return
}
