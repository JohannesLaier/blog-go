// Enthält die Methode zum IndexController (Darstellen des Blog-Post-Views im Frontend)
package blog_index

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/view"
	"de/blog/projekt/01/core/db"
	"net/http"
	"sort"
)

//Gibt ein DetailController-Objekt, das mithilfe des Handlers ("/") die notwenigen Daten an den Blog-Posts-View übergibt
// und diesen darstellt, zurück.
// Das View stellt eine Auflistung aller Blogartikel mit dem Titel, Untertitel, Autor und Einstellungsdatum dar.
func NewIndexController() *controller.Controller {
	//Neuer Controller als "IndexController"
	ctrl := controller.NewController("IndexController")

	// Handler zum Anzeigen einer Auflistung aller Post mit dem Titel und Untertitel
	// sowie dem Autor und dem Einstellungsdatum
	ctrl.AddHandler("/", func (w http.ResponseWriter, r *http.Request) {

		// Daten aus Datenbank holen
		blog_db := db.Get("blog")
		blog_posts := blog_db.GetCollection("post")
		blog_authors := blog_db.GetCollection("author")

		posts := blog_posts.GetAll()

		// Sortieren aller Posts nach Einstellungsdatum
		sort.Slice(posts, func(i, j int) bool {
			return (posts[i].(*post.Post)).GetDate().Unix() > (posts[j].(*post.Post)).GetDate().Unix()
		})

		// Autoren zu Posts ermitteln
		var posts_authors []map[string]interface{}
		for  _, p := range posts {

			blog_post := p.(*post.Post)
			blog_author_id := blog_post.GetAuthorID()
			blog_author := blog_authors.GetByID(blog_author_id)

			article := map[string]interface{} {
				"post" : blog_post,
				"author" : blog_author,
			}

			posts_authors = append(posts_authors, article)
		}

		// Daten an View übergeben und darstellen
		v := view.NewView("blog/views/blog_posts", "blog/layout/layout")
		v.SetModel(posts_authors)
		v.Write(w)
	})
	// IndexController mit dem Handler übergeben
	return ctrl
}
