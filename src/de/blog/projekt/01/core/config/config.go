//Enthält das Config-Struct und die zugehörigen Methoden.
package config

// Go-Blog von 8892993, 1734394, 1777093

import (
	"flag"
	"os"
)

//Das Struct Confi definiert die Konfigurationen
// mit dem Pfad zum Resource-Folder,
// dem HTTP-Port unter dem der Blog erreichbar ist
// und der Sessionlebensdauer in Minuten, nach der die Autoren-Sessions am Backend des Blogs ablaufen.
type Config struct {
	resource_folder string
	https_port int
	session_expire int
}

//Instanz des Config-Structs (Singleton-Pattern)
var _config *Config;

//Methode liefert das Config-Objekt zurück.
func GetConfig() Config {
	if (_config == nil) {
		_config = new(Config)
		_config.parse()
	}
	return *_config
}

//Methode liefert den HTTP-Port unter dem der Blog erreichbar ist, zurück.
func (config *Config) GetHttpsPort() int {
	return config.https_port
}

//Methode liefert die Sessionlebensdauer in Minuten,
// nach der die Autoren-Sessions am Backend des Blogs ablaufen, zurück.
func (config *Config) GetSessionExpire() int {
	return config.session_expire
}

//Methode liefert den Pfad zum Resource-Folder zurück.
func (config *Config) GetResourceFolder() string {
	return config.resource_folder
}

//Methode intialisiert das Config-Objekt mit den durch Flags angegebenen Werten beim Aufruf
// oder festgelegten default-Werten (HTTP-Port: 8443; Sessionlebendauer: 15min, Ressource-Folder: /res)
func (c *Config) parse() {
	flag.IntVar(&c.https_port, "https-port", 8443, "HTTPs Port of the Go-Blog")
	flag.IntVar(&c.session_expire, "session-expire", 15, "Session expire duration in seconds")
	flag.StringVar(&c.resource_folder, "resource-folder", "res/", "Path to the resource folder")
	flag.Parse()

	// Existenz des Resource-Folders überprüfen
	_, err := os.Stat(c.resource_folder)
	if err != nil {
		panic("Ressource folder could not be found. Please check the resource-folder parameter");
	}
}