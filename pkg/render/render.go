package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

// RendersTemplate renders templates
func RendersTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}

	// Replaced below logic with above logic. Keeping for reference.

	// parsedTemplate, err := template.ParseFiles("./../../templates/" + tmpl)
	// if err != nil {
	// 	fmt.Println("error parsing template: ", err)
	// 	return
	// }

	// err = parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("error executing parsed template: ", err)
	// 	return
	// }
}

// CreateTemplateCache creates a template cache as a map.
func CreateTemplateCache() (map[string]*template.Template, error) {

	// 	Create a cache to store our parsed templates.
	myCache := map[string]*template.Template{}

	// 	Find pages in template folder.
	pages, err := filepath.Glob("./templates/*.pages.html")
	if err != nil {
		return myCache, err
	}

	// 	Range through pages found in template folder.
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
