package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "github.com/celerbuild/celerbuild-example-go/internel/handle"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAPIEndpoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := handler.SetupRouter()

	t.Run("Test Hello Endpoint", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, "Hello World from CelerBuild!", response["message"])
	})

	t.Run("Test Version Endpoint", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/version", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, handler.Version, response["version"])
	})
}
