// Enthält die Methode zum PostController (Funktionalitäten zum Verwalten der Posts im Backend)
package admin_post

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/app/shared"
	"de/blog/projekt/01/app/models/comment"
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/core/http/session"
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/view"
	"de/blog/projekt/01/core/db"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"sort"
)

//Gibt ein PostController-Objekt, das die notwenigen Daten an die Post-Views übergibt und diese darstellt
// und alle Funktionalitäten bezüglich der Blog-Posts im Backend beinhaltet, zurück.
// Dieses enthält einen Handler für die Auflistung der Posts ("/admin/posts"), der auch das Löschen von Posts ermöglicht,
// und einen zweiten Handler ("/admin/posts-detail") für die Detail-Ansicht eines (bestehenden oder neuen) Posts,
// in dem die Eigenschaften (Titel, Untertitel, Text, Keywords) verändert
// und zugehörige Kommentare freigegeben und gelöscht werden können.
// Die Views stellen zum einen eine Auflistung aller Posts, sortiert nach Einstellungsdatum,
// und zum anderen eine Detailansicht eines Post, mit dem vollständigen Beitrag, dessen Kategorien, den Kommentaren,
// und ein Formular zum kommentieren dar.
func NewPostController() * controller.Controller {
	//Neuer Controller als "PostController"
	ctrl := controller.NewController("PostController")

	// Nur mit Session erreichbar machen
	ctrl.SetAuthWrapper(shared.SessionWrapper)

	// Handler für die Auflistung der Posts unter "/admin/posts"
	ctrl.AddHandler("/admin/posts", func (w http.ResponseWriter, r *http.Request) {

		// Aktuell eingeloggten Nutzer aus Session holen
		session_store := session.GetSessionStore()
		sess := session_store.GetCurrent(r)
		user := sess.Get("CURRENT")
		current := user.(*author.Author)

		// Parameter aus URL auslesen
		params := r.URL.Query()

		// Daten aus Datenbank holen
		blog_db := db.Get("blog")
		blog_posts := blog_db.GetCollection("post")

		// Funktionalitäten ausführen (Post löschen)
		action := template.HTMLEscapeString(params.Get("action"))
		if action == "delete" {
			post_id, _ := strconv.Atoi(params.Get("id"))
			blog_posts.Remove(post_id)
		}

		// Nur die eigenen Posts anzeigen
		posts := blog_posts.GetListFilter(func (entry db.DBCollectionEntry) bool {
			return entry.(*post.Post).GetAuthorID() == current.GetID()
		})

		// Posts absteigend nach Datum sortieren
		sort.Slice(posts, func(i, j int) bool {
			return (posts[i].(*post.Post)).GetDate().Unix() > (posts[j].(*post.Post)).GetDate().Unix()
		})

		blog_db.Store()

		// Daten an View übergeben und darstellen
		v := view.NewView("admin/views/post_list", "admin/layout/layout")
		v.SetModel(posts)
		v.Write(w)
	})

	// Handler für die Detailansicht neuer und bestehender Posts unter "/admin/post-admin"
	ctrl.AddHandler("/admin/post-detail", func (w http.ResponseWriter, r *http.Request) {

		// Aktuell eingeloggten Nutzer aus Session holen
		session_store := session.GetSessionStore()
		sess := session_store.GetCurrent(r)
		current := sess.Get("CURRENT").(*author.Author)

		// Parameter aus URL auslesen
		params := r.URL.Query()
		post_id, err := strconv.Atoi(params.Get("id"))
		if err != nil {
			post_id = 0
		}

		// Daten aus Datenbank holen
		blog_db := db.Get("blog")
		blog_posts := blog_db.GetCollection("post")
		blog_authors := blog_db.GetCollection("author")
		blog_keywords := blog_db.GetCollection("keyword")
		blog_comments := blog_db.GetCollection("comment")

		var blog_post *post.Post
		var blog_author *author.Author
		if (post_id > 0) { // Save: bestehenden Post in Post laden (als Übergabe für das View)
			blog_post = blog_posts.GetByID(post_id).(*post.Post)
			blog_author_id := blog_post.GetAuthorID()
			blog_author = blog_authors.GetByID(blog_author_id).(*author.Author)

			// Nur eigene Posts dürfen bearbeitet/gelesen werden
			if (blog_post.GetAuthorID() != current.GetID()) {
				http.Redirect(w, r, "posts", 302)
				return
			}
		} else { // New: neuen Post anlegen und in Post laden (als Übergabe für das View)
			blog_post = &post.Post{}
			blog_author = current
			blog_post.SetAuthor(current.GetID())
		}

		// Funktionalitäten ausführen (Post speichern oder löschen, Kommentar freigeben oder löschen)
		action := template.HTMLEscapeString(params.Get("action"))
		if action == "savePost" { // Post speichern
			post_title := template.HTMLEscapeString(r.FormValue("title"))
			post_subtitle := template.HTMLEscapeString(r.FormValue("subtitle"))
			post_content := r.FormValue("content")

			r.ParseForm()
			var post_keywords []int
			post_keywords_str := r.PostForm["keywords[]"]
			for _, k_id := range post_keywords_str {
				key_id, _ := strconv.Atoi(k_id)
				post_keywords = append(post_keywords, key_id)
			}

			blog_post.SetTitle(post_title)
			blog_post.SetSubTitle(post_subtitle)
			blog_post.SetContent(post_content)
			blog_post.SetKeywords(post_keywords)

			if (post_id > 0) { // Änderung an bestehendem Post speichern
				blog_posts.Update(blog_post)
			} else { // Neuen Post in DB anlegen
				blog_post.SetDate(time.Now())
				blog_posts.Add(blog_post)
			}

			//Nach Speichern zurück auf die Post-Übersicht
			http.Redirect(w, r, "posts", 302)
			return

		} else if action == "deletePost" {
			// Post aus DB löschen
			blog_posts.Remove(post_id)
			http.Redirect(w, r, "posts", 302)
			return

		} else if action == "shareComment" {
			// Kommentar freigeben
			comment_id, _ := strconv.Atoi(params.Get("comment_id"))
			comment := blog_comments.GetByID(comment_id).(*comment.Comment)
			comment.SetShare(true)
			blog_comments.Update(comment)

		} else if action == "deleteComment" {
			// Kommentar löschen
			comment_id, _ := strconv.Atoi(params.Get("comment_id"))
			blog_comments.Remove(comment_id)

		}

		// Keywords des Posts
		var keywords []map[string]interface{}
		for _ , entry := range blog_keywords.GetAll() {
			checked := false
			for _, id := range blog_post.GetKeywords() {
				if entry.GetID() == id {
					checked = true
				}
			}
			keywords = append(keywords, map[string]interface{} {
				"keyword": entry,
				"checked": checked,
			})
		}

		// Noch freizugebende Kommentare zu dem Post
		blog_post_comments_new := blog_comments.GetListFilter(func(entry db.DBCollectionEntry) bool {
			com := entry.(*comment.Comment)
			return (com.GetPost() == blog_post.GetID() && com.Share == false);
		})

		// Freigegebene Kommentare zu dem Post
		blog_post_comments := blog_comments.GetListFilter(func(entry db.DBCollectionEntry) bool {
			com := entry.(*comment.Comment)
			return (com.GetPost() == blog_post.GetID() && com.Share == true);
		})

		blog_db.Store()

		// Daten an View übergeben und darstellen
		v := view.NewView("admin/views/post_detail", "admin/layout/layout")
		v.SetModel(map[string]interface{} {
			"post" : blog_post,
			"author" : blog_author,
			"keywords" : keywords,
			"newComments" : blog_post_comments_new,
			"comments" : blog_post_comments,
		})
		v.Write(w)
	})

	// PostController mit den zwei Handlern übergeben
	return ctrl
}