package apix

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupApp() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Wrapper(Pong))
	r.GET("/invalidArgument", Wrapper(func(ctx *gin.Context) (interface{}, error) {
		return nil, NewInvalidArgumentError()
	}))

	return r
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	gin.SetMode(gin.TestMode)
}

func TestPing(t *testing.T) {
	app := setupApp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"succeeded\":true,\"resp_data\":\"pong\"}", w.Body.String())
}

func TestInvalidArgumentError(t *testing.T) {
	app := setupApp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/invalidArgument", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, InvalidArgument, w.Code)
}
