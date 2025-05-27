package main

import (
	"demo/database"
	"demo/handlers"
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	PORT string
	DSN  string
)

func main() {

	PORT = os.Getenv("PORT")
	if PORT == "" {
		flag.StringVar(&PORT, "port", "8084", "-port=8084 or --port=8084 --port 8084")
	}

	DSN = os.Getenv("DSN")
	if DSN == "" {
		flag.StringVar(&DSN, "dsn", `host=localhost user=admin password=admin123 dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`, "-dsn=db connection string or --dsn=db connection string --dsn db connection string")
	}

	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	r := gin.Default() // includes logger and also recovery feature

	// r.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.String(200, "Pong")
	// })
	db, err := database.GetConnection(DSN)
	if err != nil {
		panic(err)
	} else {
		log.Info().Msg("successfully connected to the database")
	}

	r.GET("/", handlers.Root)
	r.GET("/ping", handlers.Ping)
	r.GET("/health", handlers.Health)

	userRoute := r.Group("/private/users")
	userDB := database.NewUserDB(db)
	userHandler := handlers.NewUserHandler(userDB)
	userRoute.POST("/user", userHandler.CreateUser)
	userRoute.PUT("/user", userHandler.UpdateUser)
	userRoute.GET("/user/:id", userHandler.GetUserByID)       // path param
	userRoute.DELETE("/user/:id", userHandler.DeleteUserByID) // path para

	if err := r.Run(":" + PORT); err != nil {
		log.Error().Msg(err.Error())
	}

}

// Gorilla Mux, Beego, Echo, Fiber, Gin , Http
// Gin also uses standard http package internally

// Fast Http Routing
// Middleware supprot
// Json bind and validataion
// built in error handling
// Gouped routing
// crash free recover middleware
// zero memory allocation routing
// inbuilt logger
// working with request params

// go get -u github.com/gin-gonic/gin
