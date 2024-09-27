package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/go-clean-arch/bmi"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBmiHandler_CalculateBmi(t *testing.T) {
	mockBmi := domain.Bmi{
		Height: 100,
		Weight: 50,
	}

	tempMockBmi := mockBmi

	j, err := json.Marshal(tempMockBmi)
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequestWithContext(context.TODO(), echo.POST, "/bmi", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/bmi")

	handler := BmiHandler{
		Service: bmi.NewService(),
	}
	err = handler.CalculateBmi(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}
