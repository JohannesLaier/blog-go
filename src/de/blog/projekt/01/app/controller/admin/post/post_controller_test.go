package admin_post

import (
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/app/models/comment"
	"de/blog/projekt/01/app/models/keyword"
	"de/blog/projekt/01/core/db"
	"testing"
	"os"
	"de/blog/projekt/01/core/util"
	"github.com/stretchr/testify/assert"
	"time"
	"io/ioutil"
	"net/http/httptest"
	"de/blog/projekt/01/core/http/session"
	"net/http"
	"strings"
	"strconv"
)

func GetDB() (*db.DB, *author.Author, *post.Post, *post.Post) {
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

	blog_post2 := post.NewPost("title", "st", "content", blog_author2.GetID())
	blog_post2.AddKeyword(blog_keyword.GetID())

	blog_posts.Add(blog_post)
	blog_posts.Add(blog_post2)

	blog_comment := comment.NewComment("admin", "test", blog_post.GetID())
	blog_comment.SetShare(true)

	blog_comments := db.NewDBCollection("comment")
	blog_comments.Add(blog_comment)

	blog_db.AddCollection(blog_posts)
	blog_db.AddCollection(blog_authors)
	blog_db.AddCollection(blog_keywords)
	blog_db.AddCollection(blog_comments)

	return blog_db, blog_author, blog_post, blog_post2
}

func SendRequest(t *testing.T, author *author.Author, path string, url string) string {
	// Create Controller
	ctrl := NewPostController()
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


func TestNewPostController_Posts_List(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	_, blog_author, blog_post, blog_post2 := GetDB()

	ctrl_path := "/admin/posts"

	resp := SendRequest(t, blog_author, ctrl_path, "")

	assert.True(t, strings.Contains(resp, blog_post.Title))
	assert.True(t, strings.Contains(resp, blog_post.SubTitle))

	assert.True(t, strings.Contains(resp, blog_post2.Title))
	assert.True(t, strings.Contains(resp, blog_post2.SubTitle))
}

func TestNewPostController_Posts_Delete(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	_, blog_author, blog_post, blog_post2 := GetDB()

	ctrl_path := "/admin/posts"

	ctrl_url := "?id=" + strconv.Itoa(blog_post2.GetID())
	ctrl_url = ctrl_url + "&action=delete"

	resp := SendRequest(t, blog_author, ctrl_path, ctrl_url)

	assert.True(t, strings.Contains(resp, blog_post.Title))
}

func TestNewPostController_Posts_Detail(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	_, blog_author, blog_post, _ := GetDB()

	ctrl_path := "/admin/post-detail"
	ctrl_url := "?id=" + strconv.Itoa(blog_post.GetID())

	resp := SendRequest(t, blog_author, ctrl_path, ctrl_url)

	assert.True(t, strings.Contains(resp, blog_post.Title))
	assert.True(t, strings.Contains(resp, blog_post.SubTitle))
}