package domain

type BmiRequest struct {
	Height float32 `json:"height" validate:"required"`
	Weight float32 `json:"weight" validate:"required"`
}

type BmiResponse struct {
	Bmi    float32 `json:"bmi"`
	Result string  `json:"result"`
}
