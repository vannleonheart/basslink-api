package common

import (
	"CRM/src/lib/basslink"
	"errors"
	"time"

	"gorm.io/gorm"
)

func (s *Service) handleGetPaymentById(id string) (*basslink.RemittancePayment, error) {
	var payment basslink.RemittancePayment

	if err := s.App.DB.Connection.Where("id = ? AND status = ?", id, basslink.PaymentStatusWait).First(&payment).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}

func (s *Service) handlePaymentConfirm(id string, req *PaymentConfirmRequest) error {
	var payment basslink.RemittancePayment

	if err := s.App.DB.Connection.Preload("Remittance").Where("id = ? AND status = ?", id, basslink.PaymentStatusFailed).First(&payment).Error; err != nil {
		return err
	}

	if payment.Remittance == nil {
		return errors.New("remittance data not found")
	}

	if payment.Remittance.Status != basslink.RemittanceStatusWait {
		return errors.New("invalid remittance status")
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.RemittancePayment{}).Where("id = ? AND status = ?", id, basslink.PaymentStatusWait).Updates(map[string]interface{}{
			"status":                basslink.PaymentStatusConfirmed,
			"payment_confirm_time":  req.Date,
			"payment_confirm_proof": req.Proof,
			"payment_reference":     req.Reference,
			"updated":               time.Now().Unix(),
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(basslink.Remittance{}).Where("id = ? AND status = ?", id, basslink.RemittanceStatusWait).Updates(map[string]interface{}{
			"status":  basslink.RemittanceStatusPaymentConfirmed,
			"updated": time.Now().Unix(),
		}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
