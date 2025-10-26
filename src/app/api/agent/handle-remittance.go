package agent

import (
	"CRM/src/lib/basslink"
	"errors"
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/vannleonheart/goutil"
	"gorm.io/gorm"
)

func (s *Service) generateTransactionId(transactionType string) (string, error) {
	prefix := "TRX"

	switch transactionType {
	case "domestic":
		prefix = "DOM"
	case "international":
		prefix = "INT"
	}

	randomizer := goutil.NewRandomString(goutil.AlphaUNumCharset)
	sufix := randomizer.GenerateRange(2, 4)
	transactionId := fmt.Sprintf("%s-%s-%s", prefix, time.Now().Format("20060102"), sufix)

	var remittances []basslink.Remittance

	if err := s.App.DB.Connection.Where("id = ?", transactionId).Limit(1).Find(&remittances).Error; err != nil {
		return "", err
	}

	if len(remittances) > 0 {
		return s.generateTransactionId(transactionType)
	}

	return transactionId, nil
}

func (s *Service) getRemittances(agent *basslink.Agent, req *GetRemittanceFilter) (*[]basslink.Remittance, error) {
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

	if err := db.Where("agent_id = ?", agent.Id).Find(&remittances).Error; err != nil {
		return nil, err
	}

	return &remittances, nil
}

func (s *Service) getRemittance(agent *basslink.Agent, remittanceId string) (*basslink.Remittance, error) {
	var remittance basslink.Remittance

	if err := s.App.DB.Connection.Preload("SourceCurrency").Preload("TargetCurrency").Preload("Attachments").Where("id = ?", remittanceId).First(&remittance).Error; err != nil {
		return nil, err
	}

	if remittance.Status != basslink.RemittanceStatusSubmitted && remittance.AgentId != agent.Id {
		return nil, errors.New("record not found")
	}

	return &remittance, nil
}

