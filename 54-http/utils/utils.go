package utils

import (
	"demo/models"
	"os"

	"github.com/rs/zerolog/log"
)

var UserChan chan *models.User
var UserErrorChan chan error

const FILENAME string = "users.txt"

func init() {
	UserChan = make(chan *models.User)
	UserErrorChan = make(chan error)
	go ProcessUsers(FILENAME)
	go ProcessErrors()
}

func Init(fileName string) { // if you want to save to a non constant file name, call this in main
	UserChan = make(chan *models.User)
	UserErrorChan = make(chan error)
	go ProcessUsers(fileName)
	go ProcessErrors()
}

func ProcessUsers(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		//log.Error().Msg(err.Error())
		UserErrorChan <- err
	}

	defer file.Close()
	for user := range UserChan {
		bytes, err := user.ToSting()
		if err != nil {
			UserErrorChan <- err
			//log.Error().Msg(err.Error())
		}
		_, err = file.Write([]byte(bytes + "\n"))
		if err != nil {
			UserErrorChan <- err
			//log.Error().Msg(err.Error())
		}
	}

}

func ProcessErrors() {
	// file, err := os.Create("log.txt")
	// if err != nil {
	// 	log.Error().Msg(err.Error()) //
	// }

	// log.Logger = zerolog.New(file)
	for err := range UserErrorChan {
		log.Error().Msg(err.Error()) // what ever the error you would like to do that
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
