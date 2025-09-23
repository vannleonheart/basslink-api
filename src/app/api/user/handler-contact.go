package user

import (
	"CRM/src/lib/basslink"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Service) GetContacts(user *basslink.User) (*[]basslink.Contact, error) {
	var contacts []basslink.Contact

	if err := s.App.DB.Connection.Where("user_id = ?", user.Id).Find(&contacts).Error; err != nil {
		return nil, err
	}

	return &contacts, nil
}

func (s *Service) GetContact(user *basslink.User, contactId string) (*basslink.Contact, error) {
	var contact basslink.Contact

	if err := s.App.DB.Connection.Preload("Documents", "Accounts").Where("user_id = ?", user.Id).First(&contact, contactId).Error; err != nil {
		return nil, err
	}

	return &contact, nil
}

func (s *Service) UpdateContact(user *basslink.User, contactId string, req *UpdateContactRequest) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("user_id = ?", user.Id).First(&selectedContact, contactId).Error; err != nil {
		return err
	}

	now := time.Now().Unix()

	updatedContactData := map[string]interface{}{
		"name":           req.Name,
		"birthdate":      req.Birthdate,
		"gender":         req.Gender,
		"country":        req.Country,
		"region":         req.Region,
		"city":           req.City,
		"address":        req.Address,
		"email":          req.Email,
		"phone_code":     req.PhoneCode,
		"phone_no":       req.PhoneNo,
		"identity_type":  req.IdentityType,
		"identity_no":    req.IdentityNo,
		"occupation":     req.Occupation,
		"identity_image": req.IdentityImage,
		"portrait_image": req.PortraitImage,
		"notes":          req.Notes,
		"updated":        now,
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.Contact{}).Where("id = ? AND user_id", selectedContact.Id, user.Id).Updates(updatedContactData).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteContact(user *basslink.User, contactId string) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("user_id = ?", user.Id).First(&selectedContact, contactId).Error; err != nil {
		return err
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.ContactDocument{}).Where("contact_id = ?", selectedContact.Id).Delete(nil).Error; err != nil {
			return err
		}

		if err := tx.Model(basslink.ContactAccount{}).Where("contact_id = ?", selectedContact.Id).Delete(nil).Error; err != nil {
			return err
		}

		if err := tx.Model(basslink.Contact{}).Where("id = ? AND user_id = ?", selectedContact.Id, user.Id).Delete(nil).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateContact(user *basslink.User, req *CreateContactRequest) error {
	newContactId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	newUser := basslink.Contact{
		Id:            newContactId.String(),
		AgentId:       user.AgentId,
		UserId:        &user.Id,
		Name:          req.Name,
		Birthdate:     req.Birthdate,
		Gender:        req.Gender,
		Country:       req.Country,
		Region:        req.Region,
		City:          req.City,
		Address:       req.Address,
		Email:         req.Email,
		PhoneCode:     req.PhoneCode,
		PhoneNo:       req.PhoneNo,
		IdentityType:  req.IdentityType,
		IdentityNo:    req.IdentityNo,
		Occupation:    req.Occupation,
		IdentityImage: req.IdentityImage,
		PortraitImage: req.PortraitImage,
		Notes:         req.Notes,
		IsVerified:    false,
		Created:       now,
	}

	if err = s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&newUser).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateContactDocument(user *basslink.User, contactId string, req *CreateContactDocumentRequest) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("user_id = ?", user.Id).First(&selectedContact, contactId).Error; err != nil {
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

func (s *Service) UpdateContactDocument(user *basslink.User, contactId, documentId string, req *UpdateContactDocumentRequest) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("user_id = ?", user.Id).First(&selectedContact, contactId).Error; err != nil {
		return err
	}

	var selectedDocument basslink.ContactDocument

	if err := s.App.DB.Connection.Where("contact_id = ?", selectedContact.Id).First(&selectedDocument, documentId).Error; err != nil {
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

func (s *Service) DeleteContactDocument(user *basslink.User, contactId, documentId string) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("user_id = ?", user.Id).First(&selectedContact, contactId).Error; err != nil {
		return err
	}

	var selectedDocument basslink.ContactDocument

	if err := s.App.DB.Connection.Where("contact_id = ?", selectedContact.Id).First(&selectedDocument, documentId).Error; err != nil {
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

func (s *Service) CreateContactAccount(user *basslink.User, contactId string, req *CreateContactAccountRequest) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("user_id = ?", user.Id).First(&selectedContact, contactId).Error; err != nil {
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
		Region:      req.Region,
		City:        req.City,
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

func (s *Service) UpdateContactAccount(user *basslink.User, contactId, accountId string, req *UpdateContactAccountRequest) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("user_id = ?", user.Id).First(&selectedContact, contactId).Error; err != nil {
		return err
	}

	var selectedAccount basslink.ContactAccount

	if err := s.App.DB.Connection.Where("contact_id = ?", selectedContact.Id).First(&selectedAccount, accountId).Error; err != nil {
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

func (s *Service) DeleteContactAccount(user *basslink.User, contactId, accountId string) error {
	var selectedContact basslink.Contact

	if err := s.App.DB.Connection.Where("user_id = ?", user.Id).First(&selectedContact, contactId).Error; err != nil {
		return err
	}

	var selectedAccount basslink.ContactAccount

	if err := s.App.DB.Connection.Where("contact_id = ?", selectedContact.Id).First(&selectedAccount, accountId).Error; err != nil {
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
