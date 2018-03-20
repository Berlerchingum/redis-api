package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

// Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes
type Routes []Route

// NewRouter
// Get a router and link the routes
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
				HandlerFunc(route.HandlerFunc).
					Name(route.Name)
	}

	return router
}

var routes = Routes{
	Route{
		"Index	",
		"GET",
		"/",
		index,
	},
	Route{
		"insert",
		"POST",
		"/notes",
		createNote,
	},
	Route{
		"getAll",
		"GET",
		"/notes",
		getAllNotes,
	},
	Route{
		"GetById",
		"GET",
		"/notes/{idNote}",
		getNoteById,
	},
	Route{
		"Delete",
		"DELETE",
		"/notes/{idNote}",
		deleteNote,
	},
}