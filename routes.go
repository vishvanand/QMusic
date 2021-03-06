package main

/**************** Declares all routes/API ***************/

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"CreatePartyController",
		"POST",
		"/dashboard/createparty",
		CreatePartyController,
	},

	Route{
		"AuthorizeSpotify",
		"GET",
		"/authspotify",
		AuthorizeSpotify,
	},

	Route{
		"Dashboard",
		"GET",
		"/dashboard",
		Dashboard,
	},

	Route{
		"SearchSong",
		"GET",
		"/searchsong",
		SearchSong,
	},

	Route{
		"ViewPlaylist",
		"GET",
		"/viewplaylist",
		ViewPlaylist,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		/*********** LOGGER CODE *************/
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		/*************************************/

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler) //Analogous to Handler(route.handlerFunc)
	}

	s := http.StripPrefix("/", http.FileServer(http.Dir("./www")))
	router.PathPrefix("/").Handler(s)

	return router
}
