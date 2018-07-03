package blog_contact


import (
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/core/util"
	"testing"
	"os"
)

func TestNewContactController(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Create Controller
	ctrl := NewContactController()
	ctrl_path := "/contact"

	// Get Handler
	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	// Send Request
	assert.HTTPSuccess(t, handler, "GET", "https://localhost"+ctrl_path , nil)

	names := []string {"Jana Hockenberger", "Annika Keil", "Johannes Laier", "@lehre.mosbach.dhbw.de"}
	for _, name := range names {
		assert.HTTPBodyContains(t, handler, "GET", "https://localhost"+ctrl_path , nil, name)
	}
}
