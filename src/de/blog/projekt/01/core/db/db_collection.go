package db

// Go-Blog von 8892993, 1734394, 1777093

//Typ-Definition eines Datenbank-Collection-Eintrags als Interface mit Getter und Setter für die ID des Eintrags
type DBCollectionEntry interface {
	GetID() int
	SetID(id int)
}

//Typ-Definition einer Datenbank-Collection-Filter-Funktion
// die für einen Collection-Eintrag bestimmt, ob die Funktion darauf zutrifft.
type DBCollectionFilterFunc func (entry DBCollectionEntry) bool

//Das Struct DBCollection definiert eine Datenbank-Collection (vergleichbar einer Tabelle)
// mit einem Namen, einen Zähler für die Primärschlüssel und einem Set von Collection-Einträgen.
type DBCollection struct {
	Name string `json: "name"`
	primarykey int `json: "primary_key"`
	Set []DBCollectionEntry `json: "set"`
}

//Konstruktor für ein neues Datenbank-Collection-Objekt
// mit dem übergebenen Namen, dem Primärschlüssel von Null und einem Set für Collection-Einträge.
func NewDBCollection(name string) * DBCollection {
	collection := new(DBCollection)
	collection.Name = name
	collection.primarykey = 0
	collection.Set = make([]DBCollectionEntry, 0)
	return collection
}

//Methode liefert den Namen der Datenbank-Collection zurück.
func (collection *DBCollection) GetName() string {
	return collection.Name
}

//Methode liefert das Set mit allen Collection-Einträgen der Datenbank-Collection zurück.
func (collection * DBCollection) GetAll() []DBCollectionEntry {
	return collection.Set
}

//Methode liefert den Collection-Eintrag mit der übergebenen ID aus der Datenbank-Collection zurück.
// Wenn kein Eintrag mit der übergebenen ID existiert, liefert sie nil zurück.
func (collection * DBCollection) GetByID(id int) DBCollectionEntry {
	for _, entry := range collection.Set {
		if entry.GetID() == id {
			return entry
		}
	}
	return nil
}

//Methode liefert ein Set mit den Collection-Einträge aus der Datenbank-Collection zurück,
// die der übergebenen Funkion entsprechen.
func (collection * DBCollection) GetListFilter(filterFunc DBCollectionFilterFunc) []DBCollectionEntry {
	var result []DBCollectionEntry
	for _, entry := range collection.Set {
		if filterFunc(entry) {
			result = append(result, entry)
		}
	}
	return result
}

//Methode liefert den ersten Collection-Eintrag der Datenbank-Collection zurück,
// der der übergebenen Funktion entspricht.
func (collection * DBCollection) GetFilter(filterFunc DBCollectionFilterFunc) DBCollectionEntry {
	for _, entry := range collection.Set {
		if filterFunc(entry) {
			return entry
		}
	}
	return nil
}

//Methode fügt den übergebenen Collection-Eintrag in die Collection hinzu
// und liefert die ID des Eintrages in der Collection zurück.
func (collection * DBCollection) Add(entry DBCollectionEntry) int {
	id := collection.next()
	entry.SetID(id)
	collection.Set = append(collection.Set, entry)
	return id
}

//Methode überschreibt den Collection-Eintrag mit den Daten des übergebenen Collection-Eintrages.
func (collection * DBCollection) Update(e DBCollectionEntry) {
	for index, entry := range collection.Set {
		if e.GetID() == entry.GetID() {
			collection.Set[index] = e
		}
	}
}

//Methode entfernt den Collection-Eintrag mit der übergebenen ID aus der Datenbank-Collection.
func (collection * DBCollection) Remove(id int) {
	for index, entry := range collection.Set {
		if entry.GetID() == id {
			collection.Set = append(collection.Set[:index], collection.Set[index+1:]...)
			break
		}
	}
}

//Methode erhöht den Primärschlüsselzähler um eins und liefert diesen zurück.
func (collection * DBCollection) next() int {
	collection.primarykey++
	return collection.primarykey
}