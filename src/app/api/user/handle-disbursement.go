package user

import "CRM/src/lib/basslink"

func (s *Service) GetDisbursements(user *basslink.User) (*[]basslink.Disbursement, error) {
	var disbursements []basslink.Disbursement

	if err := s.App.DB.Connection.Where("user_id = ?", user.Id).Find(&disbursements).Error; err != nil {
		return nil, err
	}

	return &disbursements, nil
}

func (s *Service) GetDisbursement(user *basslink.User, disbursementId string) (*basslink.Disbursement, error) {
	var disbursement basslink.Disbursement

	if err := s.App.DB.Connection.Where("user_id = ?", user.Id).First(&disbursement, disbursementId).Error; err != nil {
		return nil, err
	}

	return &disbursement, nil
}

func (s *Service) CreateDisbursement(user *basslink.User) {

}

func (s *Service) UpdateDisbursement(user *basslink.User) {

}

func (s *Service) SubmitDisbursement(user *basslink.User) {

}

func (s *Service) CancelDisbursement(user *basslink.User) {

}
