package admin_categorie

import (
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/app/models/keyword"
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/core/util"
	"de/blog/projekt/01/core/db"
	"strconv"
	"testing"
	"os"
)

func GetDB() (*db.DB, *post.Post, *keyword.Keyword) {
	// Init DB
	db_author_username := "admin"
	db_author_password := "s3cr3t"

	db_blog := db.Get("blog")
	db_authors := db.NewDBCollection("author")
	db_keywords := db.NewDBCollection("keyword")

	db_keyword := keyword.NewKeyword("Test Keyword")
	db_keywords.Add(db_keyword)

	db_author := author.NewAuthor(db_author_username, db_author_password)
	db_authors.Add(db_author)

	db_posts := db.NewDBCollection("post")
	db_post := post.NewPost("Your first  blog title", "Our sub title", "<b>Blog Inhalt</b>", db_author.GetID())
	db_post.AddKeyword(db_keyword.GetID())
	db_posts.Add(db_post)

	db_blog.AddCollection(db_authors)
	db_blog.AddCollection(db_keywords)
	db_blog.AddCollection(db_posts)

	return db_blog, db_post, db_keyword
}

func TestNewCategoryController_Index(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	_, _, blog_keyword := GetDB()

	// Create Controller
	ctrl := NewCategoryController()
	ctrl_path := "/admin/categories"

	// Remove Session Protection
	ctrl.SetAuthWrapper(nil)

	// Get Handler
	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	// Send Request
	assert.HTTPSuccess(t, handler, "GET", "https://localhost"+ctrl_path , nil)
	assert.HTTPBodyContains(t, handler, "GET", "https://localhost"+ctrl_path , nil, blog_keyword.Name)
}

func TestNewCategoryController_Add(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	GetDB()

	// Create Controller
	ctrl := NewCategoryController()
	ctrl_path := "/admin/categories"

	// Remove Session Protection
	ctrl.SetAuthWrapper(nil)

	// Get Handler
	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	keyword_text := "NewKeywordäöüß"

	// Request Value
	values := map[string][]string {
		"action" : []string { "add" },
		"category" : []string { keyword_text },
	}

	// Send Request
	assert.HTTPSuccess(t, handler, "GET", "https://localhost"+ctrl_path, values)
	assert.HTTPBodyContains(t, handler, "GET", "https://localhost"+ctrl_path , nil, keyword_text)
}

func TestNewCategoryController_Delete(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Init DB
	blog_db, blog_post, blog_keyword := GetDB()
	blog_keywords := blog_db.GetCollection("keyword")

	blog_keyword2 := keyword.NewKeyword("AnotherKeyWord")
	blog_keywords.Add(blog_keyword2)

	blog_post.AddKeyword(blog_keyword2.GetID())

	// Create Controller
	ctrl := NewCategoryController()
	ctrl_path := "/admin/categories"

	// Remove Session Protection
	ctrl.SetAuthWrapper(nil)

	// Get Handler
	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	// Request Value
	values := map[string][]string {
		"id" : []string { strconv.Itoa(blog_keyword2.GetID()) },
		"action" : []string { "delete" },
	}

	// Send Request
	assert.HTTPSuccess(t, handler, "GET", "https://localhost"+ctrl_path, values)
	assert.HTTPBodyContains(t, handler, "GET", "https://localhost"+ctrl_path , nil, blog_keyword.Name)
	assert.HTTPBodyNotContains(t, handler, "GET", "https://localhost"+ctrl_path , nil, blog_keyword2.Name)
}
