package admin

import (
	"CRM/src/lib/basslink"
	"strings"
	"time"
)

func (s *Service) getDisbursements(req *GetDisbursementFilter) (*[]basslink.Disbursement, error) {
	var disbursements []basslink.Disbursement

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

	if err := db.Find(&disbursements).Error; err != nil {
		return nil, err
	}

	return &disbursements, nil
}

func (s *Service) getDisbursement(disbursementId string) (*basslink.Disbursement, error) {
	var disbursement basslink.Disbursement

	if err := s.App.DB.Connection.First(&disbursement, disbursementId).Error; err != nil {
		return nil, err
	}

	return &disbursement, nil
}
