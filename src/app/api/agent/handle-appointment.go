package agent

import "CRM/src/lib/basslink"

func (s *Service) getAppointments() (*[]basslink.Appointment, error) {
	var appointments []basslink.Appointment

	if err := s.App.DB.Connection.Order("created DESC").Find(&appointments).Error; err != nil {
		return nil, err
	}

	return &appointments, nil
}
