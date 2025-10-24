package common

import (
	"CRM/src/lib/basslink"
)

func (s *Service) handleGetRate(req *GetRateRequest) (*basslink.RateInfo, error) {
	rateInfo, err := s.App.CalculateRate(req.FromCurrency, req.ToCurrency, req.FromAmount, req.ToAmount)
	if err != nil {
		return nil, err
	}

	return rateInfo, nil
}
