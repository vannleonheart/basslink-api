package agent

import (
	"CRM/src/lib/basslink"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

func (s *Service) createDisbursement(agent *basslink.Agent, req *CreateDisbursementRequest) error {
	now := time.Now().Unix()

	disbursementId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	flFromAmount, err := req.FromAmount.Float64()
	if err != nil {
		return err
	}

	flToAmount, err := req.ToAmount.Float64()
	if err != nil {
		return err
	}

	flRate, err := req.Rate.Float64()
	if err != nil {
		return err
	}

	flFee, err := req.FeeAmount.Float64()
	if err != nil {
		return err
	}

	newDisbursement := basslink.Disbursement{
		Id:           disbursementId.String(),
		AgentId:      agent.Id,
		UserId:       nil,
		FromCurrency: req.FromCurrency,
		FromAmount:   flFromAmount,
		ToContact:    "",
		ToCurrency:   req.ToCurrency,
		ToAmount:     flToAmount,
		ToAccount:    "",
		RateCurrency: "",
		Rate:         flRate,
		FeeCurrency:  "",
		FeeAmount:    flFee,
		TransferType: "",
		TransferRef:  req.TransferReference,
		TransferDate: req.TransferDate,
		FundSource:   req.FundSource,
		Purpose:      req.Purpose,
		Notes:        req.Notes,
		Status:       basslink.DisbursementStatusNew,
		IsSettled:    false,
		Created:      now,
		Updated:      nil,
	}

	if err = s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&newDisbursement).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) updateDisbursement(agent *basslink.Agent) {

}

func (s *Service) submitDisbursement(agent *basslink.Agent) {

}

func (s *Service) cancelDisbursement(agent *basslink.Agent) {

}
