// Enthält das Struct Controller und den Typ der HandlerFunktion.
package controller

// Go-Blog von 8892993, 1734394, 1777093

import (
	"net/http"
)

//Typ-Definition einer Handler-Funktion
type HandlerFunc func (w http.ResponseWriter, r *http.Request)

//Typ-Definition einer Handler-Funktion die als Wrapper verwendet werden kann
type HandlerFuncWrapper func (f HandlerFunc) HandlerFunc

//Das Struct Controller definiert den Controller.
// Dieser hat einen Namen und beinhaltet eine Handler-Map mit HandlerFunktionen unt Pfadnamen
type Controller struct {
	Name string
	Handler map[string]HandlerFunc
	AuthWrapper HandlerFuncWrapper
}

//Kontruktor für ein neues Controller-Objekt
// mit dem übergebenen Namen und einer initialisierten Handler-Map für HandlerFunktionen.
func NewController(name string) * Controller {
	ctrl := new(Controller)
	ctrl.Name = name
	ctrl.Handler = make(map[string]HandlerFunc)
	return ctrl
}

//Methode fügt die übergebene HandlerFunktion unter dem übergebenen Pfad-String in die Handler-Map ein.
func (c * Controller) AddHandler(path string, hFunc HandlerFunc) {
	c.Handler[path] = hFunc
}

// Methode liefert die Handler-Map des Controllers zurück.
func (c * Controller) GetHandler() map[string]HandlerFunc {
	if c.AuthWrapper != nil {
		handlers := make(map[string]HandlerFunc)
		for path, handler := range c.Handler {
			handlers[path] = c.AuthWrapper(handler)
		}
		return handlers
	}
	return c.Handler
}

// Methode setzt den AuthWrapper für einen Controller
func (c * Controller) SetAuthWrapper(wrapper HandlerFuncWrapper) {
	c.AuthWrapper = wrapper
}