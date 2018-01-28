package apis

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		withCORS(Index),
	},

	Route{
		"CreateUser",
		"POST",
		"/v1/users",
		withCORS(CreateUser),
	},

	Route{
		"DeleteUser",
		"DELETE",
		"/v1/users/{uuid}",
		withCORS(DeleteUser),
	},

	Route{
		"GetAllUsers",
		"GET",
		"/v1/users",
		withCORS(GetAllUsers),
	},

	Route{
		"GetUser",
		"GET",
		"/v1/users/{uuid}",
		withCORS(GetUser),
	},

	Route{
		"PatchUser",
		"PATCH",
		"/v1/users/{uuid}",
		withCORS(PatchUser),
	},

	Route{
		"UpdateUser",
		"PUT",
		"/v1/users/{uuid}",
		withCORS(UpdateUser),
	},

	Route{
		"UploadProfileImage",
		"POST",
		"/v1/users/{uuid}/image",
		withCORS(UploadProfileImage),
	},
}

func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fn(w, r)
	}
}
