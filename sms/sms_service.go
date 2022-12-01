package sms

type SMSService interface {
	SendSMS(destinationAddr, message string) (bool, error)
}
