package models

type Room struct {
	UUIDBaseModel
	Password string `json:"password,omitempty"`
	Token string `json:"token,omitempty"`
}
