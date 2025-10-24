package agent

import (
	"CRM/src/lib/basslink"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Service) getRecipients() (*[]basslink.Recipient, error) {
	var recipients []basslink.Recipient

	if err := s.App.DB.Connection.Find(&recipients).Error; err != nil {
		return nil, err
	}

	return &recipients, nil
}

func (s *Service) getRecipient(recipientId string) (*basslink.Recipient, error) {
	var recipient basslink.Recipient

	if err := s.App.DB.Connection.Preload("Documents").Where("id = ?", recipientId).First(&recipient).Error; err != nil {
		return nil, err
	}

	return &recipient, nil
}

func (s *Service) updateRecipient(recipientId string, req *UpdateRecipientRequest) error {
	var selectedRecipient basslink.Recipient

	if err := s.App.DB.Connection.Where("id = ?", recipientId).First(&selectedRecipient).Error; err != nil {
		return err
	}

	now := time.Now().Unix()

	updatedRecipientData := map[string]interface{}{
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

	var documents []basslink.RecipientDocument

	if req.RecipientDocuments != nil && len(*req.RecipientDocuments) > 0 {
		for _, document := range *req.RecipientDocuments {
			if document.DocumentData == nil || document.DocumentType == nil || len(*document.DocumentData) == 0 || len(*document.DocumentType) == 0 {
				continue
			}

			documentId := ""

			if document.Id != nil && len(*document.Id) > 0 {
				documentId = *document.Id
			} else {
				newDocumentId, e := uuid.NewV7()
				if e != nil {
					return e
				}
				documentId = newDocumentId.String()
			}

			documents = append(documents, basslink.RecipientDocument{
				Id:           documentId,
				RecipientId:  selectedRecipient.Id,
				DocumentType: *document.DocumentType,
				DocumentData: *document.DocumentData,
				Notes:        document.Notes,
				IsVerified:   false,
				Created:      now,
				Updated:      nil,
			})
		}
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.Recipient{}).Where("id = ?", selectedRecipient.Id).Updates(updatedRecipientData).Error; err != nil {
			return err
		}

		if err := tx.Model(basslink.RecipientDocument{}).Where("recipient_id = ?", selectedRecipient.Id).Delete(nil).Error; err != nil {
			return err
		}

		if len(documents) > 0 {
			if err := tx.CreateInBatches(&documents, len(documents)).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) deleteRecipient(recipientId string) error {
	var selectedRecipient basslink.Recipient

	if err := s.App.DB.Connection.Where("id = ?", recipientId).First(&selectedRecipient).Error; err != nil {
		return err
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.RecipientDocument{}).Where("recipient_id = ?", selectedRecipient.Id).Delete(nil).Error; err != nil {
			return err
		}

		if err := tx.Model(basslink.Recipient{}).Where("id = ?", selectedRecipient.Id).Delete(nil).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) createRecipient(req *CreateRecipientRequest) error {
	newRecipientId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	newRecipient := basslink.Recipient{
		Id:               newRecipientId.String(),
		SenderId:         req.RecipientSenderId,
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

	var documents []basslink.RecipientDocument

	if req.RecipientDocuments != nil && len(*req.RecipientDocuments) > 0 {
		for _, document := range *req.RecipientDocuments {
			if document.DocumentData == nil || document.DocumentType == nil || len(*document.DocumentData) == 0 || len(*document.DocumentType) == 0 {
				continue
			}

			if newDocumentId, e := uuid.NewV7(); e == nil {
				documents = append(documents, basslink.RecipientDocument{
					Id:           newDocumentId.String(),
					RecipientId:  newRecipient.Id,
					DocumentType: *document.DocumentType,
					DocumentData: *document.DocumentData,
					Notes:        document.Notes,
					IsVerified:   false,
					Created:      now,
					Updated:      nil,
				})
			}
		}
	}

	if err = s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&newRecipient).Error; err != nil {
			return err
		}

		if len(documents) > 0 {
			if err = tx.CreateInBatches(&documents, len(documents)).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) createRecipientDocument(recipientId string, req *CreateRecipientDocumentRequest) error {
	var selectedRecipient basslink.Recipient

	if err := s.App.DB.Connection.Where("id = ?", recipientId).First(&selectedRecipient).Error; err != nil {
		return err
	}

	documentId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	newRecipientDocument := basslink.RecipientDocument{
		Id:           documentId.String(),
		RecipientId:  selectedRecipient.Id,
		DocumentType: req.DocumentType,
		DocumentData: req.DocumentData,
		Notes:        req.Notes,
		IsVerified:   false,
		Created:      now,
		Updated:      nil,
	}

	if err = s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&newRecipientDocument).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) updateRecipientDocument(recipientId, documentId string, req *UpdateRecipientDocumentRequest) error {
	var selectedRecipient basslink.Recipient

	if err := s.App.DB.Connection.Where("id = ?", recipientId).First(&selectedRecipient).Error; err != nil {
		return err
	}

	var selectedDocument basslink.RecipientDocument

	if err := s.App.DB.Connection.Where("id = ? AND recipient_id = ?", documentId, selectedRecipient.Id).First(&selectedDocument).Error; err != nil {
		return err
	}

	now := time.Now().Unix()

	updatedRecipientDocumentData := map[string]interface{}{
		"document_type": req.DocumentType,
		"document_data": req.DocumentData,
		"notes":         req.Notes,
		"updated":       now,
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.RecipientDocument{}).Where("id = ? AND recipient_id = ?", selectedDocument.Id, selectedRecipient.Id).Updates(&updatedRecipientDocumentData).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) deleteRecipientDocument(recipientId, documentId string) error {
	var selectedRecipient basslink.Recipient

	if err := s.App.DB.Connection.Where("id = ?", recipientId).First(&selectedRecipient).Error; err != nil {
		return err
	}

	var selectedDocument basslink.RecipientDocument

	if err := s.App.DB.Connection.Where("id = ? AND recipient_id = ?", documentId, selectedRecipient.Id).First(&selectedDocument).Error; err != nil {
		return err
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.RecipientDocument{}).Where("id = ? AND recipient_id = ?", selectedDocument.Id, selectedRecipient.Id).Delete(nil).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
