package handler

import (
	library "libraryapp"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags books
// @Description create book
// @ID create-book
// @Accept  json
// @Produce  json
// @Param input body library.Book true "book info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/book [post]
func (h *Handler) createBook(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	var input library.Book
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Book.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllBooksResponse struct {
	Data []library.Book `json:"data"`
}

func (h *Handler) getAllBooks(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	books, err := h.services.Book.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBooksResponse{
		Data: books,
	})
}

func (h *Handler) getBookById(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	book, err := h.services.Book.GetById(bookId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) updateBook(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var input library.UpdateBookInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Book.Update(bookId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteBook(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	err = h.services.Book.Delete(bookId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
