package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-timestamp-ms/handlers"
	"github.com/madjiebimaa/fcc-timestamp-ms/models"
	"github.com/stretchr/testify/assert"
)

func TestUnixToUTC(t *testing.T) {
	gin.SetMode(gin.TestMode)

	unix := "1644278400"
	utc := "08 Feb 22 00:00 UTC"
	date := "2022-02-08"

	sucResBody, err := json.Marshal(models.TimeStamp{
		Unix: unix,
		UTC:  utc,
	})
	assert.NoError(t, err)

	failInvResBody, err := json.Marshal(gin.H{
		"message": "invalid url param value",
	})
	assert.NoError(t, err)

	testCases := []struct {
		name       string
		param      string
		statusCode int
		resBody    []byte
	}{
		{
			name:       "success with unix param",
			param:      unix,
			statusCode: http.StatusOK,
			resBody:    sucResBody,
		},
		{
			name:       "success with date param",
			param:      date,
			statusCode: http.StatusOK,
			resBody:    sucResBody,
		},
		{
			name:       "fail invalid param",
			param:      "fail",
			statusCode: http.StatusBadRequest,
			resBody:    failInvResBody,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/api/"+tt.param, nil)
			assert.NoError(t, err)

			rec := httptest.NewRecorder()
			_, r := gin.CreateTestContext(rec)
			r.GET("/api/:unix", handlers.UnixToUTCHandler)

			r.ServeHTTP(rec, req)

			assert.Equal(t, tt.statusCode, rec.Code)
			assert.Equal(t, tt.resBody, rec.Body.Bytes())
		})
	}

}
