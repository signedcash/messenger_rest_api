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
			messages.PUT("/:id", h.updateMessage)
			messages.DELETE("/:id", h.deleteMessage)

		}

		chats := api.Group("/chats")
		{
			chats.POST("/", h.createChat)
			chats.GET("/", h.getAllChats)
			chats.GET("/:id", h.getChatById)
			chats.PUT("/:id", h.updateChat)
			chats.DELETE("/:id", h.deleteChat)

		}

		profiles := api.Group("/profiles")
		{
			profiles.POST("/", h.createProfile)
			profiles.GET("/", h.getAllProfiles)
			profiles.GET("/:id", h.getProfileById)
			profiles.PUT("/:id", h.updateProfile)
			profiles.DELETE("/:id", h.deleteProfile)

		}
	}

	return router
}
