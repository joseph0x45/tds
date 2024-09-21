package domain

import (
	"time"
)

type Susbscription struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	BoughtOn  time.Time `json:"bought_on" db:"bought_on"`
	ExpiresOn time.Time `json:"expires_on" db:"expires_on"`
}
