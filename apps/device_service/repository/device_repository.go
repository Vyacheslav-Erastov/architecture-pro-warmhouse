package repository

import (
	"context"
	"device_service/models"

	"github.com/jmoiron/sqlx"
)

type DeviceRepository struct {
	db *sqlx.DB
}

func NewDeviceRepository(db *sqlx.DB) *DeviceRepository {
	return &DeviceRepository{db: db}
}

func (r *DeviceRepository) GetAll(ctx context.Context) ([]models.Device, error) {
	devices := []models.Device{}
	err := r.db.SelectContext(ctx, &devices, "SELECT * FROM devices")
	return devices, err
}

func (r *DeviceRepository) Create(ctx context.Context, d models.DeviceCreate) (int, error) {
	var id int
	err := r.db.GetContext(ctx, &id, `
        INSERT INTO devices (name, type, status, location)
        VALUES ($1, $2, $3, $4)
		RETURNING id
    `, d.Name, d.Type, d.Status, d.Location)
	return id, err
}

func (r *DeviceRepository) GetByID(ctx context.Context, id string) (*models.Device, error) {
	var d models.Device
	err := r.db.GetContext(ctx, &d, "SELECT * FROM devices WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *DeviceRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE devices SET status=$1, last_updated=NOW() WHERE id=$2",
		status, id)
	return err
}
