package api

import (
	"fmt"

	db "github.com/LeMinh0706/simplebank/db/sqlc"
	"github.com/LeMinh0706/simplebank/token"
	"github.com/LeMinh0706/simplebank/util"
	"github.com/gin-gonic/gin"
)

// /http request for bankink service
type Server struct {
	config     util.Config
	queries    *db.Queries
	tokenMaker token.Maker
	router     *gin.Engine
}

// /Create a new http server and setup routing
func NewServer(config util.Config, queries *db.Queries) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}
	server := &Server{
		config:     config,
		queries:    queries,
		tokenMaker: tokenMaker,
	}

	/// add routes to router
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {

	router := gin.Default()

	router.POST("/users/register", server.createUser)
	router.POST("/users/login", server.loginUser)
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/users/profile", server.myProfile)
	authRoutes.PUT("/users", server.updatePosition)
	authRoutes.PUT("/users/avt", server.updateAvatar)
	authRoutes.POST("/users", server.getUsers)
	authRoutes.GET("/search", server.searchUser)
	server.router = router

	router.Static("upload/avatar", "./upload/avatar")

}

// Start runs the HTTP server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
