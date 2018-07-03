package comment

import (
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/core/db"
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/app/models/author"
)

func TestComment_Relation(t *testing.T) {
	blog_posts := db.NewDBCollection("post")
	blog_authors := db.NewDBCollection("author")
	blog_comments := db.NewDBCollection("comment")

	blog_author1 := author.NewAuthor("peter", "geheim")
	blog_author2 := author.NewAuthor("peter", "geheim")
	blog_authors.Add(blog_author1)
	blog_authors.Add(blog_author2)

	blog_post1 := post.NewPost("title", "subtitle", "content",  blog_author1.GetID())
	blog_post2 := post.NewPost("title", "subtitle", "content",  blog_author2.GetID())
	blog_posts.Add(blog_post1)
	blog_posts.Add(blog_post2)

	blog_comment1 := NewComment("MisterInternet", "text", blog_post1.GetID())
	blog_comment2 := NewComment("AnotherUsername", "example text", blog_post2.GetID())
	blog_comments.Add(blog_comment1)
	blog_comments.Add(blog_comment2)

	assert.Equal(t, blog_post1.GetID(), blog_comment1.GetPost(), "Incorrect relation to blog post")
	assert.Equal(t, blog_post2.GetID(), blog_comment2.GetPost(), "Incorrect relation to blog post")
}

func TestNewComment(t *testing.T) {
	username := "peterßßßäääöö"
	text := "this is the text of the comment ääüüü"
	post_id := 5

	c := NewComment(username, text, post_id)

	assert.Equal(t, c.GetUsername(), username, "Username doesnt match")
	assert.Equal(t, c.GetText(), text, "Text doesnt match")
	assert.Equal(t, c.GetPost(), post_id, "Post-Relataion doesnt match")
}


func TestComment_GetID(t *testing.T) {
	// Test via setter

	id := 1
	c := NewComment("admin", "test", 1)
	c.SetID(id)

	assert.Equal(t, id, c.GetID(), "could not get id")

	// Test via collection
	collection := db.NewDBCollection("testComment")

	username := "peterßßßäääöö"
	text := "this is the text of the comment ääüüü"
	post_id := 5

	c = NewComment(username, text, post_id)

	collection.Add(c)

	assert.Equal(t, 1, collection.GetAll()[0].GetID(), "Get wrong keyword id")
}

func TestComment_GetUsername(t *testing.T) {
	username := "admin"
	c := NewComment(username, "test", 1)

	assert.Equal(t, username, c.GetUsername(), "could not get username")
}

func TestComment_GetText(t *testing.T) {
	text := "test"
	c := NewComment("admin", text, 1)

	assert.Equal(t,  text, c.GetText(),"could not get text")
}

func TestComment_GetPost(t *testing.T) {
	post := 1
	c := NewComment("admin", "text", post)

	assert.Equal(t, post, c.GetPost(), "could not get post")
}

func TestComment_GetShare(t *testing.T) {
	c := NewComment("admin", "test", 1)

	assert.Equal(t, false, c.GetShare(),"wrong shareStatus")
}

func TestComment_SetShare(t *testing.T) {
	c := NewComment("admin", "test", 1)

	c.SetShare(true)

	assert.True(t, c.GetShare(), "share is still in status false")
}

func TestComment_SetID(t *testing.T) {
	c := NewComment("admin", "text", 1)

	newId := 2
	c.SetID(newId)

	assert.Equal(t, newId, c.GetID(), "id could no be set")
}

func TestComment_GetDate(t *testing.T) {
	start := time.Now()

	c := NewComment("admin", "text", 1)

	end := time.Now()

	assert.True(t, start.Unix() <= c.GetDate().Unix(), "Date must be between start and end")
	assert.True(t, end.Unix() >= c.GetDate().Unix(), "Date must be between start and end")
}