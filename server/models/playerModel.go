package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Player struct {
	UUID      string `json:"uuid" gorm:"primaryKey"`
	CreateAt  time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string   `json:"name"`
	Age         uint     `json:"age"`
	Personality string   `json:"personality"`
	Position    string   `json:"position"`
	LeftFoot    string   `json:"left_foot"`
	RightFoot   string   `json:"right_foot"`
	Height      uint     `json:"height"`
	Weight      uint     `json:"weight"`
	Attributs   Attribut `gorm:"foreignKey:UUID;references:UUID""`
}

func (player *Player) BeforeCreate(ctx *gorm.DB) (err error) {
	player.UUID = uuid.NewString()

	if uuid.Validate(player.UUID) != nil {
		err = errors.New("can't save invalid data")
	}

	return
}
