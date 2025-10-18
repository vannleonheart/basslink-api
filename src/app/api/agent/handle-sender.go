package agent

import (
	"CRM/src/lib/basslink"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Service) getUsers() (*[]basslink.Sender, error) {
	var senders []basslink.Sender

	if err := s.App.DB.Connection.Find(&senders).Error; err != nil {
		return nil, err
	}

	return &senders, nil
}

func (s *Service) getUser(senderId string) (*basslink.Sender, error) {
	var sender basslink.Sender

	if err := s.App.DB.Connection.Preload("Documents").Where("id = ?", senderId).First(&sender).Error; err != nil {
		return nil, err
	}

	return &sender, nil
}

func (s *Service) updateUser(agent *basslink.AgentUser, senderId string, req *UpdateSenderRequest) error {
	var selectedSender basslink.Sender

	if err := s.App.DB.Connection.Where("id = ?", senderId).First(&selectedSender).Error; err != nil {
		return err
	}

	now := time.Now().Unix()

	updatedUserData := map[string]interface{}{
		"sender_type":         req.SenderType,
		"name":                req.SenderName,
		"gender":              req.SenderGender,
		"birthdate":           req.SenderBirthdate,
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

	var documents []basslink.SenderDocument

	if req.SenderDocuments != nil && len(*req.SenderDocuments) > 0 {
		for _, document := range *req.SenderDocuments {
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

			documents = append(documents, basslink.SenderDocument{
				Id:           documentId,
				SenderId:     selectedSender.Id,
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
		if err := tx.Model(basslink.Sender{}).Where("id = ?", selectedSender.Id).Updates(updatedUserData).Error; err != nil {
			return err
		}

		if err := tx.Model(basslink.SenderDocument{}).Where("sender_id = ?", selectedSender.Id).Delete(nil).Error; err != nil {
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

func (s *Service) createUser(agent *basslink.Agent, req *CreateSenderRequest) error {
	newUserId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	newSender := basslink.Sender{
		Id:                newUserId.String(),
		SenderType:        req.SenderType,
		Name:              req.SenderName,
		Gender:            req.SenderGender,
		Birthdate:         req.SenderBirthdate,
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
		ZipCode:           req.SenderAddress,
		Contact:           req.SenderContact,
		Occupation:        req.SenderOccupation,
		PepStatus:         req.SenderPepStatus,
		Notes:             req.SenderNotes,
		Created:           now,
		CreatedBy:         agent.Id,
		Updated:           nil,
		UpdatedBy:         nil,
	}

	var documents []basslink.SenderDocument

	if req.SenderDocuments != nil && len(*req.SenderDocuments) > 0 {
		for _, document := range *req.SenderDocuments {
			if document.DocumentData == nil || document.DocumentType == nil || len(*document.DocumentData) == 0 || len(*document.DocumentType) == 0 {
				continue
			}

			if newDocumentId, e := uuid.NewV7(); e == nil {
				documents = append(documents, basslink.SenderDocument{
					Id:           newDocumentId.String(),
					SenderId:     newSender.Id,
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
		if err = tx.Create(&newSender).Error; err != nil {
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
