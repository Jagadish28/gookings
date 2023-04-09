package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Jagadish28/bookings/pkg/config"
	"github.com/Jagadish28/bookings/pkg/handler"
	"github.com/Jagadish28/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":4444"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this true in PROD

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	app.TemplateCache = tc
	app.UseCache = false
	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplate(&app)

	if err != nil {
		log.Fatal("unable to create template cache")
	}

	// http.HandleFunc("/", handler.Repo.Home)
	// http.HandleFunc("/about", handler.Repo.About)

	fmt.Println(fmt.Sprintf("starting application at %s", port))

	// _ = http.ListenAndServe(port, nil)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()

	log.Fatal(err)
}

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
// 	n, err := fmt.Fprint(w,"Hello World")

// 	if(err != nil){
// 		fmt.Println(err)
// 	}

// 	fmt.Println(fmt.Sprintf("Bytes written: %d",n))
// })

// func addValues(x, y int) int {
// 	return x + y
// }

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 0.0)

// 	if err != nil {
// 		fmt.Fprintf(w, "cont divide by 0")
// 		return
// 	}
// 	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f ins %f", 100.0, 0.0, f))
// }

// func divideValues(x, y float32) (float32, error) {

// 	if y <= 0 {
// 		err := errors.New("can't divide by 0")
// 		return 0, err
// 	}
// 	result := x / y

// 	return result, nil

// }
