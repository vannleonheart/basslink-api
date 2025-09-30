package admin

import "CRM/src/lib/basslink"

func (s *Service) getUsers() (*[]basslink.User, error) {
	var users []basslink.User

	if err := s.App.DB.Connection.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (s *Service) getUser(userId string) (*basslink.User, error) {
	var user basslink.User

	if err := s.App.DB.Connection.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
