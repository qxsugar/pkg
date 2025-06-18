package kit

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var CustomError = errors.New("CustomError")

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	gin.SetMode(gin.DebugMode)
}

func setupApp() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", TranslateFunc(Pong))
	r.GET("/invalidArgument", TranslateFunc(func(ctx *gin.Context) (any, error) {
		return nil, NewInvalidArgumentError()
	}))
	r.GET("/customError", TranslateFunc(func(ctx *gin.Context) (any, error) {
		return nil, CustomError
	}))

	return r
}

func TestHandler(t *testing.T) {
	app := setupApp()

	t.Run("test ping", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
		app.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		body, err := io.ReadAll(w.Body)
		assert.Equal(t, err, nil)

		respBody := RespBody{}
		err = json.Unmarshal(body, &respBody)
		assert.Equal(t, err, nil)

		assert.Equal(t, respBody.Code, OK)
		assert.Equal(t, respBody.Succeeded, true)
		assert.Equal(t, respBody.RespData, "pong")
	})

	t.Run("test invalid argument error", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/invalidArgument", nil)
		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		body, err := io.ReadAll(w.Body)
		assert.Equal(t, err, nil)

		respBody := RespBody{}
		err = json.Unmarshal(body, &respBody)
		assert.Equal(t, err, nil)
		assert.Equal(t, respBody.Code, ErrInvalidArgument)
		assert.Equal(t, respBody.Info, messages[ErrInvalidArgument])
	})

	t.Run("test custom request error", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/customError", nil)
		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		body, err := io.ReadAll(w.Body)
		assert.Equal(t, err, nil)

		respBody := RespBody{}
		err = json.Unmarshal(body, &respBody)
		assert.Equal(t, err, nil)

		assert.Equal(t, respBody.Succeeded, false)
		assert.Equal(t, respBody.Code, InternalErrorCode)
		assert.Equal(t, respBody.Info, messages[ErrInternal])
		assert.Equal(t, respBody.Desc, CustomError.Error())
	})
}
