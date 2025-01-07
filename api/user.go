package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	db "github.com/LeMinh0706/simplebank/db/sqlc"
	"github.com/LeMinh0706/simplebank/token"
	"github.com/LeMinh0706/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Register struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Fullname string  `json:"fullname"`
	Gender   int32   `json:"gender"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req Register
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	uuid, _ := uuid.NewRandom()
	user, err := server.queries.CreateUser(ctx, db.CreateUserParams{
		ID:       uuid,
		Username: req.Username,
		Password: hashedPassword,
		Fullname: req.Fullname,
		Gender:   req.Gender,
		Avt:      util.RandomAvatar(req.Gender),
		Lat:      req.Lat,
		Lng:      req.Lng,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req Login
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.queries.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("sai tài khoản hoặc mật khẩu")))
		return
	}
	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(fmt.Errorf("sai tài khoản hoặc mật khẩu")))
		return
	}
	token, err := server.tokenMaker.CreateToken(req.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, LoginResponse{AccessToken: token})
}

func (server *Server) myProfile(ctx *gin.Context) {
	auth := ctx.MustGet(authorizationPayLoadKey).(*token.Payload)

	user, err := server.queries.GetMyProfile(ctx, auth.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type UpdatePosition struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (server *Server) updatePosition(ctx *gin.Context) {
	var req UpdatePosition
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	auth := ctx.MustGet(authorizationPayLoadKey).(*token.Payload)
	user, err := server.queries.GetMyProfile(ctx, auth.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	update, err := server.queries.UpdatePosition(ctx, db.UpdatePositionParams{ID: user.ID, Lat: user.Lat, Lng: user.Lng})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, update)
}

type Position struct {
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Radius float64 `json:"radius"`
}

func (server *Server) getUsers(ctx *gin.Context) {
	var req Position
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	users, err := server.queries.GetUsers(ctx, db.GetUsersParams{
		Radians:   req.Lat,
		Radians_2: req.Lng,
		Lat:       req.Radius,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (server *Server) searchUser(ctx *gin.Context) {
	param := ctx.Query("name")
	pageStr := ctx.Query("page")
	pageSizeStr := ctx.Query("page_size")

	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	users, err := server.queries.SearchUser(ctx, db.SearchUserParams{
		Column1: sql.NullString{String: param, Valid: true},
		Limit:   int32(pageSize),
		Offset:  (int32(page) - 1) * int32(pageSize),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (server *Server) updateAvatar(ctx *gin.Context) {
	auth := ctx.MustGet(authorizationPayLoadKey).(*token.Payload)
	user, err := server.queries.GetMyProfile(ctx, auth.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	image, err := ctx.FormFile("image")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !util.ExtCheck(image.Filename) {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("ảnh không đúng định dạng")))
		return
	}
	fileName := fmt.Sprintf("upload/%s/%d%s", "avt", time.Now().Unix(), filepath.Ext(image.Filename))
	if err := ctx.SaveUploadedFile(image, fileName); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	update, err := server.queries.UpdateAvatar(ctx, db.UpdateAvatarParams{
		ID:  user.ID,
		Avt: fileName,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, update)
}
