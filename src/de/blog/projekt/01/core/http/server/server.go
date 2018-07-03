// Enthält das Struct Server und alle zugehörigen Methoden für einen Server.
package server

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/core/controller"
	"de/blog/projekt/01/core/config"
	"net/http"
	"strconv"
	"log"
	"os"
)

//Das Struct Server definiert den Server
// mit einer Port-Angabe, Private(KeyFile) und Public-Key(CertFile) zur Verschlüsselung
// und einer Liste von Controllern.
type Server struct {
	PortHttps uint16
	KeyFile string
	CertFile  string
	controller []controller.Controller
}

//Konstruktor für ein neues Server-Objekt
// mit der Port-Angabe, dem Private und Public-Key aus dem übergebenen Configurations-Objekt
func NewServer(config config.Config) *Server {
	server := new(Server)
	server.PortHttps = uint16(config.GetHttpsPort())
	server.KeyFile = config.GetResourceFolder() + "ssl/key.pem"
	server.CertFile = config.GetResourceFolder() + "ssl/cert.pem"

	_, err := os.Stat(server.KeyFile)
	if err != nil {
		panic("SSL Key File could not be found.");
	}

	_, err = os.Stat(server.CertFile)
	if err != nil {
		panic("SSL Cert File could not be found.");
	}

	return server
}

//Methode fügt den übergebenen Controller zur Controller-Liste des Servers hinzu.
func (s *Server) AddController(controller * controller.Controller) {
	s.controller = append(s.controller, *controller)
}

//Methode gibt dem Server das Pattern an,
// unter dem die (statischen) Dateien aus dem Verzeichnis(dir) aufgerufen werden können.
func (s *Server) AddFileServer(pattern, dir string) {
	http.Handle(pattern, http.StripPrefix(pattern, http.FileServer(http.Dir(dir))))
}

//Methode startet den Server
func (s *Server) Run() {

	log.Println("[Server] Initializing... ")

	for _, ctrl := range s.controller {
		log.Println("[Server] Register Controller: " + ctrl.Name)

		for path, hFunc := range ctrl.GetHandler() {
			http.HandleFunc(path, hFunc)

			log.Printf("[Server] Register Method: %s %s\n", ctrl.Name, path)
		}
	}

	log.Printf("[Server] Server is running on port: %d\n", s.PortHttps)
	log.Printf("[Server] Open the following urls in your browser:")
	log.Printf("[Server] Frontend: https://127.0.0.1:%d, Backend: https://127.0.0.1:%d/admin/", s.PortHttps, s.PortHttps)

	// Go Server starten und auf Port lauschen
	err := http.ListenAndServeTLS(":" + strconv.Itoa(int(s.PortHttps)), s.CertFile, s.KeyFile, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}