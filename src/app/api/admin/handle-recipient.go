package admin

import "CRM/src/lib/basslink"

func (s *Service) getRecipients() (*[]basslink.Recipient, error) {
	var recipients []basslink.Recipient

	if err := s.App.DB.Connection.Find(&recipients).Error; err != nil {
		return nil, err
	}

	return &recipients, nil
}

func (s *Service) getRecipient(recipientId string) (*basslink.Recipient, error) {
	var recipient basslink.Recipient

	if err := s.App.DB.Connection.Preload("Documents").Preload("Accounts").Where("id = ?", recipientId).First(&recipient).Error; err != nil {
		return nil, err
	}

	return &recipient, nil
}
