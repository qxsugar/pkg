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

	// Success cases
	r.GET("/ping", TranslateFunc(Pong))
	r.GET("/success-with-data", TranslateFunc(func(ctx *gin.Context) (any, error) {
		return map[string]interface{}{
			"message": "success",
			"data":    []int{1, 2, 3},
		}, nil
	}))
	r.GET("/success-nil-data", TranslateFunc(func(ctx *gin.Context) (any, error) {
		return nil, nil
	}))

	// Business error cases
	r.GET("/invalidArgument", TranslateFunc(func(ctx *gin.Context) (any, error) {
		return nil, NewInvalidArgumentError()
	}))
	r.GET("/notFound", TranslateFunc(func(ctx *gin.Context) (any, error) {
		return nil, NewNotFoundError().WithErr(errors.New("resource not found in database"))
	}))
	r.GET("/permissionDenied", TranslateFunc(func(ctx *gin.Context) (any, error) {
		return nil, NewPermissionDeniedError()
	}))

	// Custom error cases
	r.GET("/customError", TranslateFunc(func(ctx *gin.Context) (any, error) {
		return nil, CustomError
	}))
	r.GET("/nilError", TranslateFunc(func(ctx *gin.Context) (any, error) {
		return nil, nil // Return nil instead of (*Exception)(nil)
	}))

	// RouterGroup wrapper tests
	group := NewRouterGroup(r.Group("/api"))
	group.GET("/test", func(ctx *gin.Context) (any, error) {
		return "group test", nil
	})
	group.POST("/test", func(ctx *gin.Context) (any, error) {
		return "post test", nil
	})
	group.PUT("/test", func(ctx *gin.Context) (any, error) {
		return "put test", nil
	})
	group.DELETE("/test", func(ctx *gin.Context) (any, error) {
		return "delete test", nil
	})
	group.PATCH("/test", func(ctx *gin.Context) (any, error) {
		return "patch test", nil
	})

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

	t.Run("test success with complex data", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/success-with-data", nil)
		app.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		body, err := io.ReadAll(w.Body)
		assert.Equal(t, err, nil)

		respBody := RespBody{}
		err = json.Unmarshal(body, &respBody)
		assert.Equal(t, err, nil)

		assert.Equal(t, respBody.Code, OK)
		assert.Equal(t, respBody.Succeeded, true)
		assert.NotEqual(t, nil, respBody.RespData)
	})

	t.Run("test success with nil data", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/success-nil-data", nil)
		app.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		body, err := io.ReadAll(w.Body)
		assert.Equal(t, err, nil)

		respBody := RespBody{}
		err = json.Unmarshal(body, &respBody)
		assert.Equal(t, err, nil)

		assert.Equal(t, respBody.Code, OK)
		assert.Equal(t, respBody.Succeeded, true)
		assert.Equal(t, respBody.RespData, nil)
	})

	t.Run("test invalid argument error", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/invalidArgument", nil)
		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)

		body, err := io.ReadAll(w.Body)
		assert.Equal(t, err, nil)

		respBody := RespBody{}
		err = json.Unmarshal(body, &respBody)
		assert.Equal(t, err, nil)
		assert.Equal(t, respBody.Code, ErrInvalidArgument)
		assert.Equal(t, respBody.Info, Messages[ErrInvalidArgument])
		assert.Equal(t, respBody.Succeeded, false)
	})

	t.Run("test not found error with description", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/notFound", nil)
		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)

		body, err := io.ReadAll(w.Body)
		assert.Equal(t, err, nil)

		respBody := RespBody{}
		err = json.Unmarshal(body, &respBody)
		assert.Equal(t, err, nil)
		assert.Equal(t, respBody.Code, ErrNotFound)
		assert.Equal(t, respBody.Info, Messages[ErrNotFound])
		assert.Equal(t, respBody.Succeeded, false)
		assert.Equal(t, respBody.Desc, "resource not found in database")
	})

	t.Run("test permission denied error", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/permissionDenied", nil)
		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusForbidden, w.Code)

		body, err := io.ReadAll(w.Body)
		assert.Equal(t, err, nil)

		respBody := RespBody{}
		err = json.Unmarshal(body, &respBody)
		assert.Equal(t, err, nil)
		assert.Equal(t, respBody.Code, ErrPermissionDenied)
		assert.Equal(t, respBody.Info, Messages[ErrPermissionDenied])
		assert.Equal(t, respBody.Succeeded, false)
	})

	t.Run("test custom request error", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/customError", nil)
		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		body, err := io.ReadAll(w.Body)
		assert.Equal(t, err, nil)

		respBody := RespBody{}
		err = json.Unmarshal(body, &respBody)
		assert.Equal(t, err, nil)

		assert.Equal(t, respBody.Succeeded, false)
		assert.Equal(t, respBody.Code, InternalErrorCode)
		assert.Equal(t, respBody.Info, Messages[ErrInternal])
		assert.Equal(t, respBody.Desc, CustomError.Error())
	})

	t.Run("test nil error", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/nilError", nil)
		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		body, err := io.ReadAll(w.Body)
		assert.Equal(t, err, nil)

		respBody := RespBody{}
		err = json.Unmarshal(body, &respBody)
		assert.Equal(t, err, nil)

		assert.Equal(t, respBody.Succeeded, true)
		assert.Equal(t, respBody.Code, OK)
	})
}

