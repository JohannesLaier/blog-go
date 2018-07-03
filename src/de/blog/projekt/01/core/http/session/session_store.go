// Enthält alle Structs und Methoden zu Sessions (Session, SessionStore, ...)
package session

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/core/util"
	"de/blog/projekt/01/core/config"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	"os"
)

// Eine SessionStore-Instanz (nach dem Singleton-Pattern)
var session_store_instances *SessionStore
var session_store_instance_mutex sync.Mutex

//Methode gibt die Instanz des SessionStores zurück. (Singleton-Pattern)
func GetSessionStore() *SessionStore {
	session_store_instance_mutex.Lock()
	defer session_store_instance_mutex.Unlock()

	if session_store_instances == nil {
		session_store_instances = NewSessionStore()
	}

	return session_store_instances
}

//Das Strukt SessionStore definiert einen SessionStore
// mit einer Map der Sessions
type SessionStore struct {
	Sessions map[string]*Session `json: "sessions"`
}

//Konstruktor für ein neues SessionStore-Objekt
// mit einer Map für Sessions.
func NewSessionStore() *SessionStore {
	store := new(SessionStore)
	store.Sessions = make(map[string]*Session)
	return store
}

//Methode erzeugt eine neue Session und speichert diese im SessionStorage.
// Die Methode liefert die ID der neuen Session und das Session-Objekt selbst zurück.
func (store * SessionStore) New() (string, *Session) {
	// Jede Session ID nur einmal vergeben
	session_id := util.RandomString(32)
	for store.Sessions[session_id] != nil {
		session_id = util.RandomString(32)
	}

	session := NewSession(session_id)
	store.Sessions[session_id] = session
	return session_id, session
}

//Methode liefert die Session mit der übergebenen ID zurück, wenn die Session noch nicht abgelaufen ist.
func (store * SessionStore) Get(id string) *Session {
	session := store.Sessions[id]

	if session != nil {
		cfg := config.GetConfig()

		expireDate := session.Date.Add(time.Duration(cfg.GetSessionExpire()) * time.Minute)

		if (time.Now().Unix() > expireDate.Unix()) {
			return nil
		}
	}

	return session
}

//Methode liefert Session zur SessionID aus dem Cookie des übergebenen Http-Requests zurück.
func (store * SessionStore) GetCurrent(r *http.Request) *Session {
	cookie, err := r.Cookie("SESSIONID")
	if err == nil {
		return store.Get(cookie.Value)
	}
	return nil
}

//Methode löscht die Session mit der angebenen ID aus dem SessionStore.
func (store * SessionStore) Discard(id string) {
	delete(store.Sessions, id)
}

//Methode liest die Session-Daten aus der gespeicherten Datei im JSON-Format ein, wandelt diese in Objekte um
// und lädt sie in den SessionStore.
func (store * SessionStore) Load() {
	raw, err := ioutil.ReadFile(store.getPath())
	if err == nil {
		json.Unmarshal(raw, &store.Sessions)
	}
}

//Methode wandelt die Session-Daten in das JSON-Format um und speichert sie in einer Datei.
func (store * SessionStore) Store() {
	json_string, _ := json.Marshal(store.Sessions)
	ioutil.WriteFile(store.getPath(), json_string, 0600)
}

//Methode liefert den Pfad der SessionStore-Datei auf der Festplatte zurück.
func (store * SessionStore) getPath() string {
	conf := config.GetConfig()

	session_folder := conf.GetResourceFolder() + "sessions/"
	session_path := session_folder + "sessions.json"

	// Create folder if not exist
	os.Mkdir(session_folder, os.ModePerm)

	return session_path
}






