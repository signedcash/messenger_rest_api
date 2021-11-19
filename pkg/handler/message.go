package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	textme "github.com/signedcash/messenger_rest_api"
)

func (h *Handler) createMesssage(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input textme.Message
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.SenderId != userId {
		newErrorResponse(c, http.StatusInternalServerError, "Invalid token or message")
		return
	}

	id, err := h.services.Message.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllMessagesByChatId(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	chatId, err := strconv.Atoi(c.Param("chat_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Invalid chat id param")
		return
	}

	messenges, err := h.services.Message.GetAllByChatId(userId, chatId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if messenges == nil {
		newErrorResponse(c, http.StatusInternalServerError, "no messages for this chat id or the token is invalid")
		return
	}

	c.JSON(http.StatusOK, messenges)
}

func (h *Handler) updateMessage(c *gin.Context) {

}

func (h *Handler) deleteMessage(c *gin.Context) {

}
