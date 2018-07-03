package blog_detail

import (
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/core/db"
	"de/blog/projekt/01/core/util"
	"de/blog/projekt/01/app/models/keyword"
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/app/models/comment"
	"de/blog/projekt/01/app/models/post"
	"net/http/httptest"
	"net/http"
	"strconv"
	"testing"
	"time"
	"os"
)

func GetDB() (*db.DB, *post.Post, *post.Post, *keyword.Keyword) {
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
	blog_post1 := post.NewPost("Your first  blog title", "Our sub title", "<b>Blog Inhalt</b>", blog_author.GetID())
	blog_post2 := post.NewPost("IT News", "Programming is awesome", "<b>Nothing is better than coding stuff</b>", blog_author.GetID())
	blog_post1.AddKeyword(blog_keyword.GetID())
	blog_post2.AddKeyword(blog_keyword.GetID())
	blog_posts.Add(blog_post1)
	blog_posts.Add(blog_post2)

	blog_comment := comment.NewComment("admin", "test", blog_post1.GetID())
	blog_comment.SetShare(true)

	blog_comments := db.NewDBCollection("comment")
	blog_comments.Add(blog_comment)

	blog_db.AddCollection(blog_posts)
	blog_db.AddCollection(blog_authors)
	blog_db.AddCollection(blog_keywords)
	blog_db.AddCollection(blog_comments)

	return blog_db, blog_post1, blog_post2, blog_keyword
}

func TestNewDetailController_Posts(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	_, blog_post1, blog_post2, blog_keyword := GetDB()

	// Create Controller
	ctrl := NewDetailController()
	ctrl_path := "/detail"

	// Get Handler
	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	// Test Multiple Blog Posts
	posts := []post.Post{ *blog_post1, *blog_post2 }

	for _, post := range posts {
		// Request Params
		values := map[string][]string {
			"id" : []string { strconv.Itoa(post.GetID()) },
		}

		// Send Request
		assert.HTTPSuccess(t, handler, "GET", "https://localhost"+ctrl_path , values)

		// Blog-Post Content
		assert.HTTPBodyContains(t, handler, "GET", "https://localhost"+ctrl_path , values, post.Title)
		assert.HTTPBodyContains(t, handler, "GET", "https://localhost"+ctrl_path , values, post.SubTitle)

		// Blog-Post Keywords
		assert.HTTPBodyContains(t, handler, "GET", "https://localhost"+ctrl_path , values, blog_keyword.Name)
	}
}

func TestNewDetailController_Comment(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	_, blog_post1, _, _ := GetDB()

	// Create Controller
	ctrl := NewDetailController()
	ctrl_path := "/detail"

	// Get Handler
	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	// Comment params
	comment_name := "MrInternet15"
	comment_text := "Thank you for sharing your greate blog article"

	// Request Params
	values_add_comment := map[string][]string {
		"id" : []string { strconv.Itoa(blog_post1.GetID()) },
		"action" : []string { "comment" },
		"username" : []string { comment_name },
		"text" : []string { comment_text },
	}

	// Request Params
	values_show_post := map[string][]string {
		"id" : []string { strconv.Itoa(blog_post1.GetID()) },
	}

	// Send Add Comment Request
	msg := "Successfully saved. It will be displayed after approval by"
	assert.HTTPBodyContains(t, handler, "POST", "https://localhost"+ctrl_path , values_add_comment, msg)

	// Blog post is looked until ist gets shared by an admin
	assert.HTTPBodyNotContains(t, handler, "GET", "https://localhost"+ctrl_path , values_show_post, comment_name)
	assert.HTTPBodyNotContains(t, handler, "GET", "https://localhost"+ctrl_path , values_show_post, comment_text)
}


func TestNewDetailController_Nickname(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	_, blog_post1, _, _ := GetDB()

	// Create Controller
	ctrl := NewDetailController()
	ctrl_path := "/detail"

	// Get Handler
	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	// Comment
	comment_nickname := "MrInternet110"

	// Request params to show a blog post
	// Peparing request post data
	ctrl_path = ctrl_path + "?id=" + strconv.Itoa(blog_post1.GetID())


	// Create request and response
	req := httptest.NewRequest("POST", ctrl_path, nil)
	resp := httptest.NewRecorder()

	expiration := time.Now().Add(time.Duration(100 * time.Minute))
	cookie := http.Cookie{Name: "NICKNAME", Value: comment_nickname, Expires: expiration}
	req.AddCookie(&cookie)

	handler(resp, req)

	assert.Equal(t, 200, resp.Code)
}
