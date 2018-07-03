// Enthält das Struct Post und alle zugehörigen Methoden für einen Blog-Beitrag.
package post

// Go-Blog von 8892993, 1734394, 1777093

import "time"

//Der Struct Post defniert das Model des Posts /Blogbeitrags mit dessen Attributen.
// Diese sind eine ID, der Titel, der Untertitel, der Inhalt, das Erstellungsdatum und der Autor(ID) des Posts
// sowie eine Keyword-Liste.
type Post struct {
	Id int `json: "id"`
	Title string `json: "title"`
	SubTitle string `json: "sub_title"`
	Content string `json: "content"`
	Date time.Time `json: "date"`
	Author int `json: "author_id"`
	Keywords []int `json: "keywords"`
}

//Konstruktor für ein neues Post-Objekt eines Posts/ Blogbeitrags
// mit dem übergebenen Titel und Untertitel, dem übergebenen Post-Inhalt
// und der übergebenen ID des Autors sowie dem momentanten Zeitpunkt als Erstellungsdatum.
func NewPost(title string, subtitle string, content string, author_id int) *Post {
	p := new(Post)
	p.Title = title
	p.SubTitle = subtitle
	p.Content = content
	p.Date = time.Now()
	p.Author = author_id
	return p
}

//Methode ordent den Post einem weiteren Keyword (/Kategorie) zu.
func (p *Post) AddKeyword(keyword int) {
	p.Keywords = append(p.Keywords, keyword)
}

//Methode liefert ID des Posts zurück.
func (p * Post) GetID() int {
	return p.Id
}

//Methode liefert Inhalt des Posts zurück.
func (p * Post) GetContent() string {
	return p.Content
}

//Methode liefert Erstellungsdatum des Posts zurück.
func (p * Post) GetDate() time.Time {
	return p.Date
}

//Methode liefert ID des Autors des Posts zurück.
func (p * Post) GetAuthorID() int {
	return p.Author
}

//Methode liefert eine Liste der Keywords des Posts zurück.
func (p * Post) GetKeywords() []int {
	return p.Keywords
}

//Methode setzt den übergebenen Integer-Wert als ID des Posts.
func (p *Post) SetID(id int) {
	p.Id = id
}

//Methode setzt den übergebenen String als Titel des Posts.
func (p *Post) SetTitle(title string) {
	p.Title = title
}

//Methode setzt den übergebenen String als Untertitel des Posts.
func (p *Post) SetSubTitle(subtitle string) {
	p.SubTitle = subtitle
}

//Methode setzt den übergebenen String als Inhalt des Posts.
func (p *Post) SetContent(content string) {
	p.Content = content
}

//Methode setzt die übergebene Liste aus Integer-Werten (IDs der Keywords) als Keywords (/Kategorien) des Posts.
func (p *Post) SetKeywords(keywords []int) {
	p.Keywords = keywords
}

//Methode setzt den übergebenen IntegerWert(Autor-ID) als Autor des Posts.
func (p *Post) SetAuthor(author int) {
	p.Author = author
}

//Methode setzt das übergebene Datum als Erstellungsdatum des Posts.
func (p *Post) SetDate(date time.Time) {
	p.Date = date
}