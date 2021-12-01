package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	textme "github.com/signedcash/messenger_rest_api"
)

func (h *Handler) createChat(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input textme.Chat
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.User1Id != userId && input.User2Id != userId {
		newErrorResponse(c, http.StatusInternalServerError, "Invalid token or chat")
		return
	}

	id, err := h.services.Chat.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllChatsByUserId(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	chats, err := h.services.Chat.GetAllByUserId(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, chats)
}

func (h *Handler) getChatByUserId(c *gin.Context) {
	user1Id, err := getUserId(c)
	if err != nil {
		return
	}

	user2Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	chat, err := h.services.Chat.GetByUserId(user1Id, user2Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, chat)
}

func (h *Handler) updateChat(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input textme.UpdateChatInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Chat.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
