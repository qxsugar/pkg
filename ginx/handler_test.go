package ginx

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testError = errors.New("test")

func setupApp() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Wrapper(Pong))
	r.GET("/invalidArgument", Wrapper(func(ctx *gin.Context) (interface{}, error) {
		return nil, NewInvalidArgumentError()
	}))
	r.GET("/test_error", Wrapper(func(ctx *gin.Context) (interface{}, error) {
		return nil, testError
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

func TestHandler(t *testing.T) {
	app := setupApp()

	t.Run("test ping", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
		app.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "{\"succeeded\":true,\"resp_data\":\"pong\"}", w.Body.String())
	})

	t.Run("test invalid argument error", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/invalidArgument", nil)
		app.ServeHTTP(w, req)

		assert.Equal(t, InvalidArgument, w.Code)
	})

	t.Run("test custom request error", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test_error", nil)
		app.ServeHTTP(w, req)

		assert.Equal(t, Internal, w.Code)
	})
}
