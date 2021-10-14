package models

type Room struct {
	UUIDBaseModel
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
}
