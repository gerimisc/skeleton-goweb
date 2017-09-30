package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sugiantoaudi/skeleton-goweb/init"
	"github.com/sugiantoaudi/skeleton-goweb/routing"
	"log"
	"net/http"
	"os"
)

// Multi-line help information
var help = `Usage:
	./main [config file path]/[flags:optional]

The flags are:
	-host	set database IP address
	-port	set database port
	-user	set database user
	-pwd	set database password`

// Declare command-line flags
var host, user, pwd, connectionString string

func init() {
	flag.StringVar(&host, "host", "127.0.0.1:3306", "Database listen spec")
	flag.StringVar(&user, "user", "root", "Database user login")
	flag.StringVar(&pwd, "password", "", "Database user password")
	flag.StringVar(&connectionString, "connection", "", "Connection string info")

	flag.Parse()
}

func main() {
	// Count length of argument
	argsCount := len(os.Args)

	// Create object from db Configuration struct for JSON stream
	fileCfg := new(db.Configuration)

	if argsCount == 1 {
		fmt.Println("No arguments supplied. Using default configurations")
		fmt.Printf("%s", help)

	} else if argsCount == 2 {

		// Open file for reading
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		// Decode JSON configuration file
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&fileCfg)
		if err != nil {
			log.Fatal(err)
		}

		// Call db-init.go to connect to database
		if err := db.Connect(fileCfg); err != nil {
			log.Printf("Error in main(): %v", err)
		}

	} else {

		// Connect using flag arguments
		flagCfg := db.Configuration{Username: user, Password: pwd, ConnectionString: connectionString}
		db.Connect(&flagCfg)

	}

	// Load routes
	routing.LoadRoutes()

	err := http.ListenAndServe(":80", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
