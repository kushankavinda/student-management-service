package models

import "time"

type KeyStore struct {
	ID          string
	UserId      string
	PublicKey   string
	CreatedTime time.Time
	IsActive    bool
	DeviceId    string
	PushToken   string
	KeyId       string
}
