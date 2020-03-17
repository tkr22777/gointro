package gone

/* Practicing https://golang.org/doc/articles/wiki/ */

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

/* The method reads:
	A method named save() that takes as its receiver p, a pointer to Page
	It takes no parameter, and returns a value of type error
	The save method returns an error value because that is the type of WriterFile
	If all goes well, Page.save() will return nil (the zero value for pointers, interfaces,
	and some other types
 */
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &Page{title, body}
}


