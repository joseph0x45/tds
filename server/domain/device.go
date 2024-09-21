package domain

import (
	"time"
)

type Device struct {
	ID         string `json:"id" db:"id"`
	Version    string `json:"version" db:"version"`
	Model      string `json:"model" db:"model"`
	Number     string `json:"number" db:"number"`
	DataAmount string `json:"data_amount" db:"data_amount"`
}

type DevicePosition struct {
	DeviceID  string    `json:"device_id" db:"device_id"`
	Latitude  float64   `json:"latitude" db:"latitude"`
	Longitude float64   `json:"longitude" db:"longitude"`
	LoggedAt  time.Time `json:"logged_at" db:"logged_at"`
}
