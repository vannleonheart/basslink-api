package admin

import "CRM/src/lib/basslink"

func (s *Service) getSenders() (*[]basslink.Sender, error) {
	var senders []basslink.Sender

	if err := s.App.DB.Connection.Find(&senders).Error; err != nil {
		return nil, err
	}

	return &senders, nil
}

func (s *Service) getSenderById(senderId string) (*basslink.Sender, error) {
	var sender basslink.Sender

	if err := s.App.DB.Connection.Where("id = ?", senderId).First(&sender).Error; err != nil {
		return nil, err
	}

	return &sender, nil
}
