package view

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)

func TestNewView(t *testing.T) {
	pathView := "admin/views/login"
	pathLayout := "admin/layout/layout_login"

	view := NewView(pathView, pathLayout)
	assert.NotNil(t, view.FuncMap)
	assert.NotNil(t, view.FuncMap["date"])
	assert.Equal(t, view.PathView, pathView)
	assert.Equal(t, view.PathLayout, pathLayout)
}

func TestView_AddFunc(t *testing.T) {
	v := NewView("admin/views/login", "admin/layout/layout_login")

	v.AddFunc("test", map[string]string{})

	assert.NotEmpty(t, v.FuncMap, "Function Map is empty")
}

func TestView_AddFunc2(t *testing.T) {
	v := NewView("admin/views/login", "admin/layout/layout_login")

	assert.Equal(t, len(v.FuncMap), 1, "Length incorrect")

	v.AddFunc("test", map[string]string{})

	assert.Equal(t, len(v.FuncMap), 2, "No Function was added")
}

func TestView_SetModel(t *testing.T) {
	v := NewView("admin/views/login", "admin/layout/layout_login")
	v.SetModel(map[string]string{})

	assert.Equal(t, v.Model, map[string]string{}, "Model wasn't set")

	model := []int {1, 2, 3}
	v.SetModel(model)

	assert.Equal(t, v.Model, model, "Model couldnt be set")
}

func TestView_Write(t *testing.T) {

	backup_cwd, _ := os.Getwd()
	os.Chdir("../../")
	defer os.Chdir(backup_cwd)

	w := httptest.NewRecorder()
	v := NewView("blog/views/blog_about", "blog/layout/layout")
	v.Write(w)

	response := w.Result()

	if response.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		bodyString := string(bodyBytes)

		assert.True(t, strings.Contains(bodyString, "Blog"))
	}
}