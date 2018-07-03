// Enthält das Struct Author und alle zugehörigen Methoden für einen Autor.
package author

// Go-Blog von 8892993, 1734394, 1777093

import (
	"de/blog/projekt/01/core/util"
	"encoding/hex"
	"crypto/sha512"
)

//Der Struct Author definiert das Model des Autors mit dessen Attributen .
// Diese sind eine ID, sein Username, ein generierter Salt-Wert
// und der Passwort-Hash des damit verschlüsselten Passwortes.
type Author struct {
	Id int `json: "id"`
	Username string `json: "username""`
	Pwd_hash string `json: "password_hash"`
	Salt string `json: "password_salt"`
}

//Konstruktor für ein neues Author-Objekt eines Autors
// mit dem übergebenen Username, einem zufälligen Salt-Wert
// und dem damit berechneten Passwort-Hash des übergebenen Passworts.
func NewAuthor(username, password string) *Author {
	a := new(Author)
	a.Username = username
	a.Salt = a.generateSalt()
	a.Pwd_hash = a.hashPassword(password)
	return a
}

//Methode zum Verifizieren des übergebenen Passwortes durch Abgleich mit dem gespeicherten Passwort-Hash
// Liefert zurück ob das Passwort korrekt ist oder nicht.
func (a * Author) Verify(pwd string) bool {
	return a.hashPassword(pwd) == a.Pwd_hash
}

//Methode zum Generieren eine zufälligen Salt-Wertes der Länge 20
func (a * Author) generateSalt() string {
	return util.RandomString(20)
}

//Methode zur Verschlüsselung des übergebenen Passwortes mit dem Hashing-Algorithmus SHA512 und dem zufälligen Salt-Wert
// Liefert den Hash des übergebenen Passwortes zurück.
func (a * Author) hashPassword(password string) string {
	sha := sha512.New()
	sha.Write([]byte(password + a.Salt))
	hash := hex.EncodeToString(sha.Sum(nil))
	return hash
}

//Methode liefert ID des Autors zurück.
func (a * Author) GetID() int {
	return a.Id
}

//Methode liefert Username des Autors zurück.
func (a * Author) GetUsername() string {
	return a.Username
}

//Methode setzt den übergebenen String als Username des Autors.
func (a * Author) SetUsername(username string) {
	a.Username = username
}

//Methode generiert einen Salt-Wert und setzt den Hash des übergebenen String als Passwort des Autors.
func (a * Author) SetPassword(newPassword string) {
	a.Salt = a.generateSalt()
	a.Pwd_hash = a.hashPassword(newPassword)
}

//Methode setzt den übergebenen Integer-Wert als ID des Autors.
func (a * Author) SetID(newId int) {
	a.Id = newId
}