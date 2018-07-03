package util

import (
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/core/controller"
	"net/http/httptest"
	"net/http"
	"testing"
	"io/ioutil"
)

func Test_RandomString(t *testing.T) {
	for i := 0; i < 100; i++ {
		assert.Equal(t, i, len(RandomString(i)), "String length does not match")
	}
}

func HelperUnitTest_GetHandler_Handler(t *testing.T, ctrl_path string, ctrl_content string, handler func(w http.ResponseWriter, r *http.Request)) {
	// Create request and response
	req := httptest.NewRequest("POST", ctrl_path, nil)
	resp := httptest.NewRecorder()

	handler(resp, req)

	response := resp.Result()

	if response.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		assert.Equal(t, ctrl_content, bodyString)
	}
}

func TestUnitTest_GetHandler(t *testing.T) {
	ctrl_path := "/path"
	ctrl_content := "OK-VALUE"
	ctrl_handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(ctrl_content))
	}
	ctrl := controller.NewController("Test")
	ctrl.AddHandler(ctrl_path, ctrl_handler)

	handler := UnitTest_GetHandler(ctrl, ctrl_path)

	HelperUnitTest_GetHandler_Handler(t, ctrl_path, ctrl_content, handler)
}

func TestUnitTest_GetHandlerWrapped(t *testing.T) {
	ctrl_path := "/path"
	ctrl_content := "OK-VALUE"
	ctrl_handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(ctrl_content))
	}
	ctrl := controller.NewController("Test")
	ctrl.AddHandler(ctrl_path, ctrl_handler)

	handler := UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	HelperUnitTest_GetHandler_Handler(t, ctrl_path, ctrl_content, handler)
}