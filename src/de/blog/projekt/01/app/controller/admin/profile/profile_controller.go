// Enthält die Methode zum ProfileController (Funktionalitäten zum Verwalten der Profil-Einstellungen des eingeloggten Autors)
package admin_profile

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/app/models/author"
	"de/blog/projekt/01/app/shared"
	"de/blog/projekt/01/core/http/session"
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/view"
	"de/blog/projekt/01/core/db"
	"net/http"
	"strings"

	"html/template"
)

//Gibt ein ProfileController-Objekt, das die notwenigen Daten an den Profile-View übergibt und diesen darstellt
// und die Funktionalitäten zum Ändern der Profil-Einstellungen beinhaltet, zurück.
// Dieses enthält einen Handler ("/admin/profile"), der das Updaten der Profil-Einstellungen ermöglicht
// Das View stellt ein Formular zum Ändern der Profil-Einstellungen des eingeloggten Autors dar.
func NewProfileController() * controller.Controller {
	//Neuer Controller als "ProfileController"
	ctrl := controller.NewController("ProfileController")

	// Nur mit Session erreichbar machen
	ctrl.SetAuthWrapper(shared.SessionWrapper)

	// Handler für die Profil-Einstellungen unter "/admin/profile"
	ctrl.AddHandler("/admin/profile", func (w http.ResponseWriter, r *http.Request) {

		//Statusmeldung-Flags
		success := false
		error_empty := false
		error_password_invalid := false
		error_passwords_doesnt_match := false

		// Aktuell eingeloggten Nutzer aus Session holen
		session_store := session.GetSessionStore()
		sess := session_store.GetCurrent(r)
		user := sess.Get("CURRENT").(*author.Author)

		// Parameter aus URL auslesen
		params := r.URL.Query()

		// Funktionalitäten ausführen (Profildaten speichern)
		action := template.HTMLEscapeString(params.Get("action"))
		if action == "save" {

			// Textfeld-Eingaben holen
			user_name := template.HTMLEscapeString(r.FormValue("username"))
			user_password_old := template.HTMLEscapeString(r.FormValue("password"))
			user_password_new := template.HTMLEscapeString(r.FormValue("newPassword"))
			user_password_new2 := template.HTMLEscapeString(r.FormValue("newPassword2"))

			// Überprüfen ob Passwort korrekt
			if user.Verify(user_password_old) {
				// Prüfen ob Eingaben des neuen Passworts übereinstimmen
				if user_password_new == user_password_new2 {
					// Überprüfen, ob alle Felder ausgefüllt sind
					if (strings.TrimSpace(user_name) != "") && (strings.TrimSpace(user_password_new) != "") {
						blog_db := db.Get("blog")
						blog_authors := blog_db.GetCollection("author")

						user.SetUsername(user_name)
						user.SetPassword(user_password_new)

						blog_authors.Update(user)
						blog_db.Store()

						success = true
					} else {
						error_empty = true
					}
				} else {
					error_passwords_doesnt_match = true
				}
			} else {
				error_password_invalid = true
			}
		}

		// Daten an View übergeben und darstellen
		v := view.NewView("admin/views/profile", "admin/layout/layout")
		v.SetModel(map[string]interface{} {
			"user" : user,
			"success" : success,
			"error_empty" : error_empty,
			"error_password_invalid" : error_password_invalid,
			"error_passwords_doesnt_match" : error_passwords_doesnt_match,
		})
		v.Write(w)
	})

	// ProfileController mit dem Handler übergeben
	return ctrl
}
