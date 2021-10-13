package models

type Session struct {
	UUIDBaseModel
	RoomID  string `json:"-"`
	Room    *Room  `json:"room" gorm:"foreignKey:RoomID;references:ID"`
	IsOwner bool   `json:"is_owner"`
}
