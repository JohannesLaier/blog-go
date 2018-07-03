package keyword

import (
	"testing"
	"de/blog/projekt/01/core/db"
	"github.com/stretchr/testify/assert"
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/app/models/post"
)

func TestKeyword_Relation(t *testing.T) {
	blog_posts := db.NewDBCollection("post")
	blog_authors := db.NewDBCollection("author")
	blog_keywords := db.NewDBCollection("keyword")

	blog_author1 := author.NewAuthor("peter", "geheim")
	blog_author2 := author.NewAuthor("peter", "geheim")
	blog_authors.Add(blog_author1)
	blog_authors.Add(blog_author2)

	blog_keyword1 := NewKeyword("Keyword1")
	blog_keyword2 := NewKeyword	("Keyword2")
	blog_keywords.Add(blog_keyword1)
	blog_keywords.Add(blog_keyword2)

	blog_post1 := post.NewPost("title", "subtitle", "content",  blog_author1.GetID())
	blog_post2 := post.NewPost("title", "subtitle", "content",  blog_author2.GetID())

	blog_post1.AddKeyword(blog_keyword1.GetID())
	blog_post1.AddKeyword(blog_keyword2.GetID())

	blog_post2.AddKeyword(blog_keyword1.GetID())

	blog_posts.Add(blog_post1)
	blog_posts.Add(blog_post2)

	assert.Equal(t, []int {blog_keyword1.GetID(), blog_keyword2.GetID()}, blog_post1.GetKeywords(), "Incorrect relation to keywords")
	assert.Equal(t, []int {blog_keyword1.GetID()}, blog_post2.GetKeywords(), "Incorrect relation to keywords")

}

func TestNewKeyword(t *testing.T) {
	keyword := NewKeyword("Keywordname")

	assert.NotEmpty(t, keyword)
}

func TestKeyword_GetID(t *testing.T) {
	collection := db.NewDBCollection("testKeyword")

	keyword := NewKeyword("test")

	collection.Add(keyword)

	assert.Equal(t, 1, collection.GetAll()[0].GetID(), "Get wrong keyword id")
}

func TestKeyword_SetID(t *testing.T) {
	keyword_id := 5
	keyword_name := "keyword_üöäßßß"

	keyword := NewKeyword(keyword_name)

	assert.Equal(t, 0, keyword.GetID(), "Get wrong keyword id")

	keyword.SetID(keyword_id)

	assert.Equal(t, keyword_id, keyword.GetID(), "Get wrong keyword id")
}

func TestKeyword_GetName(t *testing.T) {
	keyword_name := "keyword_üöäßßß"

	keyword := NewKeyword(keyword_name)

	assert.Equal(t, keyword_name, keyword.GetName())
}