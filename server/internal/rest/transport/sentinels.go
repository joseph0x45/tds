package transport

import (
	"errors"
	"net/http"
)

var (
	InternalServerError = ServiceResponse{
		StatusCode: http.StatusInternalServerError,
		Data:       nil,
		Err:        errors.New("InternalServerError"),
	}
	ErrSessionCreationFailed = ServiceResponse{
		StatusCode: http.StatusInternalServerError,
		Data:       nil,
		Err:        errors.New("SessionCreationFailed"),
	}
	ErrAuthorizationFailed = ServiceResponse{
		StatusCode: http.StatusUnauthorized,
		Data:       nil,
		Err:        errors.New("AuthorizationFailed"),
	}
	ErrForbidden = ServiceResponse{
		StatusCode: http.StatusForbidden,
		Data:       nil,
		Err:        errors.New("AccessForbidden"),
	}
)

var (
	OkNoData = ServiceResponse{
		StatusCode: http.StatusOK,
		Data:       nil,
		Err:        nil,
	}
	CreatedNoData = ServiceResponse{
		StatusCode: http.StatusCreated,
		Data:       nil,
		Err:        nil,
	}
)
