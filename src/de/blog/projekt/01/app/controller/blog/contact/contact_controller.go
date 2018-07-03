// Enthält die Methode zum ContactController (Darstellen des Contact-Views im Frontend)
package blog_contact

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/view"
	"net/http"
)

//Gibt ein ContactController-Objekt, das mithilfe des Handlers den Contact-View darstellt, zurück.
// Das View stellt Kontaktinformationen dar.
func NewContactController() * controller.Controller {
	// Neuer Controller als "ContactController"
	ctrl := controller.NewController("ContactController")
	// Handler zum Anzeigen der Contact-Seite unter "/contact"
	ctrl.AddHandler("/contact", func (w http.ResponseWriter, r *http.Request) {
		v := view.NewView("blog/views/blog_contact", "blog/layout/layout")
		v.Write(w)
	})
	// ContactController mit dem Handler übergeben
	return ctrl
}
