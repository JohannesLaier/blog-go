package admin_author

import (
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/app/models/keyword"
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/core/http/session"
	"de/blog/projekt/01/core/util"
	"de/blog/projekt/01/core/db"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
	"testing"
	"time"
	"os"
)

func GetDB() (*db.DB, *author.Author) {
	// Init DB
	db_author_username := "admin"
	db_author_password := "s3cr3t"

	db_blog := db.Get("blog")
	db_authors := db.NewDBCollection("author")

	blog_keyword := keyword.NewKeyword("Cloud")
	blog_keywords := db.NewDBCollection("keyword")
	blog_keywords.Add(blog_keyword)

	db_author := author.NewAuthor(db_author_username, db_author_password)
	db_authors.Add(db_author)

	blog_posts := db.NewDBCollection("post")
	blog_post := post.NewPost("Your first  blog title", "Our sub title", "<b>Blog Inhalt</b>", db_author.GetID())
	blog_post.AddKeyword(blog_keyword.GetID())

	blog_post2 := post.NewPost("title", "st", "content", db_author.GetID())
	blog_post2.AddKeyword(blog_keyword.GetID())

	blog_posts.Add(blog_post)
	blog_posts.Add(blog_post2)

	db_blog.AddCollection(db_authors)
	db_blog.AddCollection(blog_posts)
	db_blog.AddCollection(blog_keywords)

	return db_blog, db_author
}

func SendRequest(t *testing.T, author *author.Author, ctrlpath, url string, statusCode int) string {
	// Create Controller
	ctrl := NewAuthorController()
	ctrl_path := ctrlpath

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

	assert.Equal(t, statusCode, resp.Code)

	response := resp.Result()

	if response.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		return bodyString
	}

	return ""
}


func TestNewAuthorController_Authors(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	_, author := GetDB()

	ctrl_path := "/admin/authors"

	resp := SendRequest(t, author, ctrl_path, "", 200)

	assert.True(t, strings.Contains(resp, author.GetUsername()))
}

func TestNewAuthorController_Authors_Delete(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	blog_db, blog_author := GetDB()
	blog_authors := blog_db.GetCollection("author")
	blog_author2 := author.NewAuthor("jeff", "whatson")
	blog_authors.Add(blog_author2)

	// Send Request
	ctrl_path := "/admin/authors"
	url := "?id="+strconv.Itoa(blog_author2.GetID())
	url = url + "&action=delete"
	resp := SendRequest(t, blog_author, ctrl_path, url, 200)

	// Send request -> author 2 must be available
	resp = SendRequest(t, blog_author, ctrl_path, "", 200)
	assert.True(t, strings.Contains(resp, blog_author.GetUsername()))

	// Send request -> author 2 should be deleted
	resp = SendRequest(t, blog_author, ctrl_path, "", 200)
	assert.False(t, strings.Contains(resp, blog_author2.GetUsername()))
}

func TestNewAuthorController_Author_Detail(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	GetDB()

	// Create Controller
	ctrl := NewAuthorController()
	ctrl_path := "/admin/author-detail"

	// Remove Session Protection
	ctrl.SetAuthWrapper(nil)

	// Get Handler
	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	// Send Request
	assert.HTTPSuccess(t, handler, "GET", "https://localhost"+ctrl_path , nil)
}

func TestNewAuthorController_Author_Detail_Add(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	_, blog_author := GetDB()

	// Send Request
	ctrl_path := "/admin/author-detail"

	// Init Variables
	db_author_username := "peter"
	db_author_password := "geheim"

	url := "?username=" + db_author_username
	url = url + "&password=" + db_author_password
	url = url + "&action=add"

	// Send Request
	resp := SendRequest(t, blog_author, ctrl_path, url, 302)

	ctrl_path = "/admin/authors"

	// Send Request
	resp = SendRequest(t, blog_author, ctrl_path, "", 200)
	assert.True(t, strings.Contains(resp, db_author_username))
}