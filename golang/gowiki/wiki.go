package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
)

// Wiki - series of interconnected pages, each of which has a title
// and body (the page content)
type Page struct {
	Title string
	Body  []byte
}

// Template Caching
// Load ParseFiles once at program initialization and parse all templates into a single
// template
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// Security Flaw: a user can supply an arbitrary path to be read/written on the server
// Validate titles with a regular expression
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// Validate path and extract the page title
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

// Persistent storage
// Save the page body to a text file
// octal integer literal 0600 indicates the file should be created wtih read-write
// permissions for the current user only (open(2) linux man pages)
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// Construct and load a new page
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// HTML Template helper function
// template function with error checking
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Networking code
// Wrapper function that returns a function of type http.HandlerFunc
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// viewHandler - view a wiki page
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// editHandler - load a wiki page
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// saveHandler - save a wiki page
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	// Test a basic page, saving a page and loading a page
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))

	// network handlers
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
