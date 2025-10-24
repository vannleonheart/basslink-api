package common

import (
	"CRM/src/lib/basslink"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/vannleonheart/goutil"
	"gorm.io/gorm"
)

func (s *Service) handleSearchTransaction(req *TransactionSearchRequest) (*basslink.Remittance, error) {
	var remmitance basslink.Remittance

	if err := s.App.DB.Connection.Preload("SourceCurrency").Preload("TargetCurrency").Preload("Attachments").Where("id = ? AND UPPER(from_name) = ? AND UPPER(to_name) = ?", strings.ToUpper(req.TransactionId), strings.ToUpper(req.SenderName), strings.ToUpper(req.RecipientName)).First(&remmitance).Error; err != nil {
		return nil, err
	}

	return &remmitance, nil
}

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

func (s *Service) getBankInfo(currency string) *BankInfo {
	switch strings.ToLower(currency) {
	default:
		return &BankInfo{
			BankName:     "BNI",
			BankCode:     "009",
			SwiftCode:    "BNINIDJA",
			AccountNo:    "7000-120-555",
			AccountOwner: "PT Basslink Remitansi Global",
			Currency:     "IDR",
		}
	case "usd":
		return &BankInfo{
			BankName:     "BNI",
			BankCode:     "009",
			SwiftCode:    "BNINIDJA",
			AccountNo:    "7000-120-770",
			AccountOwner: "PT Basslink Remitansi Global",
			Currency:     "USD",
		}
	case "eur":
		return &BankInfo{
			BankName:     "BNI",
			BankCode:     "009",
			SwiftCode:    "BNINIDJA",
			AccountNo:    "7000-120-883",
			AccountOwner: "PT Basslink Remitansi Global",
			Currency:     "EUR",
		}
	case "cny":
		return &BankInfo{
			BankName:     "BNI",
			BankCode:     "009",
			SwiftCode:    "BNINIDJA",
			AccountNo:    "7000-120-667",
			AccountOwner: "PT Basslink Remitansi Global",
			Currency:     "CNY",
		}
	}
}

