package sms

type mockSMSService struct {
}

func (m mockSMSService) SendSMS(destinationAddr, message string) (bool, error) {

	return true, nil
}

func NewMockSMSService() SMSService {
	return &mockSMSService{}
}
