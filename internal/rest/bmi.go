package rest

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

//go:generate mockery --name BmiService
type BmiService interface {
	GetCalculateBmiResult(bmiRequest domain.Bmi) domain.Bmi
	CalculateBmi(bmiRequest domain.Bmi) float32
	CheckBmi(bmiValue float32) domain.Bmi
}

type BmiHandler struct {
	Service BmiService
}

func NewBmiHandler(e *echo.Echo, svc BmiService) {
	handler := &BmiHandler{
		Service: svc,
	}
	e.POST("/bmi", handler.CalculateBmi)
}

func isRequestBmiValid(m *domain.Bmi) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *BmiHandler) CalculateBmi(c echo.Context) error {
	req := domain.Bmi{}

	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestBmiValid(&req); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	bmiResult := a.Service.GetCalculateBmiResult(req)

	return c.JSON(http.StatusOK, bmiResult)
}
