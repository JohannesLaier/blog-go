package shared

import (
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/core/http/session"
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/db"
	"net/http/httptest"
	"net/http"
	"testing"
	"time"
	"os"
	"io/ioutil"
)

func getDB(db_author_username, db_author_password string) (*db.DB, *author.Author) {
	// Init DB
	db_blog := db.Get("blog")
	db_authors := db.NewDBCollection("author")

	db_author := author.NewAuthor(db_author_username, db_author_password)
	db_authors.Add(db_author)

	db_blog.AddCollection(db_authors)

	return db_blog, db_author
}

func execHandler(ctrl *controller.Controller, path string, req *http.Request, resp http.ResponseWriter) {
	handler := ctrl.GetHandler()[path]
	handler(resp, req)
}

func TestSessionWrapper_Unauthorized(t *testing.T) {
	handler := func (w http.ResponseWriter, r *http.Request) {
		assert.True(t, true)
	}

	// Create request
	req := httptest.NewRequest("POST", "/", nil)
	resp := httptest.NewRecorder()

	session_wrapper := SessionWrapper(handler)
	session_wrapper(resp, req)

	assert.Equal(t, resp.Code, 302)
}

func TestSessionWrapper_Authorized(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../")
	defer os.Chdir(backup_cwd)

	handler := func (w http.ResponseWriter, r *http.Request) {
		assert.True(t, true)
	}

	// Create request
	req := httptest.NewRequest("POST", "/", nil)
	resp := httptest.NewRecorder()

	session_store := session.GetSessionStore()
	_, sess := session_store.New()

	expiration := time.Now().Add(time.Duration(15 * time.Minute))
	cookie := http.Cookie{Name: "SESSIONID", Value: sess.Id, Expires: expiration}
	req.AddCookie(&cookie)

	session_wrapper := SessionWrapper(handler)
	session_wrapper(resp, req)

	assert.Equal(t, resp.Code, 200)
}

func TestSessionWrapper_Authorized_Controller(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../")
	defer os.Chdir(backup_cwd)

	// Create Controller
	ctrl := controller.NewController("TestController")
	ctrl_path := "/admin/login"
	ctrl_response := "OK"

	ctrl.AddHandler(ctrl_path, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(ctrl_response))
	})

	// Init DB
	db_author_username := "admin"
	db_author_password := "s3cr3t"

	_, author := getDB(db_author_username, db_author_password)

	// Create request and response
	req := httptest.NewRequest("POST", "/test", nil)
	resp := httptest.NewRecorder()

	// Create Session
	session_store := session.GetSessionStore()
	id, sess := session_store.New()
	sess.Put("CURRENT", author)

	expiration := time.Now().Add(time.Duration(15 * time.Minute))
	cookie := http.Cookie{Name: "SESSIONID", Value: id, Expires: expiration}
	req.AddCookie(&cookie)

	execHandler(ctrl, ctrl_path, req, resp)

	resp.Flush()
	response := resp.Result()

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		bodyString := string(bodyBytes)

		assert.Equal(t, ctrl_response, bodyString)
	} else {
		assert.True(t, false)
	}
}
