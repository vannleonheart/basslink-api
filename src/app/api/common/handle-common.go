package common

import (
	"CRM/src/lib/basslink"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/vannleonheart/goutil"
)

func (s *Service) handleUploadFile(path string, form *multipart.Form) (*[]string, error) {
	result := make([]string, 0)

	for _, files := range form.File {
		if len(files) == 0 {
			return nil, errors.New("no file uploaded")
		}

		for _, file := range files {
			fileName := fmt.Sprintf("%s/%s", path, file.Filename)
			url, err := s.App.Storage.StorePublic(fileName, file)
			if err != nil {
				return nil, err
			}
			if url != nil {
				result = append(result, *url)
			}
		}
	}

	return &result, nil
}

func (s *Service) handleGetCurrencies() (*[]basslink.Currency, error) {
	var currencies []basslink.Currency

	if err := s.App.DB.Connection.Where("is_active = ?", true).Find(&currencies).Error; err != nil {
		return nil, err
	}

	return &currencies, nil
}

func (s *Service) handleCreateAppointment(req *CreateAppointmentRequest) (*basslink.Appointment, error) {
	var appointments []basslink.Appointment

	if err := s.App.DB.Connection.Where("(email = ? OR phone = ?) AND status = ?", req.Email, req.Phone, "new").Limit(1).Find(&appointments).Error; err != nil {
		return nil, err
	}

	if len(appointments) > 0 {
		return nil, errors.New("you have already booked an appointment with this contact")
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
