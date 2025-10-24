package common

import (
	"CRM/src/lib/basslink"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

func (s *Service) handleGetRate(req *GetRateRequest) (*GetRequestResponse, error) {
	if req.FromAmount == nil && req.ToAmount == nil {
		return nil, fmt.Errorf("from amount and to amount cannot be empty")
	}

	var rate, feeRate basslink.Rate

	if err := s.App.DB.Connection.Where("from_currency = ? AND to_currency = ?", req.FromCurrency, req.ToCurrency).First(&rate).Error; err != nil {
		return nil, err
	}

	if err := s.App.DB.Connection.Where("from_currency = ? AND to_currency = ?", "idr", req.FromCurrency).First(&feeRate).Error; err != nil {
		return nil, err
	}

	var fromAmount, toAmount, totalFee float64 = 0, 0, 0
	var feePercentage = 2.5
	var feeFixed float64 = 6500

	if req.FromAmount != nil && len(*req.FromAmount) > 0 {
		if flAmount, err := strconv.ParseFloat(*req.FromAmount, 64); err == nil {
			fromAmount = flAmount
		}
	}

	if req.ToAmount != nil && len(*req.ToAmount) > 0 {
		if flAmount, err := strconv.ParseFloat(*req.ToAmount, 64); err == nil {
			toAmount = flAmount
		}
	}

	if fromAmount == 0 && toAmount == 0 {
		return nil, fmt.Errorf("from amount and to amount cannot be empty")
	}

	if fromAmount > 0 {
		totalFee = (feePercentage / 100 * fromAmount) + (feeFixed * feeRate.Rate)
		toAmount = (fromAmount - totalFee) * rate.Rate
		if toAmount < 0 {
			toAmount = 0
		} else {
			toAmount = math.Floor(toAmount)
		}
	} else if toAmount > 0 {
		fromAmount = ((toAmount / rate.Rate) + (feeFixed * feeRate.Rate)) / ((100 - feePercentage) / 100)
		totalFee = (feePercentage / 100 * fromAmount) + (feeFixed * feeRate.Rate)
		if fromAmount < 0 {
			fromAmount = 0
		} else {
			fromAmount = math.Ceil(fromAmount)
		}
	}

	resp := &GetRequestResponse{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		FromAmount:   json.Number(fmt.Sprintf("%f", fromAmount)),
		ToAmount:     json.Number(fmt.Sprintf("%f", toAmount)),
		Rate:         json.Number(fmt.Sprintf("%f", rate.Rate)),
		FeePercent:   json.Number(fmt.Sprintf("%f", feePercentage)),
		FeeFixed:     json.Number(fmt.Sprintf("%f", feeFixed)),
		TotalFee:     json.Number(fmt.Sprintf("%f", totalFee)),
	}

	return resp, nil
}
