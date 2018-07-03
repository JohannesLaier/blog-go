// Enthält alle Structs und Methoden zur Datenbank (DB, DBCollection, DBCollectionEntry, ...)
package db

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/core/config"
	"encoding/json"
	"io/ioutil"
	"sync"
	"os"
)

//Eine Map mit Datenbank-Instanzen
var db_instances = make(map[string]*DB)
var db_instance_mutex sync.Mutex

//Methode liefert die Datenbank mit dem angegebenen Namen zurück.
func Get(name string) *DB {
	db_instance_mutex.Lock()
	defer db_instance_mutex.Unlock()

	if db_instances[name] == nil {
		db_instance := NewDB(name)
		db_instance.Load()
		db_instances[name] = db_instance
	}
	return db_instances[name]
}

//Das Struct DB definiert die Datenbank
// mit einem Namen und einer Map von Datenbank-Collections
type DB struct {
	Name string `json: "db_name"`
	Collections map[string]*DBCollection `json: "db_collections"`
}

//Konstruktor für ein neues Datenbank-Objekt
// mit dem übergebenen Namen und einer Map für Datenbank-Collections
func NewDB(name string) * DB {
	db := new(DB)
	db.Name = name
	db.Collections = make(map[string]*DBCollection)
	return db
}

//Methode fügt übergebene Datenbank-Collection on die Collection-Map der Datenbank hinzu.
func (db * DB) AddCollection(collection * DBCollection) {
	db.Collections[collection.GetName()] = collection
}

//Methodie liefert Datenbank-Collection mit dem übergebenen Namen zurück.
func (db * DB) GetCollection(name string) * DBCollection {
	return db.Collections[name]
}

//Methode liest die Datenbank-Daten aus der gespeicherten Datei im JSON-Format ein, wandelt diese in Objekte um,
// und lädt sie in die Datenbank (als Collections).
func (db * DB) Load() {
	raw, err := ioutil.ReadFile(db.getPath())
	if err == nil {
		json.Unmarshal(raw, &db.Collections)
	}
}

//Methode wandelt die Daten in das JSON-Format und speichert sie in einer Datei.
func (db * DB) Store() {
	json_string, _ := json.Marshal(db.Collections)
	ioutil.WriteFile(db.getPath(), json_string, 0600)
}

//Methode liefert den Pfad der Datenbank-Datei auf der Festplatte zurück.
func (db * DB) getPath() string {
	conf := config.GetConfig()

	database_folder := conf.GetResourceFolder() + "db/"
	database_path := database_folder  + db.Name + ".json"

	// Create folder if not exist
	os.Mkdir(database_folder, os.ModeDir)

	return database_path
}

