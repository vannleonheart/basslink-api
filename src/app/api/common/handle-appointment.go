package common

import (
	"CRM/src/lib/basslink"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/vannleonheart/goutil"
)

func (s *Service) handleCreateAppointment(req *CreateAppointmentRequest) (*basslink.Appointment, error) {
	req.Email = strings.ToLower(req.Email)

	var appointments []basslink.Appointment

	if err := s.App.DB.Connection.Where("(email = ? OR phone = ?) AND status = ?", req.Email, req.Phone, "new").Limit(1).Find(&appointments).Error; err != nil {
		return nil, err
	}

	if len(appointments) > 0 {
		return nil, errors.New("you have already booked an appointment with this contact recently")
	}

	dt, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	if _, err = time.Parse("15:04", req.Time); err != nil {
		return nil, errors.New("invalid time format")
	}

	randomizer := goutil.NewRandomString(goutil.AlphaUNumCharset)
	id := fmt.Sprintf("%s-%s", dt.Format("20060102"), randomizer.GenerateRange(2, 4))

	newAppointment := basslink.Appointment{
		Id:          id,
		Name:        req.Name,
		Company:     req.Company,
		Email:       req.Email,
		Phone:       req.Phone,
		ServiceType: req.Service,
		Date:        req.Date,
		Time:        req.Time,
		Notes:       req.Notes,
		Status:      "new",
		Created:     time.Now().Unix(),
		Updated:     nil,
	}

	if err = s.App.DB.Connection.Create(&newAppointment).Error; err != nil {
		return nil, err
	}

	return &newAppointment, nil
}
