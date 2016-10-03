package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bmizerany/pat"
	"github.com/chilts/temple"

	"internal"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func projectHandler(tmpl *temple.Temple, store internal.StoreService) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// get the template
		projectTmpl, err := tmpl.Get("project.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			return
		}

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

		// create some data
		data := struct {
			Title   string
			User    *internal.User
			Project *internal.Project
		}{
			"The Week Project",
			user,
			project,
		}

		// execute the template
		err = projectTmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(handler)
}

func userHandler(tmpl *temple.Temple, store internal.StoreService) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// get the template
		userTmpl, err := tmpl.Get("user.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			return
		}

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

		// create some data
		data := struct {
			Title string
			User  *internal.User
		}{
			"The Week Project",
			user,
		}

		// execute the template
		err = userTmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(handler)
}

func handler(tmpl *temple.Temple, store internal.StoreService) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// get the template
		index, err := tmpl.Get("index.html")
		if err != nil {
			log.Println(err)
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			return
		}

		// create some data
		data := struct {
			Title string
		}{
			"The Week Project",
		}

		// execute the template
		err = index.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(handler)
}

func main() {
	var err error

	log.Println("Started WeekProject Server")

	// read the templates
	tmpl, err := temple.NewTemple("templates", "base.html", false)
	check(err)

	// connect to the store
	mongoDbUrl := os.Getenv("MONGODB_URL")
	store, err := internal.NewMongoDbStore(mongoDbUrl)
	check(err)
	defer store.Close()

	// the router
	mux := pat.New()

	mux.Get("/:userName/:projectName", projectHandler(tmpl, store))
	mux.Get("/:userName/", userHandler(tmpl, store))
	mux.Get("/", handler(tmpl, store))

	http.Handle("/", mux)
	err = http.ListenAndServe(":8504", nil)
	log.Fatal(err)
}
