package admin_logout

import (
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/core/http/session"
	"de/blog/projekt/01/core/util"
	"net/http/httptest"
	"net/http"
	"testing"
	"time"
	"os"
)

func TestNewLogoutController(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Create Controller
	ctrl := NewLogoutController()
	ctrl_path := "/admin/logout"

	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	// Create request
	req := httptest.NewRequest("POST", "/", nil)
	resp := httptest.NewRecorder()

	session_store := session.GetSessionStore()
	_, sess := session_store.New()

	session_id := "SESSIONID"

	expiration := time.Now().Add(time.Duration(15 * time.Minute))
	cookie := http.Cookie{Name: session_id, Value: sess.Id, Expires: expiration}
	req.AddCookie(&cookie)

	handler(resp, req)

	// Check Resp Cookie
	resp_cookie := resp.Result().Cookies()[0]
	assert.Equal(t, session_id, resp_cookie.Name)
	assert.Equal(t, time.Unix(0, 0).Unix(), resp_cookie.Expires.Unix())

	// Check redirect to /admin/login
	assert.Equal(t, resp.Code, 302)
}