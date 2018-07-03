// Enthält alle Utils
package util

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/core/controller"
	"math/rand"
	"net/http"
	"time"
)

//Methode liefert einen zufälligen String in der übergebenen Länge zurück.
func RandomString(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//Methode lierfert einen Handler aus eine Controller
func UnitTest_GetHandler(ctrl *controller.Controller, path string) controller.HandlerFunc {
	return ctrl.GetHandler()[path]
}

//Methode lierfert einen Handler aus eine Controller und wrappt diesen korrekt für UnitTests
func UnitTest_GetHandlerWrapped(ctrl *controller.Controller, path string) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		handler := UnitTest_GetHandler(ctrl, path)
		handler(w, r)
	}
}