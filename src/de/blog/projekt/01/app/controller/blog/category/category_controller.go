// Enthält die Methode zum CategoryController (Darstellen des Category-Views mit den Posts der Kategorie im Frontend)
package blog_category

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/view"
	"de/blog/projekt/01/core/db"
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/app/models/keyword"
	"net/http"
	"strconv"
	"sort"
)

//Gibt ein CategoryController-Objekt, das die notwenigen Daten an das Category-View übergibt und dieses darstellt, zurück.
// Dieses enthält einen Handler("/admin/category"), der die Post zur gegebenen Kategorie und deren Autoren ermittelt.
// Das View stellt eine Auflistung aller Post der entsprechenden Kategorie dar, sortiert nach Einstellungsdatum.
func NewCategoryController() * controller.Controller {
	//Neuer Controller als "CategoryController"
	ctrl := controller.NewController("CategoryController")
	// Handler für die Auflistung der Posts der jeweiligen Kategorie unter "/category"
	ctrl.AddHandler("/category", func (w http.ResponseWriter, r *http.Request) {

		// Parameter aus URL auslesen
		params := r.URL.Query()
		id, _ := strconv.Atoi(params.Get("id"))

		// Daten aus Datenbank holen
		blog_db := db.Get("blog")
		blog_posts := blog_db.GetCollection("post")
		blog_authors := blog_db.GetCollection("author")
		blog_keywords := blog_db.GetCollection("keyword")

		// Posts der entsprechender Kategorie herausfiltern
		category_posts := []post.Post{}
		for _, p := range blog_posts.GetAll() {
			blog_post := *(p.(*post.Post))

			for _, keyword_id := range blog_post.GetKeywords() {
				if keyword_id == id {
					category_posts = append(category_posts, blog_post)
					break;
				}
			}
		}

		// Posts der entsprechender Kategorie nach Einstellungsdatum sortieren
		sort.Slice(category_posts, func(i, j int) bool {
			return category_posts[i].GetDate().Unix() > category_posts[j].GetDate().Unix()
		})

		// Autoren zu Posts der entsprechenden Kategorie
		var posts_authors []map[string]interface{}
		for  _, p := range category_posts {

			blog_author_id := p.GetAuthorID()
			blog_author := blog_authors.GetByID(blog_author_id)

			post_author := map[string]interface{} {
				"post" : p,
				"author" : blog_author,
			}

			posts_authors = append(posts_authors, post_author)
		}

		blog_keyword := *(blog_keywords.GetByID(id).(*keyword.Keyword))

		// Daten an View übergeben und darstellen
		v := view.NewView("blog/views/blog_category", "blog/layout/layout")
		v.SetModel(map[string]interface{} {
			"posts_authors" : posts_authors,
			"keyword" :     blog_keyword,
		})
		v.Write(w)
	})

	// CategoryController mit dem Handler übergeben
	return ctrl
}