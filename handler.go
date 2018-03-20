package main

import (
	"net/http"
	"html"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"io"
	"github.com/gorilla/mux"
)

// Index
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// Create
func createNote(w http.ResponseWriter, r *http.Request) {
	var note Note

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	json.Unmarshal(body, &note)
	handleFatalError(err, w)
	createdNote := insert(note)
	err = json.NewEncoder(w).Encode(createdNote)
	handleFatalError(err, w)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// Get
func getAllNotes(w http.ResponseWriter, r *http.Request) {
	notes := getAll()
	err := json.NewEncoder(w).Encode(notes)
	handleFatalError(err, w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func getNoteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	note, err := get(vars["idNote"])
	handleNotFound(err, w)
	err = json.NewEncoder(w).Encode(note)
	handleFatalError(err, w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deleteById(vars["idNote"])
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func handleFatalError(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleNotFound(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}