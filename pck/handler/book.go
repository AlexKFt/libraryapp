package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createBook(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllBooks(c *gin.Context) {

}

func (h *Handler) getBookById(c *gin.Context) {

}

func (h *Handler) updateBook(c *gin.Context) {

}

func (h *Handler) deleteBook(c *gin.Context) {

}
