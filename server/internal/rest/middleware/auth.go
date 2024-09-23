package middleware

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"server/internal/repository/postgres"
	"server/internal/rest/transport"
	"server/pkg/types"
)

type AuthorizationMiddleware struct {
	adminRepo   postgres.AdminRepo
	sessionRepo postgres.SessionRepo
	logger      types.Logger
}

func NewAuthorizationMiddleware(
	adminRepo postgres.AdminRepo,
	sessionRepo postgres.SessionRepo,
	logger types.Logger,
) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{
		adminRepo:   adminRepo,
		sessionRepo: sessionRepo,
		logger:      logger,
	}
}

func (m *AuthorizationMiddleware) Authorize() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionID := r.Header.Get("Authorization")
			if sessionID == "" {
				transport.ErrAuthorizationFailed.Write(w)
				return
			}
			session, err := m.sessionRepo.Get(sessionID)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					transport.ErrAuthorizationFailed.Write(w)
					return
				}
				m.logger.Error(err.Error())
				transport.InternalServerError.Write(w)
				return
			}
			if !session.Valid {
				transport.ErrAuthorizationFailed.Write(w)
				return
			}
			user, err := m.adminRepo.GetByID(session.UserID)
			if err != nil {
				m.logger.Error(err.Error())
				transport.ErrAuthorizationFailed.Write(w)
				return
			}
			ctx := context.WithValue(r.Context(), "user", user)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
