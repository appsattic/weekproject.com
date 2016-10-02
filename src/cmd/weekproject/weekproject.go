package main

import (
	"io"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

func projectHandler() http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, "<h2>The Week Project</h2>\n")
		io.WriteString(w, "<p>This project is to create the week project.</p>\n")
		io.WriteString(w, "<p>Started: 11:10 AM - 2 Oct 2016</p>\n")
		io.WriteString(w, "<p>(Ends)</p>\n")
	}

	return http.HandlerFunc(handler)
}

func userHandler() http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, "<h2>User : chilts</h2>\n")
		io.WriteString(w, "<p>My projects:</p>\n")
		io.WriteString(w, "<ul><li><a href='week-project'>Week Project</a></li></ul>\n")
		io.WriteString(w, "<p>(Ends)</p>\n")
	}

	return http.HandlerFunc(handler)
}

func handler() http.Handler {
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
	log.Println("Started WeekProject Server")

	mux := pat.New()

	mux.Get("/chilts/week-project", projectHandler())
	mux.Get("/chilts/", userHandler())
	mux.Get("/", handler())

	http.Handle("/", mux)
	err := http.ListenAndServe(":8504", nil)
	log.Fatal(err)
}
