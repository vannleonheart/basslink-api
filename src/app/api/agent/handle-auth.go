package agent

import (
	"CRM/src/lib/basslink"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func (s *Service) signIn(req *SignInRequest) (*SignInResponse, error) {
	var user basslink.AgentUser

	if err := s.App.DB.Connection.Preload("Agent").Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, err
	}

	var credential basslink.AgentUserCredential

	if err := s.App.DB.Connection.Where("agent_user_id = ? AND credential_type = ?", user.Id, "password").First(&credential).Error; err != nil {
		return nil, err
	}

	if !s.App.VerifyPassword(req.Password, credential.CredentialData) {
		return nil, errors.New("invalid password")
	}

	if !user.IsEnable {
		return nil, errors.New("user account is disabled")
	}

	now := time.Now()
	jwtToken, err := s.App.CreateJwtToken(jwt.MapClaims{
		"iat":  now.Unix(),
		"exp":  now.Add(time.Minute * 120).Unix(),
		"user": user.Id,
		"as":   "agent",
	})
	if err != nil {
		return nil, err
	}

	return &SignInResponse{
		Token: jwtToken,
	}, nil
}

func (s *Service) updatePassword(user *basslink.AgentUser, req *UpdatePasswordRequest) error {
	var credential basslink.AgentUserCredential

	if err := s.App.DB.Connection.Where("agent_user_id = ? AND credential_type = ?", user.Id, "password").First(&credential).Error; err != nil {
		return err
	}

	if !user.IsEnable {
		return errors.New("user account is disabled")
	}

	if !s.App.VerifyPassword(req.Password, credential.CredentialData) {
		return errors.New("invalid password")
	}

	if err := s.App.DB.Connection.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(basslink.AgentUserCredential{}).Where("agent_user_id = ? AND credential_type = ?", user.Id, "password").Updates(map[string]interface{}{
			"credential_data": s.App.HashPassword(req.NewPassword),
			"updated":         time.Now().Unix(),
		}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
