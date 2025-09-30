package admin

import "CRM/src/lib/basslink"

func (s *Service) getDisbursements() (*[]basslink.Disbursement, error) {
	var disbursements []basslink.Disbursement

	if err := s.App.DB.Connection.Preload("User").Preload("Contact").Preload("TargetAccount").Preload("TargetCurrency").Find(&disbursements).Error; err != nil {
		return nil, err
	}

	return &disbursements, nil
}

func (s *Service) getDisbursement(disbursementId string) (*basslink.Disbursement, error) {
	var disbursement basslink.Disbursement

	if err := s.App.DB.Connection.First(&disbursement, disbursementId).Error; err != nil {
		return nil, err
	}

	return &disbursement, nil
}
