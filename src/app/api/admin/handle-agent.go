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

func (s *Service) getAgents() (*[]basslink.Agent, error) {
	var agents []basslink.Agent

	if err := s.App.DB.Connection.Find(&agents).Error; err != nil {
		return nil, err
	}

	return &agents, nil
}

func (s *Service) getAgent(agentId string) (*basslink.Agent, error) {
	var agent basslink.Agent

	if err := s.App.DB.Connection.First(&agent, agentId).Error; err != nil {
		return nil, err
	}

	return &agent, nil
}

func (s *Service) updateAgent(agentId string, req *UpdateAgentRequest) error {
	var selectedAgent basslink.Agent

	if err := s.App.DB.Connection.First(&selectedAgent, agentId).Error; err != nil {
		return err
	}

	now := time.Now().Unix()

	updatedAgentData := map[string]interface{}{
		"name":       req.AgentName,
		"country":    req.Country,
		"region":     req.Region,
		"city":       req.City,
		"address":    req.Address,
		"phone_code": req.PhoneCode,
		"phone_no":   req.PhoneNo,
		"email":      req.Email,
		"website":    req.Website,
		"updated":    now,
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.Agent{}).Where("id = ?", selectedAgent.Id).Updates(updatedAgentData).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) createAgent(req *CreateAgentRequest) error {
	var existingAgentUsers []basslink.AgentUser

	if err := s.App.DB.Connection.Where("username = ?", req.Username).Limit(1).Find(&existingAgentUsers).Error; err != nil {
		return err
	}

	if len(existingAgentUsers) > 0 {
		return errors.New("username already exist")
	}

	now := time.Now().Unix()

	newAgentId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	newAgent := basslink.Agent{
		Id:         newAgentId.String(),
		Name:       req.AgentName,
		Country:    req.Country,
		Region:     req.Region,
		City:       req.City,
		Address:    req.Address,
		PhoneCode:  req.PhoneCode,
		PhoneNo:    req.PhoneNo,
		Email:      req.Email,
		Website:    req.Website,
		Timezone:   nil,
		IsVerified: false,
		IsEnable:   true,
		Created:    now,
	}

	newAgentUserId, err := uuid.NewV7()
	if err != nil {
		return err
	}

	newAgentUser := basslink.AgentUser{
		Id:        newAgentUserId.String(),
		AgentId:   newAgent.Id,
		Role:      basslink.AgentRoleOwner,
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
	newAgentUserCredential := basslink.AgentUserCredential{
		AgentUserId:    newAgentUser.Id,
		CredentialType: "password",
		CredentialData: fmt.Sprintf("%x", hashedPassword),
		Updated:        now,
	}

	if err = s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&newAgent).Error; err != nil {
			return err
		}

		if err = tx.Create(&newAgentUser).Error; err != nil {
			return err
		}

		if err = tx.Create(&newAgentUserCredential).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *Service) toggleAgentEnable(agentId string) error {
	var selectedAgent basslink.Agent

	if err := s.App.DB.Connection.First(&selectedAgent, agentId).Error; err != nil {
		return err
	}

	newAgentEnableValue := !selectedAgent.IsEnable

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		return tx.Model(basslink.Agent{}).Where("id = ?", selectedAgent.IsEnable).UpdateColumn("is_enable", newAgentEnableValue).Error
	}); err != nil {
		return err
	}

	return nil
}
