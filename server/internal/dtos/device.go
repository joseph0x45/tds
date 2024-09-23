package dtos

type DeviceRegistration struct {
	ID      string `json:"id" db:"id"`
	Version string `json:"version" db:"version"`
	Model   string `json:"model" db:"model"`
	Number  string `json:"number" db:"number"`
}

func (d *DeviceRegistration) Validate() error {
	return nil
}
