// Enthält die Methode zum AboutController (Darstellen des About-Views im Frontend)
package blog_about

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/view"
	"net/http"
)

//Gibt ein AboutController-Objekt, das mithilfe des Handlers den About-View darstellt, zurück.
// Das View stellt zusätzliche Informationen über den Blog dar.
func NewAboutController() * controller.Controller {
	// Neuer Controller als "AboutController"
	ctrl := controller.NewController("AboutController")
	// Handler zum Anzeigen der About-Seite unter "/about"
	ctrl.AddHandler("/about", func (w http.ResponseWriter, r *http.Request) {
		// View darstellen
		v := view.NewView("blog/views/blog_about", "blog/layout/layout")
		v.Write(w)
	})

	// AboutController mit dem Handler übergeben
	return ctrl
}
