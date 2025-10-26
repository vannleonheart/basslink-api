package agent

import (
	"CRM/src/lib/basslink"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

func (s *Service) getSubmissions(req *GetRemittanceFilter) (*[]basslink.Remittance, error) {
	var remittances []basslink.Remittance

	db := s.App.DB.Connection.Preload("SourceCurrency").Preload("TargetCurrency").Preload("Attachments")

	if req != nil {
		if req.Type != nil && *req.Type != "" && strings.ToLower(*req.Type) != "all" {
			db = db.Where("transfer_type = ?", *req.Type)
		}

		if req.Status != nil && *req.Status != "" && strings.ToLower(*req.Status) != "all" {
			db = db.Where("status = ?", *req.Status)
		}

		if req.Search != nil && *req.Search != "" {
			searchText := strings.ToUpper(*req.Search)
			db = db.Where("id LIKE ?", "%"+searchText+"%")
		}

		if req.Start != nil {
			if startTimestamp, err := time.Parse("2006-01-02", *req.Start); err == nil {
				db = db.Where("created >= ?", startTimestamp.Unix())
			}
		}

		if req.End != nil {
			if endTimestamp, err := time.Parse("2006-01-02", *req.End); err == nil {
				db = db.Where("created <= ?", endTimestamp.Unix())
			}
		}

		if req.FromCurrency != nil && *req.FromCurrency != "" {
			db = db.Where("from_currency = ?", *req.FromCurrency)
		}

		if req.ToCurrency != nil && *req.ToCurrency != "" {
			db = db.Where("to_currency = ?", *req.ToCurrency)
		}
	}

	if err := db.Where("status = ?", basslink.RemittanceStatusSubmitted).Find(&remittances).Error; err != nil {
		return nil, err
	}

	return &remittances, nil
}

func (s *Service) approveSubmission(agent *basslink.Agent, id string) error {
	var remittance basslink.Remittance

	if err := s.App.DB.Connection.Preload("SourceCurrency").Preload("TargetCurrency").Preload("Attachments").Where("id = ? AND status = ?", id, basslink.RemittanceStatusSubmitted).First(&remittance).Error; err != nil {
		return err
	}

	now := time.Now().Unix()
	bankInfo := s.App.GetBankInfo(remittance.FromCurrency)
	newRemittancePayment := basslink.RemittancePayment{
		Id:                  remittance.Id,
		Currency:            remittance.FromCurrency,
		Amount:              remittance.FromAmount,
		PaymentMethod:       "bank_transfer",
		PaymentData:         fmt.Sprintf("%s, no. %s a.n %s", bankInfo.BankName, bankInfo.AccountNo, bankInfo.AccountOwner),
		PaymentConfirmTime:  nil,
		PaymentConfirmProof: nil,
		PaymentReference:    nil,
		Status:              basslink.PaymentStatusWait,
		Created:             now,
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.Remittance{}).Where("id = ? AND status = ?", remittance.Id, basslink.RemittanceStatusSubmitted).Updates(map[string]interface{}{
			"agent_id": agent.Id,
			"status":   basslink.RemittanceStatusWait,
			"updated":  now,
		}).Error; err != nil {
			return err
		}

		if err := tx.Create(&newRemittancePayment).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	if remittance.NotificationEmail != nil && len(*remittance.NotificationEmail) > 0 {
		s.App.EmailMsgChannel <- &basslink.EmailNotificationMesage{
			To:       *remittance.NotificationEmail,
			Subject:  fmt.Sprintf("%s - %s", remittance.Id, "Selesaikan pembayaran untuk melanjutkan proses kirim dana"),
			Template: "remittance-submitted:1",
			Data: map[string]interface{}{
				"id":                        remittance.Id,
				"sender_name":               remittance.FromName,
				"recipient_name":            remittance.ToName,
				"to_currency":               remittance.TargetCurrency.Symbol,
				"to_amount":                 s.App.FormatCurrency(fmt.Sprintf("%f", remittance.ToAmount)),
				"bank_name":                 bankInfo.BankName,
				"bank_code":                 bankInfo.BankCode,
				"bank_swift":                bankInfo.SwiftCode,
				"account_no":                bankInfo.AccountNo,
				"account_name":              bankInfo.AccountOwner,
				"currency":                  remittance.SourceCurrency.Symbol,
				"amount":                    s.App.FormatCurrency(fmt.Sprintf("%f", remittance.FromAmount)),
				"payment_confirmation_link": fmt.Sprintf("%s/%s", s.App.Config.PaymentConfirmationLink, remittance.Id),
			},
		}
	}

	return nil
}
