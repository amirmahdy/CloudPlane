package api

import (
	db "cloudplane/db/model"
	"cloudplane/docs"
	"cloudplane/internal"
	"cloudplane/token"

	"log"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	store  db.Store
	config internal.Config
	token  token.Maker
	router *gin.Engine
}

func NewServer(config internal.Config, store db.Store) (server *Server) {
	token, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker", err)
	}
	server = &Server{
		store:  store,
		config: config,
		token:  token,
	}
	server.setupRoutes()
	// todo: add middleware
	return
}

func (server *Server) setupRoutes() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.POST("/api/user/login", server.loginUser)
	router.POST("/api/user/create", server.createUser)

	authRouters := router.Group("/").Use(authMiddleware(server.token))
	authRouters.POST("/api/profile/create", server.createProfile)

	server.router = router
}

func (server *Server) Run() error {
	return server.router.Run(server.config.ServerAddress)
}
