package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/signedcash/messenger_rest_api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.GET("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		messages := api.Group("/messages")
		{
			messages.POST("/", h.createMesssage)
			messages.GET("/:chat_id", h.getAllMessagesByChatId)
			messages.GET("/last/:chat_id", h.getLastMessageByChatId)
			messages.PUT("/:id", h.updateMessage)
			messages.DELETE("/:id", h.deleteMessage)

		}

		chats := api.Group("/chats")
		{
			chats.POST("/", h.createChat)
			chats.GET("/", h.getAllChatsByUserId)
			chats.GET("/:id", h.getChatByUserId)
			chats.PUT("/:id", h.updateChat)
		}

		profiles := api.Group("/profiles")
		{
			profiles.GET("/:id", h.getProfileByUserId)
			profiles.PUT("/", h.updateProfile)
		}

		users := api.Group("/users")
		{
			users.GET("/:id", h.getUserById)
			users.GET("/search/:name", h.getUserByName)
			users.GET("/", h.getUser)
			users.PUT("/", h.updateUser)
		}
	}

	return router
}
