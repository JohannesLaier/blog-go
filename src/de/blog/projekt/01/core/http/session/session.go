package session

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/core/config"
	"net/http"
	"time"
)

//Typ-Definition eines SessionValues als Interface
type SessionValue interface{}

//Der Struct Session definiert die Session
// mit einer ID, dem Erstellunsdatum und einer Map für die Session-Daten
type Session struct {
	Id string `json: "session_id"`
	Date time.Time `json: "session_expire"`
	Data map[string]SessionValue `json: "session_data"`
}

//Konstruktor für ein neues Session-Objekt einer Session
// mit der übergebenen ID, dem momentanten Zeitpunkt als Erstellungsdatum
// und eine Map für die Daten der Session
func NewSession(id string) *Session {
	session := new(Session)
	session.Id = id
	session.Date = time.Now()
	session.Data = make(map[string]SessionValue)
	return session
}

//Methode erstellt einen Cookie mit der SessionID und der "Lebensdauer"(Session-Erstellung + 15min) und setzt diesen.
func (session * Session) CreateCookie(w http.ResponseWriter) {
	cfg := config.GetConfig()
	expiration := time.Now().Add(time.Duration(cfg.GetSessionExpire()) * time.Minute)
	cookie := http.Cookie{Name: "SESSIONID", Value: session.Id, Expires: expiration}
	http.SetCookie(w, &cookie)
}

//Methode setzt die Lebensdauer des Cookies auf 0 sec um ihn zu löschen.
func (session * Session) DestroyCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("SESSIONID")
	cookie.Expires = time.Unix(0, 0)
	http.SetCookie(w, cookie)
}

//Methode liefert die Daten zum übergebenen Schlüssel aus den Session-Daten zurück.
func (session * Session) Get(key string) SessionValue {
	cfg := config.GetConfig()
	expireDate := session.Date.Add(time.Duration(cfg.GetSessionExpire()) * time.Minute)

	if (time.Now().Unix() > expireDate.Unix()) {
		return nil
	}

	return session.Data[key]
}

//Methode setzt den übergebenen Wert unter dem übergebenen Key in die Session-Daten.
func (session * Session) Put(key string, value SessionValue) {
	session.Data[key] = value
}

//Methode löscht die Daten des übergebenen Schlüssels aus den Session-Daten.
func (session * Session) Remove(key string) {
	delete(session.Data, key)
}

