package admin_login

import (
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/core/util"
	"de/blog/projekt/01/core/db"
	"testing"
	"os"
)

func GetDB(db_author_username, db_author_password string) (*db.DB, *author.Author) {
	// Init DB
	db_blog := db.Get("blog")
	db_authors := db.NewDBCollection("author")

	db_author := author.NewAuthor(db_author_username, db_author_password)
	db_authors.Add(db_author)

	db_blog.AddCollection(db_authors)

	return db_blog, db_author
}

func TestNewLoginController_LoginNomal(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Create Controller
	ctrl := NewLoginController()
	ctrl_path := "/admin/login"

	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	assert.HTTPSuccess(t, handler, "GET", "https://localhost/admin/login" , nil)
	assert.HTTPBodyContains(t, handler, "GET", "https://localhost/admin/login" , nil, "Username")
}

func TestNewLoginController_Login_Action(t *testing.T) {
	backup_cwd, _ := os.Getwd()
	os.Chdir("../../../../")
	defer os.Chdir(backup_cwd)

	// Create Controller
	ctrl := NewLoginController()
	ctrl_path := "/admin/login"

	handler := util.UnitTest_GetHandlerWrapped(ctrl, ctrl_path)

	db_author_username := "admin"
	db_author_password := "s3cr3t"

	GetDB(db_author_username, db_author_password)

	value := map[string][]string {
		"username" : []string { db_author_username },
		"password" : []string { db_author_password },
	}

	// Corrrect Password
	assert.HTTPRedirect(t, handler, "POST", "https://localhost/admin/login" , value)

	// Change password value
	value["password"] = []string { "invalid_pw" }

	// Invalid Password
	assert.HTTPBodyContains(t, handler, "POST", "https://localhost/admin/login" , value, "Username or password is incorrect")
}