func (s *Service) handleCreateTransaction(req *CreateRemittanceRequest) (*basslink.Remittance, error) {
	now := time.Now().Unix()

	remittanceId, err := s.generateTransactionId(req.TransferType)
	if err != nil {
		return nil, err
	}

	flFromAmount, err := req.FromAmount.Float64()
	if err != nil {
		return nil, err
	}

	flToAmount, err := req.ToAmount.Float64()
	if err != nil {
		return nil, err
	}

	var sender *basslink.Sender
	var recipient *basslink.Recipient

	newUserId, e := uuid.NewV7()
	if e != nil {
		return nil, e
	}

	gender := ""
	if req.SenderGender != nil && len(*req.SenderGender) > 0 {
		gender = *req.SenderGender
	}

	birthdate := ""
	if req.SenderBirthdate != nil && len(*req.SenderBirthdate) > 0 {
		birthdate = *req.SenderBirthdate
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
		CreatedBy:         "",
	}

	newrecipientId, e := uuid.NewV7()
	if e != nil {
		return nil, e
	}

	bankCode := ""
	if req.RecipientBankCode != nil && len(*req.RecipientBankCode) > 0 {
		bankCode = *req.RecipientBankCode
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
		BankCode:         bankCode,
		BankAccountNo:    req.RecipientBankAccountNo,
		BankAccountOwner: req.RecipientBankAccountOwner,
		Notes:            req.RecipientNotes,
		Created:          now,
		Updated:          nil,
		Documents:        nil,
	}

	strToAmount := fmt.Sprintf("%f", flToAmount)
	calculate, err := s.handleGetRate(&GetRateRequest{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		FromAmount:   nil,
		ToAmount:     &strToAmount,
	})

	if err != nil {
		return nil, err
	}

	flRate, err := calculate.Rate.Float64()
	if err != nil {
		return nil, err
	}

	flFeePercent, err := calculate.FeePercent.Float64()
	if err != nil {
		return nil, err
	}

	flFeeFixed, err := calculate.FeeFixed.Float64()
	if err != nil {
		return nil, err
	}

	totalFee, err := calculate.TotalFee.Float64()
	if err != nil {
		return nil, err
	}

	newRemittance := basslink.Remittance{
		Id:                    remittanceId,
		AgentId:               "",
		SenderId:              sender.Id,
		FromCurrency:          req.FromCurrency,
		FromAmount:            flFromAmount,
		FromSenderType:        sender.SenderType,
		FromName:              sender.Name,
		FromGender:            sender.Gender,
		FromBirthdate:         sender.Birthdate,
		FromCitizenship:       sender.Citizenship,
		FromIdentityType:      sender.IdentityType,
		FromIdentityNo:        sender.IdentityNo,
		FromRegisteredCountry: sender.RegisteredCountry,
		FromRegisteredRegion:  sender.RegisteredRegion,
		FromRegisteredCity:    sender.RegisteredCity,
		FromRegisteredAddress: sender.RegisteredAddress,
		FromRegisteredZipCode: sender.RegisteredZipCode,
		FromCountry:           sender.Country,
		FromRegion:            sender.Region,
		FromCity:              sender.City,
		FromAddress:           sender.Address,
		FromZipCode:           sender.ZipCode,
		FromContact:           sender.Contact,
		FromOccupation:        sender.Occupation,
		FromPepStatus:         sender.PepStatus,
		FromNotes:             sender.Notes,
		RecipientId:           recipient.Id,
		ToCurrency:            req.ToCurrency,
		ToAmount:              flToAmount,
		ToRecipientType:       recipient.RecipientType,
		ToRelationship:        recipient.Relationship,
		ToName:                recipient.Name,
		ToCountry:             recipient.Country,
		ToRegion:              recipient.Region,
		ToCity:                recipient.City,
		ToAddress:             recipient.Address,
		ToZipCode:             recipient.ZipCode,
		ToContact:             recipient.Contact,
		ToPepStatus:           recipient.PepStatus,
		ToAccountType:         recipient.AccountType,
		ToBankName:            recipient.BankName,
		ToBankCode:            recipient.BankCode,
		ToBankAccountNo:       recipient.BankAccountNo,
		ToBankAccountOwner:    recipient.BankAccountOwner,
		ToNotes:               recipient.Notes,
		RateCurrency:          req.FromCurrency,
		Rate:                  flRate,
		FeeCurrency:           req.FromCurrency,
		FeeAmountPercent:      flFeePercent,
		FeeAmountFixed:        flFeeFixed,
		FeeTotal:              totalFee,
		PaymentMethod:         "bank_transfer",
		TransferType:          req.TransferType,
		TransferRef:           req.TransferReference,
		FundSource:            req.FundSource,
		Purpose:               req.Purpose,
		Notes:                 req.Notes,
		Status:                basslink.RemittanceStatusSubmitted,
		IsSettled:             false,
		CreatedBy:             nil,
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
					SubmitBy:     "",
					SubmitTime:   now,
				})
			}
		}
	}

	if err = s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(sender).Error; err != nil {
			return err
		}

		if err = tx.Create(recipient).Error; err != nil {
			return err
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
		return nil, err
	}

	bankInfo := s.getBankInfo(newRemittance.FromCurrency)
	s.App.EmailMsgChannel <- &basslink.EmailNotificationMesage{
		To:       newRemittance.FromContact,
		Subject:  fmt.Sprintf("%s %s", newRemittance.Id, "Selesaikan pembayaran untuk melanjutkan proses kirim dana"),
		Template: "remittance-submitted:1",
		Data: map[string]interface{}{
			"id":                        newRemittance.Id,
			"sender_name":               newRemittance.FromName,
			"recipient_name":            newRemittance.ToName,
			"to_currency":               newRemittance.ToCurrency,
			"to_amount":                 newRemittance.ToAmount,
			"bank_name":                 bankInfo.BankName,
			"bank_code":                 bankInfo.BankCode,
			"bank_swift":                bankInfo.SwiftCode,
			"account_no":                bankInfo.AccountNo,
			"account_name":              bankInfo.AccountOwner,
			"currency":                  newRemittance.FromCurrency,
			"amount":                    newRemittance.FromAmount,
			"payment_confirmation_link": fmt.Sprintf("%s/%s", s.App.Config.PaymentConfirmationLink, newRemittance.Id),
		},
	}

	return &newRemittance, nil
}
