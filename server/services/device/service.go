package device

import (
	"database/sql"
	"errors"
	"net/http"
	"server/domain"
	"server/internal/dtos"
	"server/internal/repository/postgres"
	"server/internal/rest/transport"
	"server/pkg/types"
)

type Service interface {
	RegisterDevice(*dtos.DeviceRegistration) transport.ServiceResponse
	GetAllDevices() transport.ServiceResponse
}

type service struct {
	logger     types.Logger
	deviceRepo postgres.DeviceRepo
}

func NewService(
	logger types.Logger,
	deviceRepo postgres.DeviceRepo,
) Service {
	return &service{
		logger:     logger,
		deviceRepo: deviceRepo,
	}
}

func (s *service) RegisterDevice(d *dtos.DeviceRegistration) transport.ServiceResponse {
	_, err := s.deviceRepo.GetByID(d.ID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			s.logger.Error(err.Error())
			return transport.InternalServerError
		}
		device := &domain.Device{
			ID:         d.ID,
			Version:    d.Version,
			Model:      d.Model,
			Number:     d.Number,
			DataAmount: 0,
		}
		err = s.deviceRepo.Insert(device)
		if err != nil {
			s.logger.Error(err.Error())
			return transport.InternalServerError
		}
		return transport.CreatedNoData
	}
	return transport.ServiceResponse{
		StatusCode: http.StatusConflict,
		Data:       nil,
		Err:        errors.New("duplicate_id"),
	}
}

func (s *service) GetAllDevices() transport.ServiceResponse {
	data, err := s.deviceRepo.GetAll()
	if err != nil {
		s.logger.Error(err.Error())
		return transport.InternalServerError
	}
	return transport.ServiceResponse{
		StatusCode: http.StatusOK,
		Data: map[string]interface{}{
			"devices": data,
		},
		Err: nil,
	}
}
