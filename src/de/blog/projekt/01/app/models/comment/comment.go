// Enthält das Struct Comment und alle zugehörigen Methoden für ein Kommentar.
package comment

// Go-Blog von 8892993, 1734394, 1777093

import "time"

//Der Struct Comment definiert das Model des Kommentars mit dessen Attributen.
// Diese sind die ID des Posts zu dem das Kommentar gehört, der Nickname des Kommentators, der Kommentartext,
// das Erstellungsdatum und ein Freigabe-Status.
type Comment struct {
	Id int `json: "id"`
	Post int `json: "post_id"`
	Username string `json: "username"`
	Text string `json: "text"`
	Share bool `json: "share"`
	Date time.Time `json: "date"`
}

//Konstruktor für eine neues Comment-Objekt eines Kommentars
// mit dem übergebenen Nickname des Kommentators, dem übergebenen Kommentartext, der übergebenen ID des zugehörigen Post
// sowie dem momentanten Zeitpunkt als Erstellungsdatum.
// Das Attribut Shared ist bei Erstellung immer false, da das Kommentar erst durch den Autor freigegeben werden muss.
func NewComment(username string, text string, post int) * Comment{
	c := new(Comment)
	c.Post = post
	c.Username = username
	c.Text = text
	c.Date = time.Now()
	c.Share = false
	return c
}

//Methode liefert ID des Kommentars zurück.
func (c * Comment) GetID() int {
	return c.Id
}

//Methode liefert Username des Kommentators zurück.
func (c * Comment) GetUsername() string {
	return c.Username
}

//Methode liefert Text des Kommentars zurück.
func (c * Comment) GetText() string {
	return c.Text
}

//Methode liefert ID des Post zu dem das Kommentars gehört zurück.
func (c * Comment) GetPost() int {
	return c.Post
}

//Methode liefert den Freigabe-Status des Kommentars zurück.
func (c * Comment) GetShare() bool{
	return c.Share
}

//Methode liefert das Erstellungsdatum des Kommentars zurück.
func (c * Comment) GetDate() time.Time {
	return c.Date
}

//Methode setzt den Freigabe-Status des Kommentars.
func (c * Comment) SetShare(shareStatus bool) {
	c.Share = shareStatus
}

//Methode setzt den übergebenen Integer-Wert als ID des Kommentars.
func (c * Comment) SetID(id int) {
	c.Id = id
}