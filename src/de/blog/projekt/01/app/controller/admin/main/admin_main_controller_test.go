package admin_main

import (
	"testing"
	"de/blog/projekt/01/core/util"
	"github.com/stretchr/testify/assert"
)

func TestNewAdminMainController_01(t *testing.T) {
	// Create Controller
	ctrl := NewAdminMainController()
	ctrl_path := "/admin/"

	// Get Handler
	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	// Send Request
	assert.HTTPRedirect(t, handler, "GET", "https://localhost"+ctrl_path , nil)
}

func TestNewAdminMainController_2(t *testing.T) {
	// Create Controller
	ctrl := NewAdminMainController()
	ctrl_path := "/admin"

	// Get Handler
	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	// Send Request
	assert.HTTPRedirect(t, handler, "GET", "https://localhost"+ctrl_path , nil)
}