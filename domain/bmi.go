package domain

type Bmi struct {
	Height float32 `json:"height" validate:"required"`
	Weight float32 `json:"weight" validate:"required"`
	Bmi    float32 `json:"bmi"`
	Result string  `json:"result"`
}
