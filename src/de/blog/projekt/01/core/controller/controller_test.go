package controller

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"fmt"
)

func TestNewController(t *testing.T) {
	name := "MyController"
	ctrl := NewController(name)

	assert.NotNil(t, ctrl)
	assert.NotNil(t, ctrl.Handler)
	assert.Equal(t, name, ctrl.Name)
}

func TestController_AddHandler(t *testing.T) {
	path := "/test"
	content := "OK"

	ctrl := NewController("MyController")

	assert.Equal(t, 0, len(ctrl.GetHandler()), "Count of handlers should be zero")
	assert.NotZero(t, ctrl.GetHandler(), "handler should be an empty map")

	ctrl.AddHandler(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(w, content)
	})

	assert.Equal(t, 1, len(ctrl.GetHandler()), "Count of handlers should be zero")

	ctrl.AddHandler("/test/add", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(w, content)
	})

	assert.Equal(t, 2, len(ctrl.GetHandler()), "Count of handlers should be zero")
}

func TestController_GetHandler(t *testing.T) {
	path := "/test"
	content := "OK"

	ctrl := NewController("MyController")
	ctrl.AddHandler(path, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(content))
	})

	// Create request
	req := httptest.NewRequest("POST", "/test", nil)
	resp := httptest.NewRecorder()

	handler := ctrl.GetHandler()["/test"]
	handler(resp, req)

	resp.Flush()
	response := resp.Result();

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		resp_content := string(bodyBytes)

		assert.Equal(t, content, resp_content)
	} else {
		assert.True(t, false, "Invalid Status Code")
	}
}

func TestController_GetHandler_Auth_Handler(t *testing.T) {
	path := "/test"
	content := "OK"

	ctrl := NewController("MyController")
	ctrl.SetAuthWrapper(func (handler HandlerFunc) HandlerFunc {
		return func (w http.ResponseWriter, r *http.Request) {
			handler(w, r)
		}
	})
	ctrl.AddHandler(path, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(content))
	})

	// Create request
	req := httptest.NewRequest("POST", "/test", nil)
	resp := httptest.NewRecorder()

	handler := ctrl.GetHandler()["/test"]
	handler(resp, req)

	resp.Flush()
	response := resp.Result();

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		resp_content := string(bodyBytes)

		assert.Equal(t, content, resp_content)
	} else {
		assert.True(t, false, "Invalid Status Code")
	}
}

func TestController_SetAuthWrapper(t *testing.T) {
	path := "/test"

	ctrl := NewController("MyController")
	ctrl.SetAuthWrapper(func (handler HandlerFunc) HandlerFunc {
		return func (w http.ResponseWriter, r *http.Request) {
			handler(w, r)
		}
	})
	ctrl.AddHandler(path, func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, true)
	})

	handler := ctrl.GetHandler()["/test"]
	handler(nil, nil)
}