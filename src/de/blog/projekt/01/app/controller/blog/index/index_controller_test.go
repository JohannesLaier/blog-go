package blog_index

import (
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/core/db"
	"de/blog/projekt/01/core/util"
	"de/blog/projekt/01/app/models/keyword"
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/app/models/comment"
	"de/blog/projekt/01/app/models/post"
	"testing"
	"os"
)

func GetDB() (*db.DB, *post.Post, *post.Post) {
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

	return blog_db, blog_post1, blog_post2
}

func TestNewIndexController(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	_, blog_post1, blog_post2 := GetDB()

	// Create Controller
	ctrl := NewIndexController()
	ctrl_path := "/"

	// Get Handler
	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	// Send Request
	assert.HTTPSuccess(t, handler, "GET", "https://localhost"+ctrl_path , nil)

	// First Blog Post
	assert.HTTPBodyContains(t, handler, "GET", "https://localhost"+ctrl_path , nil, blog_post1.Title)
	assert.HTTPBodyContains(t, handler, "GET", "https://localhost"+ctrl_path , nil, blog_post1.SubTitle)

	// Second Blog Post
	assert.HTTPBodyContains(t, handler, "GET", "https://localhost"+ctrl_path , nil, blog_post2.Title)
	assert.HTTPBodyContains(t, handler, "GET", "https://localhost"+ctrl_path , nil, blog_post2.SubTitle)
}