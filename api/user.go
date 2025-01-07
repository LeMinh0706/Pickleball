package api

import (
	"net/http"

	db "github.com/LeMinh0706/simplebank/db/sqlc"
	"github.com/LeMinh0706/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Gender   int32  `json:"gender"`
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

	// image, err := ctx.FormFile("image")

	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// }

	// if !util.ExtCheck(image.Filename) {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("ảnh không đúng định dạng")))
	// }
	// fileName := fmt.Sprintf("upload/%s/%d%s", "avt", time.Now().Unix(), filepath.Ext(image.Filename))
	// if err := ctx.SaveUploadedFile(image, fileName); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }
	uuid, _ := uuid.NewRandom()
	user, err := server.queries.CreateUser(ctx, db.CreateUserParams{
		ID:       uuid,
		Username: req.Username,
		Password: hashedPassword,
		Fullname: req.Fullname,
		Gender:   req.Gender,
		Avt:      util.RandomAvatar(req.Gender),
		Lat:      0,
		Lng:      0,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (server *Server) loginUser(ctx *gin.Context) {

}
