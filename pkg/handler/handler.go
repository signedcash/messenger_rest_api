package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		messages := api.Group("/messages")
		{
			messages.POST("/", h.createMesssage)
			messages.GET("/", h.getAllMessages)
			messages.GET("/:id", h.getMessageById)
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
