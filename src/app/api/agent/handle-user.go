package agent

import (
	"CRM/src/lib/basslink"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Service) getUsers(agent *basslink.Agent) (*[]basslink.User, error) {
	var users []basslink.User

	if err := s.App.DB.Connection.Where("agent_id = ?", agent.Id).Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (s *Service) getUser(agent *basslink.Agent, userId string) (*basslink.User, error) {
	var user basslink.User

	if err := s.App.DB.Connection.Where("agent_id = ?", agent.Id).First(&user, userId).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) updateUser(agent *basslink.Agent, userId string, req *UpdateUserRequest) error {
	var selectedUser basslink.User

	if err := s.App.DB.Connection.Where("agent_id = ?", agent.Id).First(&selectedUser, userId).Error; err != nil {
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
		"username":       req.Username,
		"name":           req.Name,
		"gender":         req.Gender,
		"birthdate":      req.Birthdate,
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
		"portrait_image": req.PortraitImage,
		"identity_image": req.IdentityImage,
		"notes":          req.Notes,
		"updated":        now,
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.User{}).Where("id = ? AND agent_id", selectedUser.Id, agent.Id).Updates(updatedUserData).Error; err != nil {
			return err
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
	var existingUsers []basslink.User

	if err := s.App.DB.Connection.Where("username = ?", req.Username).Limit(1).Find(&existingUsers).Error; err != nil {
		return err
	}

	if len(existingUsers) > 0 {
		return errors.New("username already exist")
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
		Name:          req.Name,
		Gender:        req.Gender,
		Birthdate:     req.Birthdate,
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
		PortraitImage: req.PortraitImage,
		IdentityImage: req.IdentityImage,
		Notes:         req.Notes,
		IsVerified:    false,
		EmailVerified: false,
		PhoneVerified: false,
		IsEnable:      true,
		Created:       now,
	}

	newUserCredential := basslink.UserCredential{
		UserId:         newUser.Id,
		CredentialType: "password",
		CredentialData: s.App.HashPassword(req.Password),
		Updated:        now,
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

func (s *Service) toggleUserEnable(agent *basslink.Agent, userId string) error {
	var selectedUser basslink.User

	if err := s.App.DB.Connection.Where("agent_id = ?", agent.Id).First(&selectedUser, userId).Error; err != nil {
		return err
	}

	newUserEnableValue := !selectedUser.IsEnable

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		return tx.Model(basslink.User{}).Where("id = ? AND agent_id = ?", selectedUser.Id, agent.Id).Update("is_enable", newUserEnableValue).Error
	}); err != nil {
		return err
	}

	return nil
}
