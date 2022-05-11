package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang-restAPI-FR/Core/Router/Public"
	"github.com/gin-gonic/gin"
)

func TestMain(t *testing.T) {
	router := gin.New()
	Public.APIRouter(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"welcome to restAPI\",\"readme\":\"https://github.com/F-8-Developer/golang-restAPI-FR/blob/main/README.md\",\"userInfo\":\"Hello World!!!\"}", w.Body.String())
}
