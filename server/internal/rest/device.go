package rest

import (
	"encoding/json"
	"net/http"
	"server/internal/dtos"
	"server/internal/rest/middleware"
	"server/services/device"

	"github.com/go-chi/chi/v5"
)

type DeviceHandler struct {
	service device.Service
}

func (h *DeviceHandler) HandleRegisterDevice(w http.ResponseWriter, r *http.Request) {
	payload := &dtos.DeviceRegistration{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.service.RegisterDevice(payload).Write(w)
}

func (h *DeviceHandler) HandleGetAllDevices(w http.ResponseWriter, r *http.Request) {
	h.service.GetAllDevices().Write(w)
}

func NewDeviceHandler(
	r chi.Router,
	service device.Service,
	m *middleware.AuthorizationMiddleware,
) {
	h := &DeviceHandler{service: service}

	r.With(m.Authorize()).Route("/devices", func(r chi.Router) {
		r.Post("/", h.HandleRegisterDevice)
		r.Get("/", h.HandleGetAllDevices)
	})
}
