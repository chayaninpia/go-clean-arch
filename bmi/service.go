package bmi

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CalculateBmi(bmiRequest domain.BmiRequest) float32 {
	meterHeight := bmiRequest.Height
	if bmiRequest.Height > 3 {
		meterHeight = bmiRequest.Height / 100
	}

	return bmiRequest.Weight / (meterHeight * meterHeight)
}

func (s *Service) CheckBmi(bmiValue float32) domain.BmiResponse {

	result := ""
	if bmiValue <= 18.5 {
		result = "ผอมเกินไป"
	} else if bmiValue <= 25 {
		result = "น้ำหนักปกติ"
	} else if bmiValue <= 30 {
		result = "เริ่มอ้วน"
	} else if bmiValue <= 35 {
		result = "อ้วน"
	} else {
		result = "อ้วนมากผิดปกติ"
	}
	return domain.BmiResponse{
		Bmi:    bmiValue,
		Result: result,
	}
}
