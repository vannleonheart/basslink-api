package agent

import "CRM/src/lib/basslink"

func (s *Service) getDisbursements(agent *basslink.Agent) (*[]basslink.Disbursement, error) {
	var disbursements []basslink.Disbursement

	if err := s.App.DB.Connection.Where("agent_id = ?", agent.Id).Find(&disbursements).Error; err != nil {
		return nil, err
	}

	return &disbursements, nil
}

func (s *Service) getDisbursement(agent *basslink.Agent, disbursementId string) (*basslink.Disbursement, error) {
	var disbursement basslink.Disbursement

	if err := s.App.DB.Connection.Where("agent_id = ?", agent.Id).First(&disbursement, disbursementId).Error; err != nil {
		return nil, err
	}

	return &disbursement, nil
}

func (s *Service) createDisbursement(agent *basslink.Agent, req *CreateDisbursementRequest) {

}

func (s *Service) updateDisbursement(agent *basslink.Agent) {

}

func (s *Service) submitDisbursement(agent *basslink.Agent) {

}

func (s *Service) cancelDisbursement(agent *basslink.Agent) {

}
