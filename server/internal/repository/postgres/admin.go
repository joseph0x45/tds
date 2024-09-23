package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"server/domain"

	"github.com/jmoiron/sqlx"
)

type AdminRepo interface {
	Insert(*domain.Admin) error
	GetByID(string) (*domain.Admin, error)
	GetAll() ([]domain.Admin, error)
}

type adminRepo struct {
	db *sqlx.DB
}

func NewAdminRepo(db *sqlx.DB) AdminRepo {
	return &adminRepo{db}
}

func (r *adminRepo) Insert(a *domain.Admin) error {
	const query = `
    insert into admins(
      username, password
    )
    values (
      :username, :password)
  `
	_, err := r.db.NamedExec(query, a)
	if err != nil {
		return fmt.Errorf("Error while inserting admin: %w", err)
	}
	return nil
}

func (r *adminRepo) GetAll() ([]domain.Admin, error) {
	const query = "select * from admins"
	data := make([]domain.Admin, 0)
	err := r.db.Select(&data, query)
	if err != nil {
		return nil, fmt.Errorf("Error while getting all admins: %w", err)
	}
	return nil, nil
}

func (r *adminRepo) GetByID(id string) (*domain.Admin, error) {
	admin := &domain.Admin{}
	const query = "select * from admins where id=$1"
	err := r.db.Get(admin, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, fmt.Errorf("Error while getting admin by ID: %w", err)
	}
	return admin, nil
}
