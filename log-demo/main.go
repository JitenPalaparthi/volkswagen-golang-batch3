//package main

// import (
// 	"log/slog"
// 	"os"
// )

// func main() {
// 	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
// 	logger.Info("User login", "user", "alice", "id", 42)
// }

// import "go.uber.org/zap"

// func main() {
// 	logger, _ := zap.NewProduction()
// 	defer logger.Sync()

// 	logger.Info("User login", zap.String("user", "alice"), zap.Int("id", 42))
// }

package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	///log.Print("hello world")
	log.Info().Msg("Hello World")
}
