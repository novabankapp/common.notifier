package sms

type SMSService interface {
	SendSMS(sourceAddr, destinationAddr, message string) (bool, error)
}
