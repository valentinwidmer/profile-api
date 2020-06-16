package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/supercoast/profile-api/data"
	"github.com/supercoast/profile-api/handlers"
)

var (
	serverListener = os.Getenv("LISTEN_ADDRESS")
	serverPort     = os.Getenv("PORT")
)

func main() {

	l := hclog.New(&hclog.LoggerOptions{
		Name:  "profile-api",
		Level: hclog.LevelFromString("DEBUG"),
	})

	profileDB := data.NewProfileDB(l)
	profile := handlers.NewProfile(l, profileDB)

	mux := mux.NewRouter()
	getRouter := mux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/v1/profiles", profile.ListProfiles)
	getRouter.HandleFunc(`/api/v1/profiles/{email:/^\S+@\S+\.\S+$/}`, profile.GetProfile)

	postRouter := mux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/v1/profiles", profile.CreateProfile)

	addr := strings.Join([]string{serverListener, serverPort}, ":")

	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	l.Info("Starting Server", "Address", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting server: %v", err)
	}
}
