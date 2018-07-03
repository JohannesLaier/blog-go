// Enthält die Methode zum LoginController (Funktionalität zum Einloggen als Autor im Backend)
package admin_login

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/core/http/session"
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/view"
	"de/blog/projekt/01/core/db"
	"html/template"
	"net/http"
)

//Gibt ein LoginController-Objekt, das die notwenigen Daten an den Login-View übergibt und diesen darstellt
// und die Funktionalitäten zum Anmelden der Blog-Autoren ins Backend beinhaltet, zurück.
// Dieses enthält einen Handler ("/admin/login"), der die Eingaben überprüft und evtl. den Autor einloggt.
// Das View stellt ein Formular zum Einloggen im Backend dar.
func NewLoginController() * controller.Controller {
	//Neuer Controller als "LoginController"
	ctrl := controller.NewController("LoginController")

	// Handler für Login-Funktionalität als Autor ins Backend unter "/admin/login"
	ctrl.AddHandler("/admin/login", func (w http.ResponseWriter, r *http.Request) {
		//Statusmeldungs-Flag
		error := false

		// Parameter aus URL auslesen
		r.ParseForm()
		username := template.HTMLEscapeString(r.FormValue("username"))
		pwd := template.HTMLEscapeString(r.FormValue("password"))

		// Überprüfung, dass alle Felder ausgefüllt sind
		if username != "" && pwd != "" {
			// Daten aus Datenbank holen
			blog_db := db.Get("blog")
			blog_authors := blog_db.GetCollection("author")

			// Überprüfen ob angegebener Autor existiert
			blog_author := blog_authors.GetFilter(func(entry db.DBCollectionEntry) bool {
				author_entry := entry.(*author.Author)
				return (author_entry.GetUsername() == username);
			})

			if (blog_author != nil) {
				user := blog_author.(*author.Author)
				// Überprüfen ob Passwort korrekt
				if (user.Verify(pwd)) {
					session_store := session.GetSessionStore()
					_, sess := session_store.New()
					sess.Put("CURRENT", user)
					sess.CreateCookie(w)
					session_store.Store()

					// Auf Auflistung der Post im Backend weiterleiten
					http.Redirect(w, r, "posts", 302)
					return
				}
			}
			error = true
		}

		// Daten an View übergeben und darstellen
		v := view.NewView("admin/views/login", "admin/layout/layout_login")
		v.SetModel(map[string]interface{} {
			"error" : error,
		})
		v.Write(w)
	})

	// LoginController mit dem Handler übergeben
	return ctrl
}
