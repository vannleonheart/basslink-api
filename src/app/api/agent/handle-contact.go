package agent

import (
	"CRM/src/lib/basslink"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Service) getContacts() (*[]basslink.Contact, error) {
	var contacts []basslink.Contact

	if err := s.App.DB.Connection.Find(&contacts).Error; err != nil {
		return nil, err
	}

	return &contacts, nil
}

func (s *Service) getContact(contactId string) (*basslink.Contact, error) {
	var contact basslink.Contact

	if err := s.App.DB.Connection.Preload("Documents", "Accounts").Where("id = ?", contactId).First(&contact).Error; err != nil {
		return nil, err
	}

	return &contact, nil
}

func (s *Service) updateContact(contactId string, req *UpdateContactRequest) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("id = ?", contactId).First(&selectedContact).Error; err != nil {
		return err
	}

	now := time.Now().Unix()

	if req.ContactPhoneCode != nil {
		phoneCode := s.App.FormatPhoneCode(*req.ContactPhoneCode)
		req.ContactPhoneCode = &phoneCode
	}

	if req.ContactPhoneNo != nil {
		phoneNo := *req.ContactPhoneNo
		if phoneNo[0] == '0' {
			phoneNo = phoneNo[1:]
		}
		req.ContactPhoneNo = &phoneNo
	}

	updatedContactData := map[string]interface{}{
		"name":          req.ContactName,
		"birthdate":     req.ContactBirthdate,
		"gender":        req.ContactGender,
		"country":       req.ContactCountry,
		"region":        req.ContactRegion,
		"city":          req.ContactCity,
		"address":       req.ContactAddress,
		"email":         req.ContactEmail,
		"phone_code":    req.ContactPhoneCode,
		"phone_no":      req.ContactPhoneNo,
		"identity_type": req.ContactIdentityType,
		"identity_no":   req.ContactIdentityNo,
		"occupation":    req.ContactOccupation,
		"notes":         req.ContactNotes,
		"updated":       now,
	}

	var documents []basslink.ContactDocument
	var accounts []basslink.ContactAccount

	if req.ContactDocuments != nil && len(*req.ContactDocuments) > 0 {
		for _, document := range *req.ContactDocuments {
			if newDocumentId, e := uuid.NewV7(); e == nil {
				documents = append(documents, basslink.ContactDocument{
					Id:           newDocumentId.String(),
					ContactId:    selectedContact.Id,
					DocumentType: document.DocumentType,
					DocumentData: document.DocumentData,
					Notes:        document.Notes,
					IsVerified:   false,
					Created:      now,
				})
			}
		}
	}

	if req.ContactAccounts != nil && len(*req.ContactAccounts) > 0 {
		for _, account := range *req.ContactAccounts {
			if newAccountId, e := uuid.NewV7(); e == nil {
				if account.BankPhoneCode != nil {
					phoneCode := s.App.FormatPhoneCode(*account.BankPhoneCode)
					account.BankPhoneCode = &phoneCode
				}
				if account.BankPhoneNo != nil {
					phoneNo := *account.BankPhoneNo
					if phoneNo[0] == '0' {
						phoneNo = phoneNo[1:]
					}
					account.BankPhoneNo = &phoneNo
				}
				accounts = append(accounts, basslink.ContactAccount{
					Id:          newAccountId.String(),
					ContactId:   selectedContact.Id,
					AccountType: "",
					BankId:      "",
					BankName:    account.BankName,
					BankCode:    account.BankCode,
					BankSwift:   account.BankSwiftCode,
					OwnerName:   account.BankAccountName,
					No:          account.BankAccountNo,
					Country:     &account.BankCountry,
					Address:     account.BankAddress,
					Email:       account.BankEmail,
					Website:     account.BankWebsite,
					PhoneCode:   account.BankPhoneCode,
					PhoneNo:     account.BankPhoneNo,
					Notes:       account.BankNotes,
					Created:     now,
					Updated:     nil,
				})
			}
		}
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.Contact{}).Where("id = ?", selectedContact.Id).Updates(updatedContactData).Error; err != nil {
			return err
		}

		if len(documents) > 0 {
			if err := tx.CreateInBatches(&documents, len(documents)).Error; err != nil {
				return err
			}
		}

		if len(accounts) > 0 {
			if err := tx.CreateInBatches(&accounts, len(accounts)).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) deleteContact(contactId string) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("id = ?", contactId).First(&selectedContact).Error; err != nil {
		return err
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.ContactDocument{}).Where("contact_id = ?", selectedContact.Id).Delete(nil).Error; err != nil {
			return err
		}

		if err := tx.Model(basslink.ContactAccount{}).Where("contact_id = ?", selectedContact.Id).Delete(nil).Error; err != nil {
			return err
		}

		if err := tx.Model(basslink.Contact{}).Where("id = ?", selectedContact.Id).Delete(nil).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) createContact(agent *basslink.Agent, req *CreateContactRequest) error {
	newContactId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	if req.ContactPhoneCode != nil {
		phoneCode := s.App.FormatPhoneCode(*req.ContactPhoneCode)
		req.ContactPhoneCode = &phoneCode
	}

	if req.ContactPhoneNo != nil {
		phoneNo := *req.ContactPhoneNo
		if phoneNo[0] == '0' {
			phoneNo = phoneNo[1:]
		}
		req.ContactPhoneNo = &phoneNo
	}

	newContact := basslink.Contact{
		Id:           newContactId.String(),
		AgentId:      agent.Id,
		ContactType:  req.ContactType,
		Name:         req.ContactName,
		Gender:       req.ContactGender,
		Birthdate:    req.ContactBirthdate,
		Citizenship:  req.ContactCitizenship,
		IdentityType: req.ContactIdentityType,
		IdentityNo:   req.ContactIdentityNo,
		Country:      req.ContactCountry,
		Region:       req.ContactRegion,
		City:         req.ContactCity,
		Address:      req.ContactAddress,
		Email:        req.ContactEmail,
		PhoneCode:    req.ContactPhoneCode,
		PhoneNo:      req.ContactPhoneNo,
		Occupation:   req.ContactOccupation,
		Notes:        req.ContactNotes,
		IsVerified:   false,
		Created:      now,
		Updated:      nil,
	}

	var documents []basslink.ContactDocument
	var accounts []basslink.ContactAccount

	if req.ContactDocuments != nil && len(*req.ContactDocuments) > 0 {
		for _, document := range *req.ContactDocuments {
			if newDocumentId, e := uuid.NewV7(); e == nil {
				documents = append(documents, basslink.ContactDocument{
					Id:           newDocumentId.String(),
					ContactId:    newContact.Id,
					DocumentType: document.DocumentType,
					DocumentData: document.DocumentData,
					Notes:        document.Notes,
					IsVerified:   false,
					Created:      now,
				})
			}
		}
	}

	if req.ContactAccounts != nil && len(*req.ContactAccounts) > 0 {
		for _, account := range *req.ContactAccounts {
			if newAccountId, e := uuid.NewV7(); e == nil {
				if account.BankPhoneCode != nil {
					phoneCode := s.App.FormatPhoneCode(*account.BankPhoneCode)
					account.BankPhoneCode = &phoneCode
				}
				if account.BankPhoneNo != nil {
					phoneNo := *account.BankPhoneNo
					if phoneNo[0] == '0' {
						phoneNo = phoneNo[1:]
					}
					account.BankPhoneNo = &phoneNo
				}
				accounts = append(accounts, basslink.ContactAccount{
					Id:          newAccountId.String(),
					ContactId:   newContact.Id,
					AccountType: "",
					BankId:      "",
					BankName:    account.BankName,
					BankCode:    account.BankCode,
					BankSwift:   account.BankSwiftCode,
					OwnerName:   account.BankAccountName,
					No:          account.BankAccountNo,
					Country:     &account.BankCountry,
					Address:     account.BankAddress,
					Email:       account.BankEmail,
					Website:     account.BankWebsite,
					PhoneCode:   account.BankPhoneCode,
					PhoneNo:     account.BankPhoneNo,
					Notes:       account.BankNotes,
					Created:     now,
					Updated:     nil,
				})
			}
		}
	}

	if err = s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&newContact).Error; err != nil {
			return err
		}

		if len(documents) > 0 {
			if err = tx.CreateInBatches(&documents, len(documents)).Error; err != nil {
				return err
			}
		}

		if len(accounts) > 0 {
			if err = tx.CreateInBatches(&accounts, len(accounts)).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) createContactDocument(contactId string, req *CreateContactDocumentRequest) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("id = ?", contactId).First(&selectedContact).Error; err != nil {
		return err
	}

	documentId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	newContactDocument := basslink.ContactDocument{
		Id:           documentId.String(),
		ContactId:    selectedContact.Id,
		DocumentType: req.DocumentType,
		DocumentData: req.DocumentData,
		Notes:        req.Notes,
		IsVerified:   false,
		Created:      now,
	}

	if err = s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&newContactDocument).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) updateContactDocument(contactId, documentId string, req *UpdateContactDocumentRequest) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("id = ?", contactId).First(&selectedContact).Error; err != nil {
		return err
	}

	var selectedDocument basslink.ContactDocument

	if err := s.App.DB.Connection.Where("id = ? AND contact_id = ?", documentId, selectedContact.Id).First(&selectedDocument).Error; err != nil {
		return err
	}

	now := time.Now().Unix()

	updatedContactDocumentData := map[string]interface{}{
		"document_type": req.DocumentType,
		"document_data": req.DocumentData,
		"notes":         req.Notes,
		"updated":       now,
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.ContactDocument{}).Where("id = ? AND contact_id = ?", selectedDocument.Id, selectedContact.Id).Updates(&updatedContactDocumentData).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) deleteContactDocument(contactId, documentId string) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("id = ?", contactId).First(&selectedContact).Error; err != nil {
		return err
	}

	var selectedDocument basslink.ContactDocument

	if err := s.App.DB.Connection.Where("id = ? AND contact_id = ?", documentId, selectedContact.Id).First(&selectedDocument).Error; err != nil {
		return err
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.ContactDocument{}).Where("id = ? AND contact_id = ?", selectedDocument.Id, selectedContact.Id).Delete(nil).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) getContactAccounts(contactId string) (*[]basslink.ContactAccount, error) {
	var accounts []basslink.ContactAccount

	if err := s.App.DB.Connection.Where("contact_id = ?", contactId).Find(&accounts).Error; err != nil {
		return nil, err
	}

	return &accounts, nil
}

