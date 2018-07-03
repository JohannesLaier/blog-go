package admin_comment

import (
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/app/models/keyword"
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/app/models/comment"
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/core/http/session"
	"de/blog/projekt/01/core/util"
	"de/blog/projekt/01/core/db"
	"net/http/httptest"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"strconv"
	"time"
	"os"
)

func GetDB() (*db.DB, *comment.Comment, *post.Post, *author.Author) {
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

	blog_comment1 := comment.NewComment("admin", "test", blog_post.GetID())
	blog_comment1.SetShare(true)

	blog_comment2 := comment.NewComment("admin", "test", blog_post.GetID())
	blog_comment2.SetShare(false)

	blog_comments := db.NewDBCollection("comment")
	blog_comments.Add(blog_comment1)
	blog_comments.Add(blog_comment2)

	blog_db.AddCollection(blog_posts)
	blog_db.AddCollection(blog_authors)
	blog_db.AddCollection(blog_keywords)
	blog_db.AddCollection(blog_comments)

	return blog_db, blog_comment1, blog_post, blog_author
}

func SendRequest(t *testing.T, author *author.Author, path string, url string) string {
	// Create Controller
	ctrl := NewCommentController()
	ctrl_path := path

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


func TestNewCommentController(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	_, blog_comment, _, blog_author := GetDB()

	ctrl_path := "/admin/comments"

	// Send Request
	resp := SendRequest(t, blog_author, ctrl_path, "")

	assert.True(t, strings.Contains(resp, blog_comment.GetText()))
	assert.True(t, strings.Contains(resp, blog_comment.GetUsername()))
}

func TestNewCommentController_Share(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	blog_db, _, blog_post, blog_author := GetDB()

	blog_comments := blog_db.GetCollection("comment")

	blog_comment2 := comment.NewComment("NickName", "This is a comment value", blog_post.GetID())
	blog_comment2.SetShare(false)
	blog_comments.Add(blog_comment2)

	ctrl_path := "/admin/comments"

	// Request Value
	ctrl_url := "?id=" + strconv.Itoa(blog_comment2.GetID())
	ctrl_url = ctrl_url + "&action=share"

	// Send Request
	SendRequest(t, blog_author, ctrl_path, ctrl_url)

	// Check Value in Database
	blog_comment2 = blog_comments.GetByID(blog_comment2.GetID()).(*comment.Comment)

	assert.True(t, blog_comment2.GetShare())
}

func TestNewCommentController_Delete(t *testing.T) {

	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	blog_db, blog_comment, blog_post, blog_author := GetDB()

	blog_comments := blog_db.GetCollection("comment")

	blog_comment2 := comment.NewComment("NickName", "This is a comment value", blog_post.GetID())
	blog_comments.Add(blog_comment2)

	ctrl_path := "/admin/comments"

	// Request Value
	ctrl_url := "?id=" + strconv.Itoa(blog_comment2.GetID())
	ctrl_url = ctrl_url + "&action=delete"

	// Send Request
	SendRequest(t, blog_author, ctrl_path, ctrl_url)

	// Send Request
	resp := SendRequest(t, blog_author, ctrl_path, "")

	assert.True(t, strings.Contains(resp, blog_comment.GetText()))
	assert.False(t, strings.Contains(resp, blog_comment2.GetText()))
}