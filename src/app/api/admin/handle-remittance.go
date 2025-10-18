package admin

import (
	"CRM/src/lib/basslink"
	"strings"
	"time"
)

func (s *Service) getRemittances(req *GetRemittanceFilter) (*[]basslink.Remittance, error) {
	var remittances []basslink.Remittance

	db := s.App.DB.Connection.Preload("SourceCurrency").Preload("TargetCurrency").Preload("Attachments")

	if req != nil {
		if req.Status != nil && *req.Status != "" && strings.ToLower(*req.Status) != "all" {
			db = db.Where("status = ?", *req.Status)
		}

		if req.Search != nil && *req.Search != "" {
			db = db.Where("id LIKE ?", "%"+*req.Search+"%")
		}

		if req.Start != nil {
			if startTimestamp, err := time.Parse("2006-01-02", *req.Start); err == nil {
				db = db.Where("created >= ?", startTimestamp.Unix())
			}
		}

		if req.End != nil {
			if endTimestamp, err := time.Parse("2006-01-02", *req.End); err == nil {
				db = db.Where("created <= ?", endTimestamp.Unix())
			}
		}
	}

	if err := db.Find(&remittances).Error; err != nil {
		return nil, err
	}

	return &remittances, nil
}

func (s *Service) getRemittance(remittanceId string) (*basslink.Remittance, error) {
	var remittance basslink.Remittance

	if err := s.App.DB.Connection.Where("id = ?", remittanceId).First(&remittance).Error; err != nil {
		return nil, err
	}

	return &remittance, nil
}