func (s *Service) createRemittance(agent *basslink.Agent, req *CreateRemittanceRequest) error {
	now := time.Now().Unix()

	remittanceId, err := s.generateTransactionId(req.TransferType)
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

	flFeePercent, err := req.FeePercent.Float64()
	if err != nil {
		return err
	}

	flFeeFixed, err := req.FeeFixed.Float64()
	if err != nil {
		return err
	}

	var sender *basslink.Sender
	var recipient *basslink.Recipient
	var updateSenderData, updateRecipientData *map[string]interface{}

	gender := ""
	if req.SenderGender != nil && len(*req.SenderGender) > 0 {
		gender = *req.SenderGender
	}

	birthdate := ""
	if req.SenderBirthdate != nil && len(*req.SenderBirthdate) > 0 {
		birthdate = *req.SenderBirthdate
	}

	if req.SenderId != nil {
		var existingSender basslink.Sender

		if err = s.App.DB.Connection.Where("id = ?", *req.SenderId).First(&existingSender).Error; err != nil {
			return err
		}

		sender = &existingSender
		updateSenderData = &map[string]interface{}{
			"sender_type":         req.SenderType,
			"name":                req.SenderName,
			"gender":              gender,
			"birthdate":           birthdate,
			"citizenship":         req.SenderCitizenship,
			"identity_type":       req.SenderIdentityType,
			"identity_no":         req.SenderIdentityNo,
			"registered_country":  req.SenderRegisteredCountry,
			"registered_region":   req.SenderRegisteredRegion,
			"registered_city":     req.SenderRegisteredCity,
			"registered_address":  req.SenderRegisteredAddress,
			"registered_zip_code": req.SenderRegisteredZipCode,
			"country":             req.SenderCountry,
			"region":              req.SenderRegion,
			"city":                req.SenderCity,
			"address":             req.SenderAddress,
			"zip_code":            req.SenderAddress,
			"contact":             req.SenderContact,
			"occupation":          req.SenderOccupation,
			"pep_status":          req.SenderPepStatus,
			"notes":               req.SenderNotes,
			"updated":             now,
			"updated_by":          agent.Id,
		}
	} else {
		newUserId, e := uuid.NewV7()
		if e != nil {
			return e
		}

		sender = &basslink.Sender{
			Id:                newUserId.String(),
			SenderType:        req.SenderType,
			Name:              req.SenderName,
			Gender:            gender,
			Birthdate:         birthdate,
			Citizenship:       req.SenderCitizenship,
			IdentityType:      req.SenderIdentityType,
			IdentityNo:        req.SenderIdentityNo,
			RegisteredCountry: req.SenderRegisteredCountry,
			RegisteredRegion:  req.SenderRegisteredRegion,
			RegisteredCity:    req.SenderRegisteredCity,
			RegisteredAddress: req.SenderRegisteredAddress,
			RegisteredZipCode: req.SenderRegisteredZipCode,
			Country:           req.SenderCountry,
			Region:            req.SenderRegion,
			City:              req.SenderCity,
			Address:           req.SenderAddress,
			ZipCode:           req.SenderZipCode,
			Contact:           req.SenderContact,
			Occupation:        req.SenderOccupation,
			PepStatus:         req.SenderPepStatus,
			Notes:             req.SenderNotes,
			Created:           now,
			CreatedBy:         agent.Id,
		}
	}

	if req.RecipientId != nil {
		var existingRecipient basslink.Recipient

		if err = s.App.DB.Connection.Where("id = ?", *req.RecipientId).First(&existingRecipient).Error; err != nil {
			return err
		}

		recipient = &existingRecipient
		updateRecipientData = &map[string]interface{}{
			"recipient_type":     req.RecipientType,
			"relationship":       req.RecipientRelationship,
			"name":               req.RecipientName,
			"country":            req.RecipientCountry,
			"region":             req.RecipientRegion,
			"city":               req.RecipientCity,
			"address":            req.RecipientAddress,
			"zip_code":           req.RecipientZipCode,
			"contact":            req.RecipientContact,
			"pep_status":         req.RecipientPepStatus,
			"account_type":       req.RecipientAccountType,
			"bank_name":          req.RecipientBankName,
			"bank_code":          req.RecipientBankCode,
			"bank_account_no":    req.RecipientBankAccountNo,
			"bank_account_owner": req.RecipientBankAccountOwner,
			"notes":              req.RecipientNotes,
			"updated":            now,
		}
	} else {
		newrecipientId, e := uuid.NewV7()
		if e != nil {
			return e
		}

		recipient = &basslink.Recipient{
			Id:               newrecipientId.String(),
			SenderId:         sender.Id,
			RecipientType:    req.RecipientType,
			Relationship:     req.RecipientRelationship,
			Name:             req.RecipientName,
			Country:          req.RecipientCountry,
			Region:           req.RecipientRegion,
			City:             req.RecipientCity,
			Address:          req.RecipientAddress,
			ZipCode:          req.RecipientZipCode,
			Contact:          req.RecipientContact,
			PepStatus:        req.RecipientPepStatus,
			AccountType:      req.RecipientAccountType,
			BankName:         req.RecipientBankName,
			BankCode:         req.RecipientBankCode,
			BankAccountNo:    req.RecipientBankAccountNo,
			BankAccountOwner: req.RecipientBankAccountOwner,
			Notes:            req.RecipientNotes,
			Created:          now,
		}
	}

	totalFee := flFeeFixed
	if flFeePercent > 0 {
		totalFee += math.Ceil((flFeePercent * flFromAmount) / 100)
	}

	if flToAmount > ((flFromAmount - totalFee) * flRate) {
		return errors.New("received amount is greater than expected amount")
	}

	createdBy := fmt.Sprintf("%s:%s:%s", "agent", agent.Id, agent.Name)

	newRemittance := basslink.Remittance{
		Id:                    remittanceId,
		AgentId:               agent.Id,
		SenderId:              sender.Id,
		FromCurrency:          req.FromCurrency,
		FromAmount:            flFromAmount,
		FromSenderType:        req.SenderType,
		FromName:              req.SenderName,
		FromGender:            gender,
		FromBirthdate:         birthdate,
		FromCitizenship:       req.SenderCitizenship,
		FromIdentityType:      req.SenderIdentityType,
		FromIdentityNo:        req.SenderIdentityNo,
		FromRegisteredCountry: req.SenderRegisteredCountry,
		FromRegisteredRegion:  req.SenderRegisteredRegion,
		FromRegisteredCity:    req.SenderRegisteredCity,
		FromRegisteredAddress: req.SenderRegisteredAddress,
		FromRegisteredZipCode: req.SenderRegisteredZipCode,
		FromCountry:           req.SenderCountry,
		FromRegion:            req.SenderRegion,
		FromCity:              req.SenderCity,
		FromAddress:           req.SenderAddress,
		FromZipCode:           req.SenderZipCode,
		FromContact:           req.SenderContact,
		FromOccupation:        req.SenderOccupation,
		FromPepStatus:         req.SenderPepStatus,
		FromNotes:             req.SenderNotes,
		RecipientId:           recipient.Id,
		ToCurrency:            req.ToCurrency,
		ToAmount:              flToAmount,
		ToRecipientType:       req.RecipientType,
		ToRelationship:        req.RecipientRelationship,
		ToName:                req.RecipientName,
		ToCountry:             req.RecipientCountry,
		ToRegion:              req.RecipientRegion,
		ToCity:                req.RecipientCity,
		ToAddress:             req.RecipientAddress,
		ToZipCode:             req.RecipientZipCode,
		ToContact:             req.RecipientContact,
		ToPepStatus:           req.RecipientPepStatus,
		ToAccountType:         req.RecipientAccountType,
		ToBankName:            req.RecipientBankName,
		ToBankCode:            req.RecipientBankCode,
		ToBankAccountNo:       req.RecipientBankAccountNo,
		ToBankAccountOwner:    req.RecipientBankAccountOwner,
		ToNotes:               req.Notes,
		RateCurrency:          req.FromCurrency,
		Rate:                  flRate,
		FeeCurrency:           req.FromCurrency,
		FeeAmountPercent:      flFeePercent,
		FeeAmountFixed:        flFeeFixed,
		FeeTotal:              totalFee,
		PaymentMethod:         req.PaymentMethod,
		TransferType:          req.TransferType,
		TransferRef:           req.TransferReference,
		FundSource:            req.FundSource,
		Purpose:               req.Purpose,
		Notes:                 req.Notes,
		Status:                basslink.RemittanceStatusWait,
		IsSettled:             false,
		CreatedBy:             &createdBy,
		ApprovedBy:            nil,
		ReleasedBy:            nil,
		ApprovedAt:            nil,
		ReleasedAt:            nil,
		Created:               now,
		Updated:               nil,
		SourceCurrency:        nil,
		TargetCurrency:        nil,
		Attachments:           nil,
	}

	var files []basslink.RemittanceAttachment

	if req.Files != nil && len(*req.Files) > 0 {
		for _, file := range *req.Files {
			if len(file) == 0 {
				continue
			}
			if newFileId, e := uuid.NewV7(); e == nil {
				files = append(files, basslink.RemittanceAttachment{
					Id:           newFileId.String(),
					RemittanceId: newRemittance.Id,
					Attachment:   file,
					SubmitBy:     agent.Id,
					SubmitTime:   now,
				})
			}
		}
	}

	if err = s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if req.SenderId == nil {
			if err = tx.Create(sender).Error; err != nil {
				return err
			}
		} else {
			if err = tx.Model(basslink.Sender{}).Where("id = ?", sender.Id).Updates(*updateSenderData).Error; err != nil {
				return err
			}
		}

		if req.RecipientId == nil {
			if err = tx.Create(recipient).Error; err != nil {
				return err
			}
		} else {
			if err = tx.Model(basslink.Recipient{}).Where("id = ?", recipient.Id).Updates(*updateRecipientData).Error; err != nil {
				return err
			}
		}

		if err = tx.Create(&newRemittance).Error; err != nil {
			return err
		}

		if len(files) > 0 {
			if err = tx.CreateInBatches(&files, len(files)).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) cancelRemittance(agent *basslink.Agent, id string, req *RemittanceCancelRequest) error {
	var remittance basslink.Remittance

	if err := s.App.DB.Connection.Preload("SourceCurrency").Preload("TargetCurrency").Preload("Attachments").Where("id = ?", id).First(&remittance).Error; err != nil {
		return err
	}

	eligibleRemittanceStatuses := []string{
		basslink.RemittanceStatusSubmitted,
		basslink.RemittanceStatusWait,
		basslink.RemittanceStatusPaymentConfirmed,
	}

	if !slices.Contains(eligibleRemittanceStatuses, remittance.Status) {
		return errors.New("invalid status")
	}

	now := time.Now().Unix()

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.Remittance{}).Where("id = ? AND status IN ?", remittance.Id, eligibleRemittanceStatuses).Updates(map[string]interface{}{
			"agent_id":        agent.Id,
			"status":          basslink.RemittanceStatusCancelled,
			"processed_at":    now,
			"processed_by":    agent.Id,
			"processed_notes": req.Reason,
			"updated":         now,
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(basslink.RemittancePayment{}).Where("id = ? AND status IN ?", remittance.Id, []string{
			basslink.PaymentStatusWait,
			basslink.PaymentStatusConfirmed,
		}).Updates(map[string]interface{}{
			"status":  basslink.PaymentStatusFailed,
			"notes":   req.Reason,
			"updated": now,
		}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	if remittance.NotificationEmail != nil && len(*remittance.NotificationEmail) > 0 {
		s.App.EmailMsgChannel <- &basslink.EmailNotificationMesage{
			To:       *remittance.NotificationEmail,
			Subject:  fmt.Sprintf("%s - %s", remittance.Id, "Permintaan pengiriman dana telah dibatalkan"),
			Template: "remittance-cancel:1",
			Data: map[string]interface{}{
				"id":             remittance.Id,
				"sender_name":    remittance.FromName,
				"recipient_name": remittance.ToName,
				"to_currency":    remittance.TargetCurrency.Symbol,
				"to_amount":      s.App.FormatCurrency(fmt.Sprintf("%f", remittance.ToAmount)),
			},
		}
	}

	return nil
}

func (s *Service) completeRemittance(agent *basslink.Agent, id string, req *RemittanceCompleteRequest) error {
	var remittance basslink.Remittance

	if err := s.App.DB.Connection.Preload("SourceCurrency").Preload("TargetCurrency").Preload("Attachments").Where("id = ?", id).First(&remittance).Error; err != nil {
		return err
	}

	eligibleRemittanceStatuses := []string{
		basslink.RemittanceStatusWait,
		basslink.RemittanceStatusPaymentConfirmed,
	}

	if !slices.Contains(eligibleRemittanceStatuses, remittance.Status) {
		return errors.New("invalid status")
	}

	now := time.Now().Unix()

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.Remittance{}).Where("id = ? AND status IN ?", remittance.Id, eligibleRemittanceStatuses).Updates(map[string]interface{}{
			"agent_id":            agent.Id,
			"status":              basslink.RemittanceStatusCompleted,
			"processed_at":        now,
			"processed_by":        agent.Id,
			"processed_reference": req.Reference,
			"processed_notes":     req.Notes,
			"updated":             now,
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(basslink.RemittancePayment{}).Where("id = ? AND status IN ?", remittance.Id, []string{
			basslink.PaymentStatusWait,
			basslink.PaymentStatusConfirmed,
		}).Updates(map[string]interface{}{
			"status":  basslink.PaymentStatusCompleted,
			"updated": now,
		}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	if remittance.NotificationEmail != nil && len(*remittance.NotificationEmail) > 0 {
		s.App.EmailMsgChannel <- &basslink.EmailNotificationMesage{
			To:       *remittance.NotificationEmail,
			Subject:  fmt.Sprintf("%s - %s", remittance.Id, "Pengiriman dana telah berhasil diproses"),
			Template: "remittance-done:1",
			Data: map[string]interface{}{
				"id":             remittance.Id,
				"sender_name":    remittance.FromName,
				"recipient_name": remittance.ToName,
				"to_currency":    remittance.TargetCurrency.Symbol,
				"to_amount":      s.App.FormatCurrency(fmt.Sprintf("%f", remittance.ToAmount)),
			},
		}
	}

	return nil
}