func (s *Service) createContactAccount(contactId string, req *CreateContactAccountRequest) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("id = ?", contactId).First(&selectedContact).Error; err != nil {
		return err
	}

	accountId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	newContactAccount := basslink.ContactAccount{
		Id:          accountId.String(),
		ContactId:   selectedContact.Id,
		AccountType: req.AccountType,
		BankId:      req.BankId,
		BankName:    req.BankName,
		BankCode:    req.BankCode,
		BankSwift:   req.BankSwift,
		OwnerName:   req.OwnerName,
		No:          req.No,
		Country:     req.Country,
		Address:     req.Address,
		Email:       req.Email,
		Website:     req.Website,
		PhoneCode:   req.PhoneCode,
		PhoneNo:     req.PhoneNo,
		Notes:       req.Notes,
		Created:     now,
	}

	if err = s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&newContactAccount).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) updateContactAccount(contactId, accountId string, req *UpdateContactAccountRequest) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("id = ?", contactId).First(&selectedContact).Error; err != nil {
		return err
	}

	var selectedAccount basslink.ContactAccount

	if err := s.App.DB.Connection.Where("id = ? AND contact_id = ?", accountId, selectedContact.Id).First(&selectedAccount).Error; err != nil {
		return err
	}

	now := time.Now().Unix()

	updatedContactAccountData := map[string]interface{}{
		"account_type": req.AccountType,
		"bank_id":      req.BankId,
		"bank_name":    req.BankName,
		"bank_code":    req.BankCode,
		"bank_swift":   req.BankSwift,
		"owner_name":   req.OwnerName,
		"no":           req.No,
		"country":      req.Country,
		"region":       req.Region,
		"city":         req.City,
		"address":      req.Address,
		"email":        req.Email,
		"website":      req.Website,
		"phone_code":   req.PhoneCode,
		"phone_no":     req.PhoneNo,
		"notes":        req.Notes,
		"updated":      now,
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.ContactAccount{}).Where("id = ? AND contact_id = ?", selectedAccount.Id, selectedContact.Id).Updates(&updatedContactAccountData).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) deleteContactAccount(contactId, accountId string) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("id = ?", contactId).First(&selectedContact).Error; err != nil {
		return err
	}

	var selectedAccount basslink.ContactAccount

	if err := s.App.DB.Connection.Where("id = ? AND contact_id = ?", accountId, selectedContact.Id).First(&selectedAccount).Error; err != nil {
		return err
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.ContactAccount{}).Where("id = ? AND contact_id = ?", selectedAccount.Id, selectedContact.Id).Delete(nil).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
