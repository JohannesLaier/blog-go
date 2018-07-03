package blog_about

import (
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/core/util"
	"testing"
	"os"
)

func TestNewAboutController(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Create Controller
	ctrl := NewAboutController()
	ctrl_path := "/about"

	// Get Handler
	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	// Send Request
	assert.HTTPSuccess(t, handler, "GET", "https://localhost"+ctrl_path , nil)

	// Check page contaning matrikel no
	matrikel_nr := []string {"8892993", "1734394", "1777093"}
	for _, mat_nr := range matrikel_nr {
		assert.HTTPBodyContains(t, handler, "GET", "https://localhost"+ctrl_path , nil, mat_nr)
	}
}