package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

// global variables
var templates *template.Template
var client *redis.Client

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/contact", contactHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// request index page handle
func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := client.LRange(client.Context(), "comments", 0, 10).Result()

	if err != nil {
		return
	}
	templates.ExecuteTemplate(w, "index.html", comments)
}

// request index page POST handle
func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// get the comment in html tag comment
	comment := r.PostForm.Get("comment")
	// push the comment to the comments list
	client.LPush(client.Context(), "comments", comment)
	// redirect tp / when the submit form
	http.Redirect(w, r, "/", 302)
}

// request contact page handle
func contactHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "contact.html", "This is the contact Page")
}

// request about page handle
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "about.html", "This is the about Page")
}
