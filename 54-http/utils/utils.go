package utils

import (
	"demo/models"
	"os"

	"github.com/rs/zerolog/log"
)

var UserChan chan *models.User

const FILENAME string = "users.txt"

func init() {
	UserChan = make(chan *models.User)
	go ProcessUsers(FILENAME)
}

func Init(fileName string) {
	UserChan = make(chan *models.User)
	go ProcessUsers(fileName)
}

func ProcessUsers(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	defer file.Close()
	for user := range UserChan {
		bytes, err := user.ToSting()
		if err != nil {
			log.Error().Msg(err.Error())
		}
		_, err = file.Write([]byte(bytes + "\n"))
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}

}

func SaveFile(filename string, user *models.User) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := user.ToSting()
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(bytes + "\n"))
	if err != nil {
		return err
	}

	return nil
}
