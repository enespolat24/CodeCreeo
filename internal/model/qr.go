package model

type QRCode struct {
	UDID   string `gorm:"primaryKey" json:"udid"`
	Url    string `json:"url"`
	UserID uint   `json:"user_id"`
}
