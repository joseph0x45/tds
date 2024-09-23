package postgres

import (
	"fmt"
	"server/domain"

	"github.com/jmoiron/sqlx"
)

type SessionRepo interface {
	Insert(*domain.Session) error
	Get(string) (*domain.Session, error)
}

type sessionRepo struct {
	db *sqlx.DB
}

func NewSessionRepo(db *sqlx.DB) SessionRepo {
	return &sessionRepo{db}
}

func (r *sessionRepo) Insert(s *domain.Session) error {
	const query = `
    insert into sessions(
      id, user_id
    )
    values(
      :id, :user_id
    )
  `
	_, err := r.db.NamedExec(query, s)
	if err != nil {
		return fmt.Errorf("Error while inserting session: %w", err)
	}
	return nil
}

func (r *sessionRepo) Get(id string) (*domain.Session, error) {
	s := &domain.Session{}
	const query = "select * from sessions where id=$1"
	err := r.db.Get(s, query)
	if err != nil {
		return nil, fmt.Errorf("Error while getting session: %w", err)
	}
	return s, nil
}
