package main

import (
	"demo/handlers"
	"demo/utils"
	"flag"
	"net/http"
	"os"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	println("main calling --1 in utils")
}

func init() {
	println("main calling --2 in utils")
}

var (
	PORT string
)

func main() {
	// file, err := os.Create("log.txt")
	// if err != nil {
	// 	log.Error().Msg(err.Error())
	// }

	// log.Logger = log.Output(file)

	//args := os.Args
	//fmt.Println(args)

	// println("--------> assigning filename")
	// utils.Filename = "users.txt"
	// utils.Process("users.dat")

	PORT = os.Getenv("PORT")
	if PORT == "" {
		flag.StringVar(&PORT, "port", "8084", "-port=8084 or --port=8084 --port 8084")
		flag.Parse()
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/health", handlers.Health)

	userhandler := handlers.NewUserHandler("users.txt")

	http.HandleFunc("/user", userhandler.Create)

	log.Info().Msg("Server started and running on port " + PORT)

	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Error().Msg(err.Error())
		close(utils.UserChan)
		close(utils.UserErrorChan)
		runtime.Goexit()
	}
}

// Trace
// Debug
// Info
// Warning
// Error
// Panic
// Fatal
