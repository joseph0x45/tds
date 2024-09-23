package postgres

import (
	"fmt"
	"server/domain"

	"github.com/jmoiron/sqlx"
)

type DeviceRepo interface {
	Insert(*domain.Device) error
	GetAll() ([]domain.Device, error)
	GetByID(string) (*domain.Device, error)
}

type deviceRepo struct {
	db *sqlx.DB
}

func NewDeviceRepo(db *sqlx.DB) DeviceRepo {
	return &deviceRepo{db}
}

func (r *deviceRepo) Insert(d *domain.Device) error {
	const query = `
    insert into devices(
      id, version, model, number, data_amount
    )
    values(
      :id, :version, :model, :number, :data_amount
    )
  `
	_, err := r.db.NamedExec(query, d)
	if err != nil {
		return fmt.Errorf("Error while inserting new device: %w", err)
	}
	return nil
}

func (r *deviceRepo) GetAll() ([]domain.Device, error) {
	data := make([]domain.Device, 0)
	const query = "select * from devices"
	err := r.db.Select(&data, query)
	if err != nil {
		return nil, fmt.Errorf("Error while getting all devices: %w", err)
	}
	return data, err
}

func (r *deviceRepo) GetByID(id string) (*domain.Device, error) {
	d := &domain.Device{}
	const query = "select * from devices where id=$1"
	err := r.db.Get(d, query, id)
	if err != nil {
		return nil, fmt.Errorf("Error while getting device by ID: %w", err)
	}
	return d, nil
}
