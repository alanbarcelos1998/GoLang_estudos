package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// global variables
var templates *template.Template
var client *redis.Client
var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/test", testeGetHandler).Methods("GET")
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// request index page handle
func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := client.LRange("comments", 0, 10).Result()

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
	client.LPush("comments", comment)
	// redirect tp / when the submit form
	http.Redirect(w, r, "/", 302)
}

// request contact page handle
func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.html", nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	session, _ := store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
}

func testeGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	untyped, ok := session.Values["username"]
	if !ok {
		return
	}

	username, ok := untyped.(string)
	if !ok {
		return
	}

	w.Write([]byte(username))
}
