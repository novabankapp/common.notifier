package sms

type SMPPConfig struct {
	SMSC       string
	SystemID   string
	Password   string
	SystemType string
}
type smppService struct {
	config SMPPConfig
}

func (s smppService) SendSMS(sourceAddr, destinationAddr, message string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewSMPPService(config SMPPConfig) SMSService {
	return &smppService{
		config: config,
	}
}
