package agent

import (
	"CRM/src/lib/basslink"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *Service) SignIn(req *SignInRequest) (*SignInResponse, error) {
	var user basslink.AgentUser

	if err := s.App.DB.Connection.Preload("Agent").Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, err
	}

	var credential basslink.AgentUserCredential

	if err := s.App.DB.Connection.Where("agent_user_id = ? AND credential_type = ?", user.Id, "password").First(&credential).Error; err != nil {
		return nil, err
	}

	hasher := sha256.New()
	hasher.Write([]byte(req.Password))
	hashedPassword := hasher.Sum(nil)

	if fmt.Sprintf("%x", hashedPassword) != credential.CredentialData {
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
