package agent

import (
	"CRM/src/lib/basslink"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Service) getDisbursements(agent *basslink.Agent) (*[]basslink.Disbursement, error) {
	var disbursements []basslink.Disbursement

	if err := s.App.DB.Connection.Preload("User").Preload("Contact").Preload("TargetAccount").Preload("TargetCurrency").Where("agent_id = ?", agent.Id).Find(&disbursements).Error; err != nil {
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

	flFee, err := req.Fee.Float64()
	if err != nil {
		return err
	}

	if req.CustomerPhoneCode != nil {
		phoneCode := s.App.FormatPhoneCode(*req.CustomerPhoneCode)
		req.CustomerPhoneCode = &phoneCode
	}

	if req.CustomerPhoneNo != nil {
		phoneNo := *req.CustomerPhoneNo
		if phoneNo[0] == '0' {
			phoneNo = phoneNo[1:]
		}
		req.CustomerPhoneNo = &phoneNo
	}

	if req.BeneficiaryPhoneCode != nil {
		phoneCode := s.App.FormatPhoneCode(*req.BeneficiaryPhoneCode)
		req.BeneficiaryPhoneCode = &phoneCode
	}

	if req.BeneficiaryPhoneNo != nil {
		phoneNo := *req.BeneficiaryPhoneNo
		if phoneNo[0] == '0' {
			phoneNo = phoneNo[1:]
		}
		req.BeneficiaryPhoneNo = &phoneNo
	}

	if req.BankPhoneCode != nil {
		phoneCode := s.App.FormatPhoneCode(*req.BankPhoneCode)
		req.BankPhoneCode = &phoneCode
	}

	if req.BankPhoneNo != nil {
		phoneNo := *req.BankPhoneNo
		if phoneNo[0] == '0' {
			phoneNo = phoneNo[1:]
		}
		req.BankPhoneNo = &phoneNo
	}

	var user *basslink.User
	var contact *basslink.Contact
	var account *basslink.ContactAccount
	var updateUserData, updateContactData, updateAccountData *map[string]interface{}

	if req.FromCustomer != nil {
		var existingUsers basslink.User

		if err = s.App.DB.Connection.Where("id = ?", *req.FromCustomer).First(&existingUsers).Error; err != nil {
			return err
		}

		user = &existingUsers
		updateUserData = &map[string]interface{}{
			"user_type":     req.CustomerType,
			"name":          req.CustomerName,
			"gender":        req.CustomerGender,
			"birthdate":     req.CustomerBirthdate,
			"citizenship":   req.CustomerCitizenship,
			"identity_type": req.CustomerIdentityType,
			"identity_no":   req.CustomerIdentityNo,
			"country":       req.CustomerCountry,
			"region":        req.CustomerRegion,
			"city":          req.CustomerCity,
			"address":       req.CustomerAddress,
			"email":         req.CustomerEmail,
			"phone_code":    req.CustomerPhoneCode,
			"phone_no":      req.CustomerPhoneNo,
			"occupation":    req.CustomerOccupation,
			"notes":         req.CustomerNotes,
			"updated":       now,
		}
	} else {
		newUserId, e := uuid.NewV7()
		if e != nil {
			return e
		}

		user = &basslink.User{
			Id:            newUserId.String(),
			AgentId:       agent.Id,
			Username:      nil,
			UserType:      req.CustomerType,
			Name:          req.CustomerName,
			Gender:        req.CustomerGender,
			Birthdate:     req.CustomerBirthdate,
			Citizenship:   req.CustomerCitizenship,
			IdentityType:  req.CustomerIdentityType,
			IdentityNo:    req.CustomerIdentityNo,
			Country:       req.CustomerCountry,
			Region:        req.CustomerRegion,
			City:          req.CustomerCity,
			Address:       req.CustomerAddress,
			Email:         req.CustomerEmail,
			PhoneCode:     req.CustomerPhoneCode,
			PhoneNo:       req.CustomerPhoneNo,
			Occupation:    req.CustomerOccupation,
			Notes:         req.CustomerNotes,
			IsVerified:    false,
			EmailVerified: false,
			PhoneVerified: false,
			IsEnable:      false,
			Created:       now,
			Updated:       nil,
		}
	}

	if req.ToContact != nil {
		var existingContacts basslink.Contact

		if err = s.App.DB.Connection.Where("id = ?", *req.ToContact).First(&existingContacts).Error; err != nil {
			return err
		}

		contact = &existingContacts
		updateContactData = &map[string]interface{}{
			"contact_type":  req.BeneficiaryType,
			"name":          req.BeneficiaryName,
			"gender":        req.BeneficiaryGender,
			"birthdate":     req.BeneficiaryBirthdate,
			"citizenship":   req.BeneficiaryCitizenship,
			"identity_type": req.BeneficiaryIdentityType,
			"identity_no":   req.BeneficiaryIdentityNo,
			"country":       req.BeneficiaryCountry,
			"region":        req.BeneficiaryRegion,
			"city":          req.BeneficiaryCity,
			"address":       req.BeneficiaryAddress,
			"email":         req.BeneficiaryEmail,
			"phone_code":    req.BeneficiaryPhoneCode,
			"phone_no":      req.BeneficiaryPhoneNo,
			"occupation":    req.BeneficiaryOccupation,
			"notes":         req.BeneficiaryNotes,
			"updated":       now,
		}
	} else {
		newContactId, e := uuid.NewV7()
		if e != nil {
			return e
		}

		contact = &basslink.Contact{
			Id:           newContactId.String(),
			AgentId:      agent.Id,
			ContactType:  req.BeneficiaryType,
			Name:         req.BeneficiaryName,
			Gender:       req.BeneficiaryGender,
			Birthdate:    req.BeneficiaryBirthdate,
			Citizenship:  req.BeneficiaryCitizenship,
			IdentityType: req.BeneficiaryIdentityType,
			IdentityNo:   req.BeneficiaryIdentityNo,
			Country:      req.BeneficiaryCountry,
			Region:       req.BeneficiaryRegion,
			City:         req.BeneficiaryCity,
			Address:      req.BeneficiaryAddress,
			Email:        req.BeneficiaryEmail,
			PhoneCode:    req.BeneficiaryPhoneCode,
			PhoneNo:      req.BeneficiaryPhoneNo,
			Occupation:   req.BeneficiaryOccupation,
			Notes:        req.BeneficiaryNotes,
			IsVerified:   false,
			Created:      now,
			Updated:      nil,
		}
	}

	if req.ToAccount != nil {
		var existingAccounts basslink.ContactAccount

		if err = s.App.DB.Connection.Where("id = ?", *req.ToAccount).First(&existingAccounts).Error; err != nil {
			return err
		}

		account = &existingAccounts
		updateAccountData = &map[string]interface{}{
			"bank_name":  req.BankName,
			"bank_code":  req.BankCode,
			"bank_swift": req.BankSwiftCode,
			"owner_name": req.BankAccountName,
			"no":         req.BankAccountNo,
			"country":    &req.BankCountry,
			"address":    req.BankAddress,
			"email":      req.BankEmail,
			"website":    req.BankWebsite,
			"phone_code": req.BankPhoneCode,
			"phone_no":   req.BankPhoneNo,
			"notes":      req.BankNotes,
			"updated":    now,
		}
	} else {
		newAccountId, e := uuid.NewV7()
		if e != nil {
			return e
		}

		account = &basslink.ContactAccount{
			Id:          newAccountId.String(),
			ContactId:   contact.Id,
			AccountType: "",
			BankId:      "",
			BankName:    req.BankName,
			BankCode:    req.BankCode,
			BankSwift:   req.BankSwiftCode,
			OwnerName:   req.BankAccountName,
			No:          req.BankAccountNo,
			Country:     &req.BankCountry,
			Address:     req.BankAddress,
			Email:       req.BankEmail,
			Website:     req.BankWebsite,
			PhoneCode:   req.BankPhoneCode,
			PhoneNo:     req.BankPhoneNo,
			Notes:       req.BankNotes,
			Created:     now,
			Updated:     nil,
		}
	}

	newDisbursement := basslink.Disbursement{
		Id:           disbursementId.String(),
		AgentId:      agent.Id,
		UserId:       &user.Id,
		FromCurrency: req.FromCurrency,
		FromAmount:   flFromAmount,
		ToContact:    contact.Id,
		ToCurrency:   req.ToCurrency,
		ToAmount:     flToAmount,
		ToAccount:    account.Id,
		RateCurrency: req.FromCurrency,
		Rate:         flRate,
		FeeCurrency:  req.FromCurrency,
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
		if req.FromCustomer == nil {
			if err = tx.Create(user).Error; err != nil {
				return err
			}
		} else {
			if err = tx.Model(basslink.User{}).Where("id = ?", user.Id).Updates(*updateUserData).Error; err != nil {
				return err
			}
		}

		if req.ToContact == nil {
			if err = tx.Create(contact).Error; err != nil {
				return err
			}
		} else {
			if err = tx.Model(basslink.Contact{}).Where("id = ?", contact.Id).Updates(*updateContactData).Error; err != nil {
				return err
			}
		}

		if req.ToAccount == nil {
			if err = tx.Create(account).Error; err != nil {
				return err
			}
		} else {
			if err = tx.Model(basslink.ContactAccount{}).Where("id = ? AND contact_id = ?", account.Id, contact.Id).Updates(*updateAccountData).Error; err != nil {
				return err
			}
		}

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
