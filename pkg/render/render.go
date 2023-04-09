package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Jagadish28/bookings/pkg/config"
	"github.com/Jagadish28/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	// create template cache
	// check app config
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("couldn't get teplate from temp cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println(err)
	}
	// parsedTemplate, _ := template.ParseFiles("../../templates/"+tmpl, "../../templates/base.html")
	// err := parsedTemplate.Execute(w, nil)

	// if err != nil {
	// 	fmt.Println("prseing err", err)
	// 	return
	// }

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*.page.html")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return cache, err
		}
		matches, err := filepath.Glob("../../templates/*.layout.html")

		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.html")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}
	return cache, nil
}

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {

// 	var tmpl *template.Template
// 	var err error

// 	_, inMap := tc[t]

// 	if !inMap {
// 		log.Println("creating template cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		log.Println("using cache")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)

// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(s string) error {
// 	templates := []string{
// 		fmt.Sprintf("../../templates/%s", s),
// 		"../../templates/base.html",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return err
// 	}

// 	tc[s] = tmpl
// 	return nil
// }
