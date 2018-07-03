package post

import (
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/core/db"
)

func TestPost_GetContent(t *testing.T) {
	content := "Test"
	p:= NewPost("Title", "SubTitle", content, 5)

	assert.Equal(t, content, p.GetContent(), "Get wrong content")
}

func TestPost_GetDate(t *testing.T) {
	start := time.Now()

	p:= NewPost("Title", "SubTitle", "Test", 5)

	end := time.Now()

	assert.True(t, start.Unix() <= p.GetDate().Unix(), "Get wrong date")
	assert.True(t, end.Unix() >= p.GetDate().Unix(), "Get wrong date")
}

func TestPost_GetAuthorID(t *testing.T) {
	author := 1
	p:= NewPost("Title", "SubTitle", "Test", author)

	assert.Equal(t, author, p.GetAuthorID(), "Get wrong author (id)")
}

func TestPost_GetID(t *testing.T) {
	collection := db.NewDBCollection("testPost")

	p:= NewPost("Title", "SubTitle", "Test", 1)

	collection.Add(p)

	assert.Equal(t, 1, collection.GetAll()[0].GetID(), "Get wrong post id")
}


func TestPost_GetKeywords(t *testing.T) {
	title := "Post Title"
	subtitle := "This is a awesome subtitle"
	content := "<h1>Post</h1>\n<p>Test Content</p>"
	author := 1

	p := NewPost(title, subtitle, content, author)
	p.AddKeyword(1)
	p.AddKeyword(5)
	p.AddKeyword(13)
	p.AddKeyword(15)

	assert.Equal(t, []int{1, 5, 13, 15}, p.GetKeywords())

	p.AddKeyword(20)

	assert.Equal(t, []int{1, 5, 13, 15, 20}, p.GetKeywords())
}

func TestPost_Setter(t *testing.T) {
	title := "Post Title"
	subtitle := "This is a awesome subtitle"
	content := "<h1>Post</h1>\n<p>Test Content</p>"
	author := 1

	p := NewPost(title, subtitle, content, author)

	assert.Equal(t, title, p.Title)
	assert.Equal(t, subtitle, p.SubTitle)
	assert.Equal(t, content, p.Content)
	assert.Equal(t, author, p.GetAuthorID())

	newId := 15
	newTitle := "Another new special äüößß Title"
	newSubTitle := "This subtitle is great for testing äöüßß"
	newContent := "<div>I like this content</div>"
	newAuthor := 1234
	newKeywords := []int{1, 2, 3}
	newDate := time.Now()

	p.SetID(newId)
	p.SetTitle(newTitle)
	p.SetSubTitle(newSubTitle)
	p.SetContent(newContent)
	p.SetAuthor(newAuthor)
	p.SetKeywords(newKeywords)
	p.SetDate(newDate)

	assert.Equal(t, newId, p.Id)
	assert.Equal(t, newTitle, p.Title)
	assert.Equal(t, newSubTitle, p.SubTitle)
	assert.Equal(t, newContent, p.Content)
	assert.Equal(t, newAuthor, p.GetAuthorID())
	assert.Equal(t, newKeywords, p.GetKeywords())
	assert.Equal(t, newDate, p.GetDate())
}

