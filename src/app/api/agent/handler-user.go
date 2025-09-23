package agent

import (
	"CRM/src/lib/basslink"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Service) GetUsers(agent *basslink.Agent) (*[]basslink.AgentUser, error) {
	var users []basslink.AgentUser

	if err := s.App.DB.Connection.Where("agent_id = ?", agent.Id).Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (s *Service) GetUser(agent *basslink.Agent, userId string) (*basslink.AgentUser, error) {
	var user basslink.AgentUser

	if err := s.App.DB.Connection.Where("agent_id = ?", agent.Id).First(&user, userId).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) UpdateUser(agent *basslink.Agent, userId string, req *UpdateUserRequest) error {
	var selectedUser basslink.AgentUser

	if err := s.App.DB.Connection.Where("agent_id = ?", agent.Id).First(&selectedUser, userId).Error; err != nil {
		return err
	}

	if selectedUser.Role == basslink.AgentRoleOwner {
		return errors.New("cannot update owner user")
	}

	if req.Role == basslink.AgentRoleOwner {
		return errors.New("cannot set role as owner")
	}

	if selectedUser.Username != req.Username {
		var existingUsers []basslink.AgentUser

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
		if err := tx.Model(basslink.AgentUser{}).Where("id = ? AND agent_id", selectedUser.Id, agent.Id).Updates(updatedUserData).Error; err != nil {
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

			if err := tx.Model(basslink.AgentUserCredential{}).Where("agent_user_id = ?", selectedUser.Id).Updates(updatedCredentialData).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUser(agent *basslink.Agent, userId string) error {
	var selectedUser basslink.AgentUser

	if err := s.App.DB.Connection.Where("agent_id = ?", agent.Id).First(&selectedUser, userId).Error; err != nil {
		return err
	}

	if selectedUser.Role == basslink.AgentRoleOwner {
		return errors.New("cannot delete owner user")
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.AgentUserCredential{}).Where("agent_user_id = ?", selectedUser.Id).Delete(nil).Error; err != nil {
			return err
		}

		if err := tx.Model(basslink.AgentUser{}).Where("id = ? AND agent_id = ?", selectedUser.Id, agent.Id).Delete(nil).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateUser(agent *basslink.Agent, req *CreateUserRequest) error {
	var existingUsers []basslink.AgentUser

	if err := s.App.DB.Connection.Where("username = ?", req.Username).Limit(1).Find(&existingUsers).Error; err != nil {
		return err
	}

	if len(existingUsers) > 0 {
		return errors.New("username already exist")
	}

	if req.Role == basslink.AgentRoleOwner {
		return errors.New("cannot set role as root")
	}

	newUserId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	now := time.Now().Unix()

	newUser := basslink.AgentUser{
		Id:        newUserId.String(),
		AgentId:   agent.Id,
		Role:      req.Role,
		Username:  req.Username,
		Name:      req.Name,
		Email:     req.Email,
		PhoneCode: req.PhoneCode,
		PhoneNo:   req.PhoneNo,
		IsEnable:  true,
		Created:   now,
	}

	hasher := sha256.New()
	hasher.Write([]byte(req.Password))
	hashedPassword := hasher.Sum(nil)
	newUserCredential := basslink.AgentUserCredential{
		AgentUserId:    newUser.Id,
		CredentialType: "password",
		CredentialData: fmt.Sprintf("%x", hashedPassword),
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

func (s *Service) ToggleUserEnable(agent *basslink.Agent, userId string) error {
	var selectedUser basslink.AgentUser

	if err := s.App.DB.Connection.Where("agent_id = ?", agent.Id).First(&selectedUser, userId).Error; err != nil {
		return err
	}

	if selectedUser.Role == basslink.AgentRoleOwner {
		return errors.New("cannot toggle owner")
	}

	newUserEnableValue := !selectedUser.IsEnable

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		return tx.Model(basslink.AgentUser{}).Where("id = ? AND agent_id = ?", selectedUser.Id, agent.Id).Update("is_enable", newUserEnableValue).Error
	}); err != nil {
		return err
	}

	return nil
}
