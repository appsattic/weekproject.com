package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bmizerany/pat"

	"internal"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func projectHandler(store internal.StoreService) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// see if we can get this user
		userName := r.URL.Query().Get(":userName")

		user, err := store.GetUser(userName)
		if err != nil {
			log.Println(err)
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			return
		}
		if user == nil {
			http.NotFound(w, r)
			return
		}

		// see if we can get this project
		projectName := r.URL.Query().Get(":projectName")

		project, err := store.GetProject(projectName)
		log.Printf("project=%#v\n", project)
		if err != nil {
			log.Println(err)
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			return
		}
		if project == nil {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, "<h2>"+project.Title+"</h2>\n")
		io.WriteString(w, "<p>This project is to create the week project.</p>\n")
		io.WriteString(w, "<p>Started: 11:10 AM - 2 Oct 2016</p>\n")
		io.WriteString(w, "<p>(Ends)</p>\n")
	}

	return http.HandlerFunc(handler)
}

func userHandler(store internal.StoreService) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// see if we can get this user
		userName := r.URL.Query().Get(":userName")

		user, err := store.GetUser(userName)
		if err != nil {
			log.Println(err)
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			return
		}
		if user == nil {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, "<h2>User : "+user.Name+"</h2>\n")
		io.WriteString(w, "<p>My projects:</p>\n")
		io.WriteString(w, "<ul><li><a href='week-project'>Week Project</a></li></ul>\n")
		io.WriteString(w, "<p>(Ends)</p>\n")
	}

	return http.HandlerFunc(handler)
}

func handler(store internal.StoreService) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, "<p>Hello, The Week Project!<p>\n")
		io.WriteString(w, "<p>See the first ever <a href='/chilts/week-project'>Week Project</a>.<p>\n")
		io.WriteString(w, "<p>From the first ever user <a href='/chilts/'>chilts</a>.<p>\n")
		io.WriteString(w, "<p>(Ends)</p>\n")
	}

	return http.HandlerFunc(handler)
}

func main() {
	var err error

	log.Println("Started WeekProject Server")

	// connect to the store
	mongoDbUrl := os.Getenv("MONGODB_URL")
	store, err := internal.NewMongoDbStore(mongoDbUrl)
	check(err)
	defer store.Close()

	// the router
	mux := pat.New()

	mux.Get("/:userName/:projectName", projectHandler(store))
	mux.Get("/:userName/", userHandler(store))
	mux.Get("/", handler(store))

	http.Handle("/", mux)
	err = http.ListenAndServe(":8504", nil)
	log.Fatal(err)
}
