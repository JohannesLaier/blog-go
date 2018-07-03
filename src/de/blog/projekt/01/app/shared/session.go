//Enthält alle Wrapper-Funktionalitäten, die von allen Controllern verwendet werden können.
package shared

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/core/http/session"
	"de/blog/projekt/01/core/controller"
	"net/http"
)

//Methode stellt einen Wrapper dar, der prüft ob ein Nutzer zur Zeit angemeldet ist.
func SessionWrapper(handler controller.HandlerFunc) controller.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		session_store := session.GetSessionStore()
		sess := session_store.GetCurrent(r)
		if sess != nil {
			handler(w, r)
		} else {
			http.Redirect(w, r, "/admin/login", 302)
		}
	}
}
