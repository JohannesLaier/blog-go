// Enthält das Struct Keyword und alle zugehörigen Methoden für eine Kategorie.
package keyword

// Go-Blog von 8892993, 1734394, 1777093

//Der Struct Keyword definiert das Model des Schlüsselwortes / der Kategorie mit dessen Attributen .
// Diese sind eine ID und der Name der Kategorie.
type Keyword struct {
	Id int `json: "id"`
	Name string `json: "name"`
}

//Konstruktor für ein neues Keyword-Objekt eines Schlüsselwortes
// mit dem übergebenen Kategorie-Name
func NewKeyword(name string) *Keyword {
	k := new(Keyword)
	k.Name = name
	return k
}

//Methode liefert ID des Schlüsselwortes / der Kategorie zurück.
func (k * Keyword) GetID() int {
	return k.Id
}

//Methode liefert den Namen des Schlüsselwortes / der Kategorie zurück.
func (k * Keyword) GetName() string {
	return k.Name
}

//Methode setzt den übergebenen Integer-Wert als ID des Schlüsselwortes / der Kategorie.
func (k * Keyword) SetID(id int) {
	k.Id = id
}