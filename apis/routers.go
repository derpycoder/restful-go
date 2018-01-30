package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

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
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"CORSPreflight",
		"OPTIONS",
		"/v1/users",
		nil,
	},
	Route{
		"CreateUser",
		"POST",
		"/v1/users",
		CreateUser,
	},

	Route{
		"DeleteUser",
		"DELETE",
		"/v1/users/{id}",
		DeleteUser,
	},
	Route{
		"GetAllUsers",
		"GET",
		"/v1/users",
		GetAllUsers,
	},
	Route{
		"GetUser",
		"GET",
		"/v1/users/{id}",
		GetUser,
	},
	Route{
		"PatchUser",
		"PATCH",
		"/v1/users/{id}",
		PatchUser,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/v1/users/{id}",
		UpdateUser,
	},
	Route{
		"UploadProfileImage",
		"POST",
		"/v1/users/{id}/image",
		UploadProfileImage,
	},
}

type Options struct {
	Documentation string `json:"documentation"`
	Endpoint      string `json:"endpoint"`
}

func appendHeader(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Status Code", strconv.Itoa(statusCode))
}

func commonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		if r.Method == "OPTIONS" {
			options := Options{"To learn how to use this endpoint, please refer", "http://abhijit-kar.com/swagger/"}
			json, err := json.Marshal(options)
			if err != nil {
				w.Write([]byte(err.Error()))
				appendHeader(w, http.StatusOK)
				return
			}

			w.Write(json)
			appendHeader(w, http.StatusOK)
			return
		}

		fn.ServeHTTP(w, r)
	}
}
