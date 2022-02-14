package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-timestamp-ms/middlewares"
	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c, r := gin.CreateTestContext(rec)
	mid := middlewares.CORS()
	r.ServeHTTP(rec, req)
	mid(c)

	assert.Equal(t, "*", rec.Header().Get("Access-Control-Allow-Origin"))
}
