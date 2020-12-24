package queries

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func writeError(httpStatus int, w http.ResponseWriter, context string, err error) {
	var e RestError
	e.Msg = err.Error()
	e.Context = context
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatus)

	json.NewEncoder(w).Encode(e)
}

func sendHTTPResponse(w http.ResponseWriter, v interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(v)
}

//HandleBook return JSON formated Book
func HandleBook(w http.ResponseWriter, r *http.Request) {

	var resp string
	var bookid uint64
	var err error

	switch r.Method {
	case http.MethodGet:
		//resp = fmt.Sprintf("GET Book id: %v\n", vr["id"])
		bookid, err = strconv.ParseUint(mux.Vars(r)["id"], 10, 32)

		if err != nil {
			writeError(http.StatusBadRequest, w, "/books/{id} request handler: parsing ID", err)
			return
		}
		book, err := SQLQueryBook(bookid)
		if err != nil {
			writeError(http.StatusBadRequest, w, "/books/{id} request handler: fetch book in database", err)
			return
		}
		sendHTTPResponse(w, book)
	case http.MethodPost:
		resp = fmt.Sprintf("POST Book id: %v\n", mux.Vars(r)["id"])
	case http.MethodDelete:
		err = SQLQueryDeleteBook(bookid)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Method %v not handled", r.Method)
		return
	}
	fmt.Fprint(w, resp)
}

//HandleBookLink queies information about a book link and storage
func HandleBookLink(w http.ResponseWriter, r *http.Request) {

	var bookid uint64
	var err error
	var booklink *BookLink

	switch r.Method {
	case http.MethodGet:
		bookid, err = strconv.ParseUint(mux.Vars(r)["id"], 10, 32)

		if err != nil {
			writeError(http.StatusBadRequest, w, "/books/{id}/link request handler: parsing book ID", err)
			return
		}
		booklink, err = SQLQueryBookLink(bookid)
		if err != nil {
			writeError(http.StatusBadRequest, w, fmt.Sprintf("Book link request handler: fetch database book id=%v", bookid), err)
			return
		}
		sendHTTPResponse(w, booklink)

	case http.MethodPost:
		w.WriteHeader(http.StatusNotImplemented)
	case http.MethodDelete:
		w.WriteHeader(http.StatusNotImplemented)
	default:
		writeError(http.StatusMethodNotAllowed, w, fmt.Sprintf("/books/{id}/link request handler: Book ID=%v", bookid),
			fmt.Errorf("Method %v not allowed", r.Method))
		return
	}
}

//HandleBookSubjects return list of subjects associated with a book ID
func HandleBookSubjects(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		bookid, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)

		if err != nil {
			writeError(http.StatusBadRequest, w, fmt.Sprintf("/books/{id}/link request handler: parsing book ID=%v", bookid), err)
			return
		}

		bookSubjects, err := SQLQueryBookSubjects(bookid)

		if err != nil {
			writeError(http.StatusBadRequest, w, fmt.Sprintf("/books/{id}/subjects: fetch database book ID=%v", bookid), err)
			return
		}

		sendHTTPResponse(w, bookSubjects)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

//GetPage return Page in JSON format: images titles, links etc
func GetPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		params := mux.Vars(r)
		var pageNumber, nbItems uint64
		var err error

		pageNumber, err = strconv.ParseUint(params["num"], 10, 32)
		if err != nil {
			writeError(http.StatusBadRequest, w, fmt.Sprintf("GetPage handler: parse page number=%v", params["num"]), err)
			return
		}
		nbItems, err = strconv.ParseUint(params["items"], 10, 32)
		if err != nil {
			writeError(http.StatusBadRequest, w, fmt.Sprintf("GetPage handler: parse number items per page=%v", params["num"]), err)
			panic(err)
		}
		items, err := SQLQueryPage(pageNumber, nbItems)
		if err != nil {
			writeError(http.StatusBadRequest, w, fmt.Sprintf("Page handler: request page %v (%v items)", pageNumber, nbItems), err)
			return
		}
		sendHTTPResponse(w, items)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

//HandleSubjects return list of subjects
func HandleSubjects(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		subjects, err := SQLQuerySubjects()
		if err != nil {
			writeError(http.StatusBadRequest, w, "/subjects request handler", err)
		}
		sendHTTPResponse(w, subjects)
		return
	}
	writeError(http.StatusMethodNotAllowed, w, "/subjects request", fmt.Errorf("Method %v not allowed", r.Method))
}
