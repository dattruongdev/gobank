package domain

import "context"

type Preset struct {
	PayeeID string  `json:"payee_id" gorm:"primaryKey"`
	Payee   AppUser `json:"payee" gorm:"foreignKey:PayeeID"`
	PayerID string  `json:"payer_id" gorm:"primaryKey"`
}

type PresetRepository interface {
	Create(c context.Context, p Preset) error
	GetAll(c context.Context, userId string) ([]Preset, error)
	Delete(c context.Context, payerId string, payerIds ...string) error
}
