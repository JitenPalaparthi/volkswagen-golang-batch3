package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//dsn := `host=localhost user=admin password=admin123 dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`

const TIMEOUT = 10
const TIMEOUTCHECKDURATION = 1 // seconds

func GetConnection(dsn string) (*gorm.DB, error) {
	timeout := time.After(TIMEOUT * time.Second)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	count := 1
	if err != nil {
		for {
			select {
			case <-timeout:
				return nil, errors.New("tried to connect to the dabase but timedout")
			default:
				log.Info().Msg(fmt.Sprint("trynig to connect for number of times-->", count))
				time.Sleep(time.Second)
				db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
				if err == nil {
					return db, nil
				}
				count++
			}
		}
	} else {
		return db, err
	}
}
