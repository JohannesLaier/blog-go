//Enthält das View-Struct und alle Methoden zum Aufbau einer HTML-Methode aus Modell, View und Layout.
package view

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/core/config"
	"text/template"
	"net/http"
	"bytes"
	"time"
	"path"
	"log"
)

//Verzeichnis der Templates
const tpl_directory  = "tpl/"

//Das Struct View definiert eine HTML-Seite
// durch das Modell und die Pfade zum View und dem Layout.
// Es enthält außerdem eine FunctionMap.
type View struct {
	PathView string
	PathLayout string
	FuncMap template.FuncMap
	Model interface{}
}

//Konstruktor für ein neues View-Objekt
// mit den übergebenen Pfaden zum View und dem Layout
// und einer FuncMap, die das Datumsformat beinhaltet.
func NewView(pathView, pathLayout string) * View {
	view := new(View)
	view.PathView = pathView
	view.PathLayout = pathLayout
	view.FuncMap = make(template.FuncMap)
	view.FuncMap["date"] = func(date time.Time) string {
		return date.Format("02.01.2006 15:04")
	}
	return view
}

//Methode setzt das übergebene Model als Model des Views.
func (view *View) SetModel(model interface{}) {
	view.Model = model
}

//Methode fügt in die FuncMap die übergebene Funktion unter dem übergebenen Namen hinzu.
func (view *View) AddFunc(name string, fun interface{}) {
	view.FuncMap[name] = fun
}

//Methode fügt die Modells, das View und das Layout zu einer kompletten HTML-Seite zusammen
// und übergibt diese dem HTTP-Response-Writer
func (view *View) Write(writer http.ResponseWriter) {
	//Pfade zum HTML des Views und Layouts ermitteln
	config := config.GetConfig()
	tpl_dir := config.GetResourceFolder() + tpl_directory
	tpl_path_view := tpl_dir + view.PathView + ".tpl"
	tpl_path_layout := tpl_dir + view.PathLayout + ".tpl"

	//View-HTML aus Datei auslesen ...
	t_view, error := template.New(path.Base(tpl_path_view)).Funcs(view.FuncMap).ParseFiles(tpl_path_view)
	if error != nil {
		log.Fatal(error)
		log.Fatal("[View] Could not find template (view): " + tpl_path_view)
		return
	}

	var tpl bytes.Buffer

	// ... und mit den Daten des Modells befüllen
	if error := t_view.Execute(&tpl, view.Model); error != nil {
		log.Fatal(error)
		return
	}

	//Komplettes HTML des Views als String als Inhalt für das Layout
	content_view := tpl.String()

	//Layout-HTML aus Datei laden ...
	t_layout, error := template.New(path.Base(tpl_path_layout)).Funcs(view.FuncMap).ParseFiles(tpl_path_layout)
	if error != nil {
		log.Fatal(error)
		log.Fatal("[View] Could not find template (layout): " + tpl_path_layout)
		return
	}

	// ... und mit dem View als Inhalt befüllen.
	// Komplettes HTML des Seite (Layout+View+Modell) an HTTP-Response-Writer übergeben
	t_layout.Execute(writer, map[string]string{
		"content" : content_view,
	})
}

