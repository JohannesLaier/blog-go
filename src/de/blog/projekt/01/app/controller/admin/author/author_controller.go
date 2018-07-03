// Enthält die Methode zum AuthorController (Funktionalitäten zum Verwalten der Autoren)
package admin_author

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/app/shared"
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/core/http/session"
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/view"
	"de/blog/projekt/01/core/db"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"sort"
)

//Gibt ein neues AuthorController-Objekt, das die notwenigen Daten an die Author-Views übergibt und diese darstellt
// und die Funktionalitäten zum Verwalten der Blog-Autoren beinhaltet, zurück.
// Dieses enthält einen Handler für die Auflistung der Autoren ("/admin/authors"),
// der auch das Löschen von Autoren ermöglicht,
// und einen zweiten Handler ("/admin/author-detail") zum Erstellen eines neuen Autors.
// Die Views stellen zum einen eine Auflistung aller Autoren dar,
// und zum anderen ein Formular zum Anlegen eines neuen Autoren.
func NewAuthorController() * controller.Controller {
	//Neuer Controller als "AuthorController"
	ctrl := controller.NewController("AuthorController")

	// Nur mit Session erreichbar machen
	ctrl.SetAuthWrapper(shared.SessionWrapper)

	// Handler für die Auflistung der Autoren unter "/admin/authors"
	ctrl.AddHandler("/admin/authors", func (w http.ResponseWriter, r *http.Request) {
		//Statusmeldungs-Flag
		success_delete := false

		// Parameter aus URL auslesen
		params := r.URL.Query()

		// Daten aus Datenbank holen
		blog_db := db.Get("blog")
		blog_authors := blog_db.GetCollection("author")
		blog_posts := blog_db.GetCollection("post")

		// Aktuell eingeloggten Nutzer aus Session holen
		session_store := session.GetSessionStore()
		sess := session_store.GetCurrent(r)
		current := sess.Get("CURRENT").(*author.Author)

		// Funktionalitäten ausführen (löschen)
		action := template.HTMLEscapeString(params.Get("action"))
		if action == "delete" {
			author_id, _ := strconv.Atoi(params.Get("id"))
			blog_authors.Remove(author_id)
			success_delete = true

			// Blog Artikel des Authors an anderen Admin (eingeloggten Nutzer) übertragen
			for _, entry := range blog_posts.GetAll() {
				blog_post := entry.(*post.Post)
				if blog_post.GetAuthorID() == current.GetID() {
					blog_post.SetAuthor(current.GetID())
					blog_posts.Update(blog_post)
				}
			}
		}

		// Eingeloggten Nutzer selbst aus der Liste der Autoren entfernen, so dass man sich selbst nicht löschen kann
		var authors []author.Author
		all_authers := blog_authors.GetAll()
		for _, entry := range all_authers {
			if (current.GetID() != entry.GetID()) {
				authors = append(authors, *(entry.(*author.Author)))
			}
		}

		// Alphabetisches Sortieren der Autoren
		sort.Slice(authors, func(i, j int) bool {
			authorA := authors[i].GetUsername()
			authorB := authors[j].GetUsername()

			return strings.Compare(authorA, authorB) == -1
		})

		blog_db.Store()

		// Daten an View übergegben und darstellen
		v := view.NewView("admin/views/author_list", "admin/layout/layout")
		v.SetModel(map[string]interface{}{
			"authors": authors,
			"success_delete":  success_delete,
		})
		v.Write(w)
	})

	// Handler für die Detailansicht zum Erstellen neuer Autoren unter "/admin/author-admin"
	ctrl.AddHandler("/admin/author-detail", func (w http.ResponseWriter, r *http.Request) {

		//Statusmeldungs-Flag
		error_username := false
		error_empty := false

		// Parameter aus URL auslesen
		params := r.URL.Query()

		// Daten aus Datenbank holen
		blog_db := db.Get("blog")
		blog_authors := blog_db.GetCollection("author")

		// Funktionalitäten ausführen (neuen Autor erstellen)
		action := template.HTMLEscapeString(params.Get("action"))
		if action == "add" {
			author_username := template.HTMLEscapeString(r.FormValue("username"))
			author_password := template.HTMLEscapeString(r.FormValue("password"))

			if !((strings.TrimSpace(author_username) != "") && (strings.TrimSpace(author_password) != "")) {
				error_empty = true
			} else {

				// Überprüfen ob angegebener Username bereits existiert
				username_exist := blog_authors.GetFilter(func(entry db.DBCollectionEntry) bool {
					author_entry := entry.(*author.Author)
					return (author_entry.GetUsername() == author_username);
				})

				if username_exist != nil {
					error_username = true
				} else {
					//Wenn er noch nicht existiert: Neuen Autor mit gegebenen Werten anlegen
					blog_author := author.NewAuthor(author_username, author_password)
					blog_authors.Add(blog_author)

					http.Redirect(w, r, "authors", 302)
					return
				}
			}
		}

		blog_db.Store()

		// Daten an View übergeben und darstellen
		v := view.NewView("admin/views/author_detail", "admin/layout/layout")
		v.SetModel(map[string]interface{} {
			"error_username" : error_username,
			"error_empty" : error_empty,

		})
		v.Write(w)
	})

	// AuthorController mit den zwei Handlern übergeben
	return ctrl
}