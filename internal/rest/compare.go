package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CompareService interface {
	CompareBy2Word(word1, word2 string) (int, int)
	CompareByte(word string) int
}

type CompareHandler struct {
	Service CompareService
}

func NewCompareHandler(e *echo.Echo, svc CompareService) {
	handler := &CompareHandler{
		Service: svc,
	}
	e.GET("/compare", handler.CompareKorkaiandA)
}

func (a *CompareHandler) CompareKorkaiandA(c echo.Context) error {

	first, second := a.Service.CompareBy2Word("a", "ก")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"a": first,
		"ก": second,
	})
}
