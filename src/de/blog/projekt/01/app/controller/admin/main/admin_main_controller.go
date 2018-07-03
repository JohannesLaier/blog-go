// Enthält die Methode zum AdminMainController (Funktionalität zum Weiterleiten ins Backend)
package admin_main

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/core/controller"
	"net/http"
)

//Gibt ein AdminMainController-Objekt,
// das die Funktionalität zum richtigen Weiterleiten zum Backend beinhaltet, zurück.
// Dieses enthält einen Handler ("/admin") und einen Handler ("/admin/"),
// der auf die Post-Übersicht im Backend weiterleitet.
func NewAdminMainController() * controller.Controller {
	//Neuer Controller als "AdminMainController"
	ctrl := controller.NewController("AdminMainController")

	// Handler für Weiterleitung von"/admin/" zu "/admin/login"
	ctrl.AddHandler("/admin/", func (w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/admin/posts", 302)
	})

	// Handler für Weiterleitung von"/admin/" zu "/admin/login"
	ctrl.AddHandler("/admin", func (w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/admin/posts", 302)
	})

	// AdminMainController mit den Handlern übergeben
	return ctrl
}
