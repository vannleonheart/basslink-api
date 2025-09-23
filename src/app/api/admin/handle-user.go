package admin

import (
	"CRM/src/lib/basslink"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Service) getUsers() (*[]basslink.Administrator, error) {
	var users []basslink.Administrator

	if err := s.App.DB.Connection.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (s *Service) getUser(userId string) (*basslink.Administrator, error) {
	var user basslink.Administrator

	if err := s.App.DB.Connection.First(&user, userId).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) updateUser(userId string, req *UpdateUserRequest) error {
	var selectedUser basslink.Administrator

	if err := s.App.DB.Connection.First(&selectedUser, userId).Error; err != nil {
		return err
	}

	if selectedUser.Role == basslink.AdministratorRoleRoot {
		return errors.New("cannot update root user")
	}

	if req.Role == basslink.AdministratorRoleRoot {
		return errors.New("cannot set role as root")
	}

	if selectedUser.Username != req.Username {
		var existingUsers []basslink.Administrator

		if err := s.App.DB.Connection.Where("username = ?", req.Username).Limit(1).Find(&existingUsers).Error; err != nil {
			return err
		}

		if len(existingUsers) > 0 {
			return errors.New("username already exist")
		}
	}

	now := time.Now().Unix()

	updatedUserData := map[string]interface{}{
		"role":       req.Role,
		"username":   req.Username,
		"name":       req.Name,
		"email":      req.Email,
		"phone_code": req.PhoneCode,
		"phone_no":   req.PhoneNo,
		"updated":    now,
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.Administrator{}).Where("id = ?", selectedUser.Id).Updates(updatedUserData).Error; err != nil {
			return err
		}

		if req.Password != nil && *req.Password != "" {
			hasher := sha256.New()
			hasher.Write([]byte(*req.Password))
			hashedPassword := hasher.Sum(nil)
			updatedCredentialData := map[string]interface{}{
				"credential_data": fmt.Sprintf("%x", hashedPassword),
				"updated":         now,
			}

			if err := tx.Model(basslink.AdministratorCredential{}).Where("administrator_id = ?", selectedUser.Id).Updates(updatedCredentialData).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) deleteUser(userId string) error {
	var selectedUser basslink.Administrator

	if err := s.App.DB.Connection.First(&selectedUser, userId).Error; err != nil {
		return err
	}

	if selectedUser.Role == basslink.AdministratorRoleRoot {
		return errors.New("cannot delete root user")
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.AdministratorCredential{}).Where("administrator_id = ?", selectedUser.Id).Delete(nil).Error; err != nil {
			return err
		}

		if err := tx.Model(basslink.Administrator{}).Where("id = ?", selectedUser.Id).Delete(nil).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) createUser(req *CreateUserRequest) error {
	var existingUsers []basslink.Administrator

	if err := s.App.DB.Connection.Where("username = ?", req.Username).Limit(1).Find(&existingUsers).Error; err != nil {
		return err
	}

	if len(existingUsers) > 0 {
		return errors.New("username already exist")
	}

	if req.Role == basslink.AdministratorRoleRoot {
		return errors.New("cannot set role as root")
	}

	newUserId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	newUser := basslink.Administrator{
		Id:        newUserId.String(),
		Role:      req.Role,
		Username:  req.Username,
		Name:      req.Name,
		Email:     req.Email,
		PhoneCode: req.PhoneCode,
		PhoneNo:   req.PhoneNo,
		IsEnable:  true,
		Created:   now,
		Updated:   nil,
	}

	hasher := sha256.New()
	hasher.Write([]byte(req.Password))
	hashedPassword := hasher.Sum(nil)
	newUserCredential := basslink.AdministratorCredential{
		AdministratorId: newUser.Id,
		CredentialType:  "password",
		CredentialData:  fmt.Sprintf("%x", hashedPassword),
		Updated:         now,
	}

	if err = s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&newUser).Error; err != nil {
			return err
		}

		if err = tx.Create(&newUserCredential).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) toggleUserEnable(userId string) error {
	var selectedUser basslink.Administrator

	if err := s.App.DB.Connection.First(&selectedUser, userId).Error; err != nil {
		return err
	}

	if selectedUser.Role == basslink.AdministratorRoleRoot {
		return errors.New("cannot toggle root user")
	}

	newUserEnableValue := !selectedUser.IsEnable

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		return tx.Model(basslink.Administrator{}).Where("id = ?", selectedUser.Id).Update("is_enable", newUserEnableValue).Error
	}); err != nil {
		return err
	}

	return nil
}
