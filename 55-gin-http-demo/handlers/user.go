package handlers

import (
	"demo/database"
	"demo/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UserHandler struct {
	database.IUserDB // prmoted field
}
type IUserHandler interface {
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	DeleteUserByID(ctx *gin.Context)
}

func NewUserHandler(iuserdb database.IUserDB) IUserHandler {
	return &UserHandler{iuserdb}
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	user := new(models.User)

	// bytes, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	//
	// }
	// err = json.Unmarshal(bytes, user)
	err := ctx.Bind(user)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}

	err = user.Validate()

	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}

	//user.Id = uint(rand.IntN(1000))
	user.Status = "active"
	user.LastModifed = uint64(time.Now().Unix())

	user, err = u.Create(user)
	if err != nil {
		log.Error().Msg(err.Error())
		ctx.String(http.StatusBadRequest, "Seems something went wrong")
		ctx.Abort()
		return
	}

	ctx.JSON(201, user) // return

}

func (u *UserHandler) UpdateUser(ctx *gin.Context) {
	user := new(models.User)

	// bytes, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	//
	// }
	// err = json.Unmarshal(bytes, user)
	err := ctx.Bind(user)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		ctx.Abort()
		return
	}

	_, err = u.GetBy(fmt.Sprint(user.Id))
	if err != nil {
		log.Error().Msg(err.Error())
		ctx.String(http.StatusBadRequest, "user does not exist")
		ctx.Abort()
		return
	}

	user.Status = "active"
	user.LastModifed = uint64(time.Now().Unix())
	user, err = u.Update(user)
	if err != nil {
		log.Error().Msg(err.Error())
		ctx.String(http.StatusBadRequest, "Seems something went wrong")
		ctx.Abort()
		return
	}

	ctx.JSON(201, user) // return

}

func (u *UserHandler) GetUserByID(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.String(http.StatusBadRequest, "invalid path param id")
		ctx.Abort()
		return
	}

	user, err := u.GetBy(id)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.String(http.StatusBadRequest, "record not found")
		} else {
			log.Error().Msg(err.Error())
			ctx.String(http.StatusBadRequest, "Seems something went wrong")
		}
		ctx.Abort()
		return
	}
	ctx.JSON(200, user) // return
}

func (u *UserHandler) DeleteUserByID(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.String(http.StatusBadRequest, "invalid path param id")
		ctx.Abort()
		return
	}

	err := u.DeleteBy(id)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.String(http.StatusBadRequest, "record not found")
		} else {
			log.Error().Msg(err.Error())
			ctx.String(http.StatusBadRequest, "Seems something went wrong")
		}
		ctx.Abort()
		return
	}
	ctx.String(200, "user with id ->"+id+" successfully deleted") // return
}
