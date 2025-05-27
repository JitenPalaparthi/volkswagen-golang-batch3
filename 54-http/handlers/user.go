package handlers

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"io"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type UserHandler struct {
	FileName string
}
type IUserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(filename string) IUserHandler {
	return &UserHandler{filename}
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	//if r.Method != "POST" {
	if r.Method != http.MethodPost {
		w.Write([]byte("not a valid http request method"))
		w.WriteHeader(http.StatusNotImplemented)
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Msg(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := new(models.User)
	err = json.Unmarshal(bytes, user) // unmarshal data coming from requst to user model object
	if err != nil {
		log.Error().Msg(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = user.Validate()
	if err != nil {
		log.Error().Msg(err.Error())
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Id = uint(rand.IntN(10000))
	user.Status = "active"
	user.LastModifed = uint64(time.Now().Unix())

	// err = utils.SaveFile(u.FileName, user)
	// if err != nil {
	// 	log.Error().Msg(err.Error())
	// 	w.Write([]byte(err.Error()))
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	utils.UserChan <- user // sending user to the channel

	w.Write([]byte("user successfully inserted"))
	//w.WriteHeader(http.StatusCreated)

}
