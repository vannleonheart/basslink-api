package common

import "CRM/src/lib/basslink"

func (s *Service) handleGetCurrencies() (*[]basslink.Currency, error) {
	var currencies []basslink.Currency

	if err := s.App.DB.Connection.Find(&currencies).Error; err != nil {
		return nil, err
	}

	return &currencies, nil
}
