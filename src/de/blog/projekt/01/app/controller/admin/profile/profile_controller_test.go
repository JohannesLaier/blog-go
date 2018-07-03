package admin_profile

import (
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/app/models/comment"
	"de/blog/projekt/01/app/models/keyword"
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/core/db"
	"de/blog/projekt/01/core/http/session"
	"de/blog/projekt/01/core/util"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"strings"
	"testing"
	"time"
	"os"
)

func GetDB() (*db.DB, *author.Author, *post.Post) {
	// Init DB
	blog_db := db.Get("blog")

	blog_keyword := keyword.NewKeyword("Cloud")
	blog_keywords := db.NewDBCollection("keyword")
	blog_keywords.Add(blog_keyword)
	blog_keywords.Add(keyword.NewKeyword("IT"))
	blog_keywords.Add(keyword.NewKeyword("Security"))

	blog_authors := db.NewDBCollection("author")
	blog_author := author.NewAuthor("admin", "123456")
	blog_author2 := author.NewAuthor("peter", "secret")
	blog_authors.Add(blog_author)
	blog_authors.Add(blog_author2)

	blog_posts := db.NewDBCollection("post")
	blog_post := post.NewPost("Your first  blog title", "Our sub title", "<b>Blog Inhalt</b>", blog_author.GetID())
	blog_post.AddKeyword(blog_keyword.GetID())
	blog_posts.Add(blog_post)

	blog_comment := comment.NewComment("admin", "test", blog_post.GetID())
	blog_comment.SetShare(true)

	blog_comments := db.NewDBCollection("comment")
	blog_comments.Add(blog_comment)

	blog_db.AddCollection(blog_posts)
	blog_db.AddCollection(blog_authors)
	blog_db.AddCollection(blog_keywords)
	blog_db.AddCollection(blog_comments)

	return blog_db, blog_author, blog_post
}

func TestNewProfileController(t *testing.T) {
	SendRequest(t, "")
}

func SendRequest(t *testing.T, url string) string {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	_, author, _ := GetDB()

	// Create Controller
	ctrl := NewProfileController()
	ctrl_path := "/admin/profile"

	// Remove Session Protection
	ctrl.SetAuthWrapper(nil)

	// Get Handler
	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	ctrl_path = ctrl_path + url

	// Create request and response
	req := httptest.NewRequest("POST", ctrl_path, nil)
	resp := httptest.NewRecorder()

	// Create Session
	session_store := session.GetSessionStore()
	id, sess := session_store.New()
	sess.Put("CURRENT", author)

	expiration := time.Now().Add(time.Duration(15 * time.Minute))
	cookie := http.Cookie{Name: "SESSIONID", Value: id, Expires: expiration}

	req.AddCookie(&cookie)

	handler(resp, req)

	assert.Equal(t, 200, resp.Code)

	response := resp.Result()

	if response.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		return bodyString
	}

	return ""
}

func TestNewProfileController_Change_Password_Success(t *testing.T) {
	author_username := "admin"
	author_passowrd := "123456"
	author_passowrd_new := "geheim"

	ctrl_path := "?username=" + author_username
	ctrl_path = ctrl_path + "&password=" + author_passowrd
	ctrl_path = ctrl_path + "&newPassword=" + author_passowrd_new
	ctrl_path = ctrl_path + "&newPassword2=" + author_passowrd_new
	ctrl_path = ctrl_path + "&action=save"

	resp := SendRequest(t, ctrl_path)

	msg := "Successfully saved"

	assert.True(t, strings.Contains(resp, msg))
}

func TestNewProfileController_Change_Password_Not_Equal_PW(t *testing.T) {
	author_username := "admin"
	author_passowrd := "123456"
	author_passowrd_new := "geheim"

	ctrl_path :=  "?username=" + author_username
	ctrl_path = ctrl_path + "&password=" + author_passowrd
	ctrl_path = ctrl_path + "&newPassword=" + author_passowrd_new
	ctrl_path = ctrl_path + "&newPassword2=" + "other_pw"
	ctrl_path = ctrl_path + "&action=save"

	resp := SendRequest(t, ctrl_path)

	msg := "Passwords does not match"

	assert.True(t, strings.Contains(resp, msg))
}

func TestNewProfileController_Change_Password_Empty_PW(t *testing.T) {
	author_username := "admin"
	author_passowrd := "123456"

	ctrl_path := "?username=" + author_username
	ctrl_path = ctrl_path + "&password=" + author_passowrd
	ctrl_path = ctrl_path + "&newPassword="
	ctrl_path = ctrl_path + "&newPassword2="
	ctrl_path = ctrl_path + "&action=save"

	resp := SendRequest(t, ctrl_path)

	msg := "All fields must be specified"

	assert.True(t, strings.Contains(resp, msg))
}

func TestNewProfileController_Change_Password_Invalid_Password(t *testing.T) {
	author_username := "admin"
	author_passowrd_new := "geheim"

	ctrl_path := "?username=" + author_username
	ctrl_path = ctrl_path + "&password=" + "invalid_password"
	ctrl_path = ctrl_path + "&newPassword=" + author_passowrd_new
	ctrl_path = ctrl_path + "&newPassword2=" + author_passowrd_new
	ctrl_path = ctrl_path + "&action=save"

	resp := SendRequest(t, ctrl_path)

	msg := "Password is incorrect"

	assert.True(t, strings.Contains(resp, msg))
}