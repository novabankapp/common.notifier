package email

type mockMailService struct {
}

func (m mockMailService) SendEmail(to []string, message string) (bool, error) {
	return true, nil
}

func NewMockMailService() MailService {
	return &mockMailService{}
}
