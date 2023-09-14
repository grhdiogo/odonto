package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"odonto/internal/infra/config"
	"odonto/internal/infra/data/pgclient"
	ws "odonto/internal/interf"
	"os"
	"time"

	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const contractVersion = "v1"

func errControl() {
	if r := recover(); r != nil {
		e := r.(error)
		log.Printf("Application shutdown caused by error: '%s'\n", e.Error())
	}
}

func main() {
	//capture panic error
	defer errControl()
	var (
		version  string
		yamlFile string
	)
	// flags
	flag.StringVar(&yamlFile, "c", "configs/context/app-dev.yaml", "Load application settings from path file name")
	flag.StringVar(&version, "v", "0.0.0", "Set application version")
	flag.Parse()
	//load configurations (panic if yamlFile not exists)
	log.Printf("Load configurations in %s", yamlFile)
	config.Init(yamlFile)
	// port
	port := os.Getenv("PORT")
	if port == "" {
		port = "18080"
	}
	//port for app wait
	log.Printf("Listen in %s", port)
	//log for initialize aplication
	log.Printf("Initialized application context: %v", version)
	//init web service
	ctx := context.Background()
	wsvc := ws.NewWebService(ctx, contractVersion)
	wsvc.Init()
	// configure pg database manager
	settings := config.GetSettings()
	dbconfig, err := settings.GetDatabases("postgresql")
	if err != nil {
		panic(err)
	}
	pgconfig := pgclient.Config{
		Host: dbconfig.Host,
		Port: dbconfig.Port,
		DBNm: dbconfig.DBNm,
		User: dbconfig.User,
		Pswd: dbconfig.Pswd,
	}
	pgclient.SetConfiguration(pgconfig)
	//logger for
	loggedRouter := handlers.LoggingHandler(os.Stdout, wsvc.GetRouters())
	//server setup
	srv := &http.Server{
		Handler: loggedRouter,
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		//Good practice: enforce timeouts for servers you create!
		WriteTimeout: 800 * time.Second,
		ReadTimeout:  800 * time.Second,
	}
	//log initializing webapp
	log.Printf("Initializing webapp")
	//log
	log.Fatal(srv.ListenAndServe())
}
