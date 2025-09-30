package agent

import (
	"CRM/src/lib/basslink"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Service) getUsers() (*[]basslink.User, error) {
	var users []basslink.User

	if err := s.App.DB.Connection.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (s *Service) getUser(userId string) (*basslink.User, error) {
	var user basslink.User

	if err := s.App.DB.Connection.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) updateUser(userId string, req *UpdateUserRequest) error {
	var selectedUser basslink.User

	if err := s.App.DB.Connection.Where("id = ?", userId).First(&selectedUser).Error; err != nil {
		return err
	}

	if selectedUser.Username != req.Username {
		var existingUsers []basslink.User

		if err := s.App.DB.Connection.Where("username = ?", req.Username).Limit(1).Find(&existingUsers).Error; err != nil {
			return err
		}

		if len(existingUsers) > 0 {
			return errors.New("username already exist")
		}
	}

	now := time.Now().Unix()

	updatedUserData := map[string]interface{}{
		"username":      req.Username,
		"user_type":     req.CustomerType,
		"name":          req.CustomerName,
		"gender":        req.CustomerGender,
		"birthdate":     req.CustomerBirthdate,
		"country":       req.CustomerCountry,
		"region":        req.CustomerRegion,
		"city":          req.CustomerCity,
		"address":       req.CustomerAddress,
		"email":         req.CustomerEmail,
		"phone_code":    req.CustomerPhoneCode,
		"phone_no":      req.CustomerPhoneNo,
		"identity_type": req.CustomerIdentityType,
		"identity_no":   req.CustomerIdentityNo,
		"occupation":    req.CustomerOccupation,
		"notes":         req.CustomerNotes,
		"updated":       now,
	}

	var documents []basslink.UserDocument

	if req.CustomerDocuments != nil && len(*req.CustomerDocuments) > 0 {
		for _, document := range *req.CustomerDocuments {
			if newDocumentId, e := uuid.NewV7(); e == nil {
				documents = append(documents, basslink.UserDocument{
					Id:           newDocumentId.String(),
					UserId:       selectedUser.Id,
					DocumentType: document.DocumentType,
					DocumentData: document.DocumentData,
					Notes:        document.Notes,
					Created:      now,
				})
			}
		}
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.User{}).Where("id = ?", selectedUser.Id).Updates(updatedUserData).Error; err != nil {
			return err
		}

		if len(documents) > 0 {
			if err := tx.CreateInBatches(&documents, len(documents)).Error; err != nil {
				return err
			}
		}

		if req.Password != nil && *req.Password != "" {
			updatedCredentialData := map[string]interface{}{
				"credential_data": s.App.HashPassword(*req.Password),
				"updated":         now,
			}

			if err := tx.Model(basslink.UserCredential{}).Where("user_id = ?", selectedUser.Id).Updates(updatedCredentialData).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) createUser(agent *basslink.Agent, req *CreateUserRequest) error {
	if req.Username != nil {
		var existingUsers []basslink.User

		if err := s.App.DB.Connection.Where("username = ?", req.Username).Limit(1).Find(&existingUsers).Error; err != nil {
			return err
		}

		if len(existingUsers) > 0 {
			return errors.New("username already exist")
		}
	}

	newUserId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	newUser := basslink.User{
		Id:            newUserId.String(),
		AgentId:       agent.Id,
		Username:      req.Username,
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
		IsEnable:      true,
		Created:       now,
		Updated:       nil,
	}

	var documents []basslink.UserDocument

	if req.CustomerDocuments != nil && len(*req.CustomerDocuments) > 0 {
		for _, document := range *req.CustomerDocuments {
			if newDocumentId, e := uuid.NewV7(); e == nil {
				documents = append(documents, basslink.UserDocument{
					Id:           newDocumentId.String(),
					UserId:       newUser.Id,
					DocumentType: document.DocumentType,
					DocumentData: document.DocumentData,
					Notes:        document.Notes,
					Created:      now,
				})
			}
		}
	}

	if err = s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&newUser).Error; err != nil {
			return err
		}

		if len(documents) > 0 {
			if err = tx.CreateInBatches(&documents, len(documents)).Error; err != nil {
				return err
			}
		}

		if req.Password != nil && len(*req.Password) > 0 {
			newUserCredential := basslink.UserCredential{
				UserId:         newUser.Id,
				CredentialType: "password",
				CredentialData: s.App.HashPassword(*req.Password),
				Updated:        now,
			}
			if err = tx.Create(&newUserCredential).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) toggleUserEnable(userId string) error {
	var selectedUser basslink.User

	if err := s.App.DB.Connection.Where("id = ?", userId).First(&selectedUser).Error; err != nil {
		return err
	}

	newUserEnableValue := !selectedUser.IsEnable

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		return tx.Model(basslink.User{}).Where("id = ?", selectedUser.Id).Update("is_enable", newUserEnableValue).Error
	}); err != nil {
		return err
	}

	return nil
}
