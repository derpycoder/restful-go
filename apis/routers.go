package apis

import (
	"encoding/json"
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
		handler = commonHeaders(route.HandlerFunc)
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method, "OPTIONS").
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"CreateUser",
		"POST",
		"/v0/users",
		CreateUser,
	},

	Route{
		"DeleteUser",
		"DELETE",
		"/v0/users/{id}",
		DeleteUser,
	},
	Route{
		"GetAllUsers",
		"GET",
		"/v0/users",
		GetAllUsers,
	},
	Route{
		"GetUser",
		"GET",
		"/v0/users/{id}",
		GetUser,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/v0/users/{id}",
		UpdateUser,
	},
	Route{
		"UploadProfileImage",
		"POST",
		"/v0/users/{id}/image",
		UploadProfileImage,
	},
}

type Options struct {
	Documentation string `json:"documentation"`
	Endpoint      string `json:"endpoint"`
}

func commonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET, PUT, DELETE")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		if r.Method == "OPTIONS" {
			options := Options{"To learn how to use this endpoint, please refer", "http://abhijit-kar.com/swagger/"}
			json, err := json.Marshal(options)
			if err != nil {
				http.Error(w, options.Documentation+" "+options.Endpoint, http.StatusUnprocessableEntity)
				return
			}

			w.Write(json)
			return
		}

		fn.ServeHTTP(w, r)
	}
}
