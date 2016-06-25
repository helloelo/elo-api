package app

import (
	"net/http"

	"elo-api/internal/handler"

	"github.com/gorilla/mux"
)

// Route is a struct containing routes informations
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is the list of routes availables
type Routes []Route

// InitRouter initialize routes and router
func InitRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"Login",
		"GET",
		"/login",
		handler.Login,
	},
	Route{
		"User",
		"GET",
		"/user/{userID}",
		handler.GetUser,
	},
	Route{
		"Notifications",
		"GET",
		"/notifications/{userID}",
		handler.GetNotifications,
	},
	Route{
		"Top Players",
		"GET",
		"/top-players/{gameID}",
		handler.GetTopPlayers,
	},
	Route{
		"Notification",
		"POST",
		"/notification",
		handler.PostNotification,
	},
	Route{
		"Approve",
		"POST",
		"/approve",
		handler.PostApprove,
	},
	Route{
		"Decline",
		"POST",
		"/decline",
		handler.PostDecline,
	},
	Route{
		"Create Match",
		"POST",
		"/match",
		handler.PostMatch,
	},
	Route{
		"Put Match result",
		"PUT",
		"/match/{matchID}",
		handler.PutResult,
	},
}
