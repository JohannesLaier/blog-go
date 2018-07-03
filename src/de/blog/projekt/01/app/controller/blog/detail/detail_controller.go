// Enthält die Methode zum DetailController (Darstellen des Post-Detail-Views im Frontend)
package blog_detail

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/app/models/comment"
	"de/blog/projekt/01/app/models/keyword"
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/view"
	"de/blog/projekt/01/core/db"
	"html/template"
	"net/http"
	"strconv"
	"sort"
)

//Gibt ein DetailController-Objekt, das mithilfe des Handlers die notwenigen Daten an den Post-Detail-View übergibt
// und diesen darstellt, zurück. Der Handler ("/detail") ermöglicht außerdem das Verfassen und Speichern neuer Kommentare.
// Das View stellt einen Blogartikel detailliert mit Titel, Untertitel, Beitrag, Kategorien, Kommentare und
// einem Formular zum Verfassen neuer Kommentaren dar.
func NewDetailController() * controller.Controller {
	//Neuer Controller als "DetailController"
	ctrl := controller.NewController("DetailController")

	// Handler zum Anzeigen des Posts, der Kategorien und der Kommentare und zum Erstellen neuer Kommentare
	ctrl.AddHandler("/detail", func (w http.ResponseWriter, r *http.Request) {

		//Statusmeldung-Flag
		success := false

		// Parameter aus URL auslesen
		params := r.URL.Query()
		post_id, _ := strconv.Atoi(params.Get("id"))

		// Daten aus Datenbank holen
		blog_db := db.Get("blog")
		blog_posts := blog_db.GetCollection("post")
		blog_authors := blog_db.GetCollection("author")
		blog_keywords := blog_db.GetCollection("keyword")
		blog_comments := blog_db.GetCollection("comment")

		// Evtl. vorhandener Nickname des Kommentators aus Cookie lesen
		nickname_cookie, error := r.Cookie("NICKNAME")
		var nickname string
		if (error == nil) {
			nickname = template.HTMLEscapeString(nickname_cookie.Value)
		}

		// Funktionalitäten ausführen (Kommentar hinzufügen)
		action := template.HTMLEscapeString(params.Get("action"))
		if action == "comment" {
			comment_text := template.HTMLEscapeString(r.FormValue("text"))
			comment_username := template.HTMLEscapeString(r.FormValue("username"))

			blog_comment := comment.NewComment(comment_username, comment_text, post_id)
			blog_comments.Add(blog_comment)

			// Nickname des Kommentators in Cookie speichern
			cookie := http.Cookie{Name: "NICKNAME", Value: comment_username}
			http.SetCookie(w, &cookie)
			nickname = comment_username

			success = true
		}

		// Post und Autorinformationen aus DB-Daten
		blog_post := (blog_posts.GetByID(post_id)).(*post.Post)
		blog_author_id := blog_post.GetAuthorID()
		blog_author := blog_authors.GetByID(blog_author_id)

		// Kategorie-Namen aus DB-Daten
		keywords := []keyword.Keyword{}
		for _, id := range blog_post.GetKeywords() {
			keyword := *(blog_keywords.GetByID(id).(*keyword.Keyword))
			keywords = append(keywords, keyword)
		}

		// Freigegebene Kommentare zum Post
		blog_post_comments := blog_comments.GetListFilter(func(entry db.DBCollectionEntry) bool {
			com := entry.(*comment.Comment)
			return (com.GetPost() == blog_post.GetID() && com.Share == true);
		})

		sort.Slice(blog_post_comments, func(i, j int) bool {
			return (blog_post_comments[i].(*comment.Comment)).GetDate().Unix() > (blog_post_comments[j].(*comment.Comment)).GetDate().Unix()
		})

		blog_db.Store()

		// Daten an View übergeben und darstellen
		v := view.NewView("blog/views/blog_detail", "blog/layout/layout")
		v.SetModel(map[string]interface{} {
			"nickname" : nickname,
			"post" : blog_post,
			"author" : blog_author,
			"keywords" : keywords,
			"comments" : blog_post_comments,
			"success" : success,
		})
		v.Write(w)
	})

	// DetailController mit dem Handler übergeben
	return ctrl
}
