package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/samarasteph/bookstore/web-server/queries"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.RequestURI)

	// Returns hello world! as a response
	fmt.Fprintln(w, "Hello world!")
}

func downloadBook(w http.ResponseWriter, r *http.Request) {
}
func uploadBook(w http.ResponseWriter, r *http.Request) {
}
func authentifyUser(w http.ResponseWriter, r *http.Request) {
}

func main() {

	port := flag.Int("port", 8000, "server listening port")
	addr := flag.String("addr", "127.0.0.1", "listening address")
	dbaddr := flag.String("dbserver", "dbserver", "db server address")
	imgDir := flag.String("img", "./IMG", "img folder static path")

	flag.Parse()

	queries.DBConnectionStr = fmt.Sprintf("ro_user@%v:ro_aldu@tcp(%v)/booksdb", *dbaddr, *dbaddr)

	// registers handleHello to GET /hello
	r := mux.NewRouter()
	r.HandleFunc("/", handleHome)
	r.HandleFunc("/books/{id:[1-9][0-9]*}", queries.HandleBook)
	r.HandleFunc("/books/{id:[1-9][0-9]*}/link", queries.HandleBookLink)
	r.HandleFunc("/books/{id:[1-9][0-9]*}/subjects", queries.HandleBookSubjects)
	r.HandleFunc("/subjects", queries.HandleSubjects)
	r.HandleFunc("/page/{num:[1-9][0-9]*}/{items:[1-9][0-9]*}", queries.GetPage)
	r.HandleFunc("/auth", authentifyUser)
	r.HandleFunc("/upload", uploadBook)
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir(*imgDir))))

	listAddr := *addr + ":" + strconv.Itoa(*port)
	fmt.Println("Server listening on " + listAddr + "...")
	fmt.Println("Database connection=" + queries.DBConnectionStr)

	srv := &http.Server{
		Handler: r,
		Addr:    listAddr,
	}
	log.Fatal(srv.ListenAndServe())
}
