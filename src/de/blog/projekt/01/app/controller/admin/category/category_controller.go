// Enthält die Methode zum CategoryController (Funktionalitäten zum Verwalten der Kategorien)
package admin_categorie

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/app/shared"
	"de/blog/projekt/01/app/models/keyword"
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/view"
	"de/blog/projekt/01/core/db"
	"html/template"
	"net/http"
	"strconv"
	"sort"
	"strings"
)

//Gibt ein CategoryController-Objekt, das die notwenigen Daten an den Categories-View übergibt und diesen darstellt
// und alle Funktionalitäten bezüglich der Kategorien im Backend beinhaltet, zurück.
// Dieses enthält einen Handler für die Auflistung der Posts ("/admin/categories"),
// der auch das Erstellen und Löschen von Kategorien ermöglicht.
// Das View stellt eine Auflistung der Kategorien und ein Formular zu Erstellung neuer Kategorien dar.
func NewCategoryController() * controller.Controller {
	//Neuer Controller als "CategoryController"
	ctrl := controller.NewController("CategoryController")

	// Nur mit Session erreichbar machen
	ctrl.SetAuthWrapper(shared.SessionWrapper)

	// Handler für die Auflistung der Kategorien unter "/admin/categories"
	ctrl.AddHandler("/admin/categories", func (w http.ResponseWriter, r *http.Request) {

		//Statusmeldung-Flags
		success_save := false
		success_delete := false

		// Parameter aus URL auslesen
		params := r.URL.Query()

		// Parameter aus URL auslesen
		blog_db := db.Get("blog")
		blog_posts := blog_db.GetCollection("post")
		blog_keywords := blog_db.GetCollection("keyword")

		// Funktionalitäten ausführen (neue Kategorie erstellen, bestehende Kategorie löschen)
		action := template.HTMLEscapeString(params.Get("action"))
		if action == "add" { // Neue Kategorie erstellen

			keyword_text := template.HTMLEscapeString(r.FormValue("category"))
			keyword_new := keyword.NewKeyword(keyword_text)
			blog_keywords.Add(keyword_new)
			success_save = true

		} else if action == "delete" { // Kategorie löschen

			keyword_id, _ := strconv.Atoi(params.Get("id"))
			blog_keywords.Remove(keyword_id)
			success_delete = true

			// Keyword-Relation aus Posts entfernen
			for  _, p := range blog_posts.GetAll() {

				blog_post := p.(*post.Post)
				blog_post_keywords := blog_post.GetKeywords()

				for index, key_id := range blog_post_keywords {
					if (key_id == keyword_id) {
						blog_post_keywords = append(blog_post_keywords[:index], blog_post_keywords[index+1:]...)
					}
				}

				blog_post.SetKeywords(blog_post_keywords)
				blog_posts.Update(blog_post)
			}


		}

		keywords := blog_keywords.GetAll()

		// Alphabetisches Sortieren der Kategorien
		sort.Slice(keywords, func(i, j int) bool {
			keywordA := keywords[i].(*keyword.Keyword).GetName()
			keywordB := keywords[j].(*keyword.Keyword).GetName()

			return strings.Compare(keywordA, keywordB) == -1
		})

		blog_db.Store()

		// Daten an View übergeben und darstellen
		v := view.NewView("admin/views/categorie_list", "admin/layout/layout")
		v.SetModel(map[string]interface{}{
			"keywords": keywords,
			"success_save":  success_save,
			"success_delete":  success_delete,
		})
		v.Write(w)
	})

	// AuthorController mit dem Handler übergeben
	return ctrl
}