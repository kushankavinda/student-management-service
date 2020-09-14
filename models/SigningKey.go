package models
import "time"
type Key struct {
	UserName string
	KeyId string
	Type int
	PrivateKey   string
	PublicKey   string
	CreatedTime time.Time
	IsActive    bool
	DeviceId    string
}
