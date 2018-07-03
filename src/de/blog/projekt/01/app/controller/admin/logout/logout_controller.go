// Enthält die Methode zum LogoutController (Funktionalität zum Ausloggen als Autor vom Backend)
package admin_logout

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/http/session"
	"net/http"
)

//Gibt ein LogoutController-Objekt, das die Funktionalitäten zum Abmelden
// der Blog-Autoren vom Backend beinhaltet, zurück.
// Dieses enthält einen Handler ("/admin/logout"), der den aktuell eingeloggten Nutzer abmeldet.
func NewLogoutController() * controller.Controller {
	//Neuer Controller als "LogoutController"
	ctrl := controller.NewController("LogoutController")

	// Handler für Logout-Funktionalität als Autor aus dem Backend unter "/admin/logout"
	ctrl.AddHandler("/admin/logout", func (w http.ResponseWriter, r *http.Request) {

		// Aktuell eingeloggten Nutzer aus Session holen und entfernen
		session_store := session.GetSessionStore()
		sess := session_store.GetCurrent(r)
		if sess != nil {
			sess.Remove("CURRENT")
			sess.DestroyCookie(w, r)
			session_store.Store()
		}

		// Auf Login zum Backend weiterleiten
		http.Redirect(w, r, "login", 302)
	})

	// LogoutController mit dem Handler übergeben
	return ctrl
}