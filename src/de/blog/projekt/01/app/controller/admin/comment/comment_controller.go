// Enthält die Methode zum CommentController (Funktionalitäten zum Verwalten der Kommentare im Backend)
package admin_comment

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/app/shared"
	"de/blog/projekt/01/app/models/comment"
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/view"
	"de/blog/projekt/01/core/db"
	"html/template"
	"net/http"
	"strconv"
	"de/blog/projekt/01/app/models/post"
	"de/blog/projekt/01/core/http/session"
	"de/blog/projekt/01/app/models/author"
	"sort"
)

//Gibt ein CommentController-Objekt, das die notwenigen Daten an den CommentView(Backend) übergibt und diesen darstellt
// und alle Funktionalitäten bezüglich der Kommentare im Backend beinhaltet, zurück.
// Dieses enthält einen Handler für die Auflistung der Kommentare ("/admin/comments")
// aufgeteilt in freizugebende und freigegebene Kommentrare, der auch das Freigeben und Löschen von Kommentaren ermöglicht.
// Das View stellt eine Auflistung aller freizugebenden und freigegebene Kommentare dar.
func NewCommentController() * controller.Controller {

	//Neuer Controller als "CommentController"
	ctrl := controller.NewController("CommentController")

	// Nur mit Session erreichbar machen
	ctrl.SetAuthWrapper(shared.SessionWrapper)

	// Handler für die Auflistung der freizugebenden und freigegebenen Kommentare unter "/admin/comments"
	ctrl.AddHandler("/admin/comments", func (w http.ResponseWriter, r *http.Request) {

		// Aktuell eingeloggten Nutzer aus Session holen
		session_store := session.GetSessionStore()
		sess := session_store.GetCurrent(r)
		user := sess.Get("CURRENT")
		current := user.(*author.Author)

		// Parameter aus URL auslesen
		params := r.URL.Query()

		blog_db := db.Get("blog")
		blog_comments := blog_db.GetCollection("comment")
		blog_posts := blog_db.GetCollection("post")

		// Funktionalitäten ausführen (Kommentar freigeben oder löschen)
		action := template.HTMLEscapeString(params.Get("action"))
		if action == "share" { // Kommentar freigeben
			comment_id, _ := strconv.Atoi(params.Get("id"))
			comment := blog_comments.GetByID(comment_id).(*comment.Comment)
			comment.SetShare(true)
			blog_comments.Update(comment)
		} else if action == "delete" { // Kommentar löschen
			comment_id, _ := strconv.Atoi(params.Get("id"))
			blog_comments.Remove(comment_id)
		}

		// Posts des aktuell eingeloggten Autors
		posts := blog_posts.GetListFilter(func (entry db.DBCollectionEntry) bool {
			return entry.(*post.Post).GetAuthorID() == current.GetID()
		})

		// Kommentare zu Posts des aktuell eingeloggten Autors herausfinden
		// und in Freizugebende und Freigegebene sortieren
		comments := blog_comments.GetAll()

		sort.Slice(comments, func(i, j int) bool {
			return comments[i].(*comment.Comment).GetDate().Unix() > comments[j].(*comment.Comment).GetDate().Unix()
		})

		comments_old := make([]map[string]interface{}, 0)
		comments_new := make([]map[string]interface{}, 0)

		for _, entry := range comments {
			comment := *(entry.(*comment.Comment))
			comment_post_id := comment.GetPost()

			// Nur Comments zu Posts des aktuell eingeloggten Autors
			for _, entry_post := range posts {
				if entry_post.GetID() == comment_post_id {

					entry := blog_posts.GetByID(comment_post_id)
					comment_post := entry.(*post.Post)

					comment_map := map[string]interface{} {
						"comment" : comment,
						"post" : comment_post,
					}

					// Freizugebendes Kommentar in die Liste der neuen Kommentare einfügen
					if comment.Share == false {
						comments_new = append(comments_new, comment_map)
					} else {
						comments_old = append(comments_old, comment_map)
					}

					break
				}
			}
		}

		blog_db.Store()

		// Daten an View übergeben und darstellen
		v := view.NewView("admin/views/comment_list", "admin/layout/layout")
		v.SetModel(map[string]interface{} {
			"comments" : comments_old,
			"newComments" : comments_new,
		})
		v.Write(w)
	})

	// CommentController mit dem Handler übergeben
	return ctrl
}