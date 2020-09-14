package models

type MobileResponse struct {
	UserId        string
	CorrelationId string
	IsValidate    bool
	KeyId         string
	Signature     string
}
