package models

import (
	"encoding/json"
	"time"
)

type DeviceCreate struct {
	Name     string          `db:"name" json:"name"`
	Type     string          `db:"type" json:"type" validate:"oneof=sensor gates heating_controller light_bulb camera"`
	Status   json.RawMessage `db:"status" json:"status"`
	Location string          `db:"location" json:"location"`
}

type Device struct {
	DeviceCreate
	ID          int       `db:"id" json:"id"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	LastUpdated time.Time `db:"last_updated" json:"last_updated"`
}