func TestRouterGroup(t *testing.T) {
	app := setupApp()

	testCases := []struct {
		method       string
		path         string
		expectedData string
	}{
		{http.MethodGet, "/api/test", "group test"},
		{http.MethodPost, "/api/test", "post test"},
		{http.MethodPut, "/api/test", "put test"},
		{http.MethodDelete, "/api/test", "delete test"},
		{http.MethodPatch, "/api/test", "patch test"},
	}

	for _, tc := range testCases {
		t.Run("test "+tc.method+" method", func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tc.method, tc.path, nil)
			app.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)

			body, err := io.ReadAll(w.Body)
			assert.Equal(t, err, nil)

			respBody := RespBody{}
			err = json.Unmarshal(body, &respBody)
			assert.Equal(t, err, nil)

			assert.Equal(t, respBody.Succeeded, true)
			assert.Equal(t, respBody.Code, OK)
			assert.Equal(t, respBody.RespData, tc.expectedData)
		})
	}
}

func TestHTTPStatusCodes(t *testing.T) {
	r := gin.New()
	
	// Test various error codes and their HTTP status mappings
	testCases := []struct {
		name           string
		errorFunc      func() *Exception
		expectedHTTP   int
		expectedCode   int
	}{
		{"InvalidArgument", NewInvalidArgumentError, http.StatusBadRequest, ErrInvalidArgument},
		{"FailedPrecondition", NewFailedPreconditionError, http.StatusBadRequest, ErrFailedPrecondition},
		{"OutOfRange", NewOutOfRangeError, http.StatusBadRequest, ErrOutOfRange},
		{"Unauthenticated", NewUnauthenticatedError, http.StatusUnauthorized, ErrUnauthenticated},
		{"PermissionDenied", NewPermissionDeniedError, http.StatusForbidden, ErrPermissionDenied},
		{"NotFound", NewNotFoundError, http.StatusNotFound, ErrNotFound},
		{"Aborted", NewAbortedError, http.StatusConflict, ErrAborted},
		{"AlreadyExists", NewAlreadyExistsError, http.StatusConflict, ErrAlreadyExists},
		{"ResourceExhausted", NewResourceExhaustedError, http.StatusTooManyRequests, ErrResourceExhausted},
		{"Cancelled", NewCancelledError, 499, ErrCancelled},
		{"DataLoss", NewDataLossError, http.StatusInternalServerError, ErrDataLoss},
		{"Unknown", NewUnknownError, http.StatusInternalServerError, ErrUnknown},
		{"Internal", NewInternalError, http.StatusInternalServerError, ErrInternal},
		{"NotImplemented", NewNotImplementedError, http.StatusNotImplemented, ErrNotImplemented},
		{"Unavailable", NewUnavailableError, http.StatusServiceUnavailable, ErrUnavailable},
		{"DeadlineExceeded", NewDeadlineExceededError, http.StatusGatewayTimeout, ErrDeadlineExceeded},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			path := "/" + tc.name
			r.GET(path, TranslateFunc(func(ctx *gin.Context) (any, error) {
				return nil, tc.errorFunc()
			}))

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, path, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHTTP, w.Code)

			body, err := io.ReadAll(w.Body)
			assert.Equal(t, err, nil)

			respBody := RespBody{}
			err = json.Unmarshal(body, &respBody)
			assert.Equal(t, err, nil)

			assert.Equal(t, tc.expectedCode, respBody.Code)
			assert.Equal(t, false, respBody.Succeeded)
		})
	}
}

func TestTranslateFunc_ProductionMode(t *testing.T) {
	// Test in production mode where debug info is not shown
	originalMode := gin.Mode()
	gin.SetMode(gin.ReleaseMode)
	defer gin.SetMode(originalMode)

	r := gin.New()
	r.GET("/business-error", TranslateFunc(func(ctx *gin.Context) (any, error) {
		return nil, NewNotFoundError().WithErr(errors.New("sensitive error details"))
	}))
	r.GET("/custom-error", TranslateFunc(func(ctx *gin.Context) (any, error) {
		return nil, errors.New("sensitive custom error")
	}))

	t.Run("business error in production", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/business-error", nil)
		r.ServeHTTP(w, req)

		body, err := io.ReadAll(w.Body)
		assert.Equal(t, err, nil)

		respBody := RespBody{}
		err = json.Unmarshal(body, &respBody)
		assert.Equal(t, err, nil)

		assert.Equal(t, respBody.Succeeded, false)
		assert.Equal(t, respBody.Code, ErrNotFound)
		assert.Equal(t, respBody.Info, Messages[ErrNotFound])
		assert.Equal(t, respBody.Desc, "") // Should be empty in production
	})

	t.Run("custom error in production", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/custom-error", nil)
		r.ServeHTTP(w, req)

		body, err := io.ReadAll(w.Body)
		assert.Equal(t, err, nil)

		respBody := RespBody{}
		err = json.Unmarshal(body, &respBody)
		assert.Equal(t, err, nil)

		assert.Equal(t, respBody.Succeeded, false)
		assert.Equal(t, respBody.Code, InternalErrorCode)
		assert.Equal(t, respBody.Info, Messages[ErrInternal])
		assert.Equal(t, respBody.Desc, "") // Should be empty in production
	})
}
